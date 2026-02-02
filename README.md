# Daburi Zero — 仕様書
「これ、家にあったっけ？」「家にあるのにまた同じものを買ってしまった。。。」を画像検索で解決する、うっかり重複買い防止プラットフォーム

## 概要
「Daburi Zero」は、洗剤、調味料、掃除用具など、日用品の「ダブリ買い」をゼロにするための管理アプリです。本サービスでは、「画像（パッケージデザイン）そのもの」をベクトル化して検索することで、在庫状況を照合する体験を提供します。

# 1.要件サマリ
- 商品を撮影するだけで在庫登録・検索ができること
- 画像の類似度（特徴量ベクトル）に基づいて、既存在庫との一致を判定すること
- 在庫の有無を「％（類似度）」と「画像比較」で直感的にユーザーへ伝えること
- 外出先（オフラインに近い環境）でも高速に検索結果を返せること

## 機能 (MVP)
1. 認証機能
    - ユーザー登録 / ログイン
    - JWTを使用

2. 顔登録・Embedding生成
    - ユーザーが自身の顔画像をアップロード
    - Python 製の ML サービスで顔特徴量（Embedding）を生成
    - Embeddingをデータベースに保存する

3. 顔類似検索 (中核機能)
    - 検索したい顔画像（芸能人・人物写真など）をアップロード
    - 入力画像から Embedding を生成
    - 登録ユーザーの顔Embeddingと類似度を計算
    - 類似度順にユーザー一覧を表示（Top N）

4. 検索結果表示
    - 顔サムネイル
    - 類似度に基づく並び順
    - 最低限のプロフィール情報 (ニックネーム等)

5. メッセージ送信
    - 気になった相手に1日最大3人までメッセージを送信可能

# 2. 全体図
```mermaid
flowchart LR

%% Frontend
subgraph FE["Frontend (Next.js / Mobile)"]
    UI["カメラUI / 在庫一覧 / 照合結果表示"]
end

%% Backend
subgraph API_Layer["API Layer (Go)"]
    AUTH["Auth Handler"]
    UPLOAD["Inventory Handler (登録・検索窓口)"]
    VEC_QUERY["Vector Search Engine"]
end

%% Embedding Server
subgraph PY_Layer["ML Layer (Python)"]
    PY["Feature Extractor (ResNet/ViT) / Image Preprocessing"]
end

%% Database
subgraph DB_Layer["Database Layer"]
    DB[(PostgreSQL + pgvector)]
end

%% フロー
UI -->|商品画像アップロード| UPLOAD
UPLOAD -->|gRPC| PY
PY -->|特徴量ベクトル返却| UPLOAD
UPLOAD -->|ベクトル検索 / 保存| DB
DB -->|検索結果返却| UPLOAD
UPLOAD -->|判定結果（類似度＋画像）| UI
```

# 3. ディレクトリ構成
backend
```
go-backend/
├── cmd/
│   └── api/
│       └── main.go           # エントリーポイント (DI、ルーティング設定、サーバー起動)
├── internal/
│   ├── domain/               # ビジネスロジックの中心 (外部ライブラリに依存しない)
│   │   ├── entity/           # エンティティ & 値オブジェクト
│   │   │   ├── user.go       # User Entity, UserId ValueObject
│   │   │   └── face.go       # FaceEmbedding ValueObject
│   │   ├── repository/       # リポジトリの「インターフェース」定義
│   │   │   ├── user_repo.go
│   │   │   └── face_repo.go
│   │   └── service/          # ドメインサービス (純粋なドメインロジックがあれば)
│   │       └── similarity.go # 例: 類似度判定の閾値ロジックなど
│   │
│   ├── usecase/              # アプリケーションロジック
│   │   ├── user_usecase.go   # "ユーザー登録する" などの処理フロー
│   │   ├── match_usecase.go  # "画像を元に類似ユーザーを探す" フロー
│   │   └── inputport/        # UseCaseへの入力データの定義 (DTO的な役割)
│   │
│   ├── controller/            # 入出力の変換 (Controller/Presenter)
│   │   ├── http/             # REST API ハンドラ (Echo/Ginなど)
│   │   │   ├── handler.go
│   │   │   ├── request.go    # JSONリクエストの構造体
│   │   │   └── response.go   # JSONレスポンスの構造体
│   │   └── websocket/        # チャット用 WebSocket ハンドラ
│   │
│   └── infrastructure/       # 技術的な詳細実装 (DB, 外部API)
│       ├── persistence/      # リポジトリの実装、永続化処理 (PostgreSQL + pgvector)
│       │   ├── db.go
│       │   ├── user_repo_impl.go
│       │   └── face_repo_impl.go
│       ├── grpc/             # PythonサービスへのgRPCクライアント実装
│       │   └── face_client.go
│       └── router/           # Webフレームワークのルーティング設定
│
├── pkg/                      # プロジェクト外でも使える汎用ユーティリティ (Logger, Errorなど)
├── api/                      # gRPCの .proto ファイル定義
│   └── proto/
│       └── face_service.proto
├── go.mod
└── go.sum
```

frontend
```
src/
├── app/                        # 1. ルーティング (Next.js App Router)
│   ├── layout.tsx              # 全体共通（フォント、ヘッダー、フッター）
│   ├── page.tsx                # ランディングページ (Desktop-2)
│   ├── login/
│   │   └── page.tsx            # ログイン画面 (Desktop-1)
│   ├── signup/
│   │   └── page.tsx            # 新規登録画面
│   └── onboarding/
│       └── page.tsx            # 登録後フロー (名前入力、顔登録など)
├── components/                 # 2. ロジックを持たない共通UI部品
│   ├── Button.tsx              # Tailwindで装飾したボタン
│   ├── Input.tsx               # 入力フォーム部品
│   └── Card.tsx                # 白い枠線などの共通コンテナ
├── features/                   # 3. 画面（機能）ごとのロジックと部品
│   ├── common/                 # 複数機能で使うロジック入り部品
│   └── routes/                 # 特定ページ専用のコンポーネント
│       ├── auth/               # ログイン・新規登録用
│       │   ├── components/     # LoginForm.tsx, SignupForm.tsx
│       │   ├── hooks.ts        # ログイン処理、バリデーション
│       │   └── endpoint.ts     # Go APIを叩く関数 (Infrastructure)
│       └── onboarding/         # 登録後のステップ用 (FigmaのDesktop-14等)
│           ├── components/     # NameForm.tsx, FaceUpload.tsx
│           ├── hooks.ts        # ステップ管理、アニメーション制御
│           └── endpoint.ts     # プロフィール保存API
├── hooks/                      # 4. プロジェクト全体で使うReact Hooks
├── utils/                      # 5. 純粋な関数（日付変換、エラー処理など）
└── constants/                  # 6. 定数（APIのベースURL、文言など）
```

<!-- websocketによるチャット機能をfrontendのflowchart LR追加する -->

# 4. 処理の流れ
### ① Frontend → Go API
- 画像アップロード
    - Next.js から顔画像をアップロードする
    - 画像は multipart/form-data 形式で送信される
    - 認証済みユーザーのみ利用可能とする

### ② Go → Python (gRPC)
画像をGo APIからPythonサーバへ送る。
Pythonは以下のものを担当する。
- 顔検出（Vision API）
- 顔前処理 (アライメント)
- 512次元 embedding 抽出

### ③ Python → Go (gRPC)
- 解析結果の返却
    - 顔 Embedding（512次元ベクトル）

### ④ 類似度検索をGoが行う
- 検索および永続化
Go APIは以下を担当します。
- 受け取ったEmbeddingをデータベースに保存
    - ベクトル検索：pgvector
- 顔検索時は、Embeddingを用いて類似度計算を実行
- 類似度順にユーザーをソート (TOP N抽出)
※ 類似度検索のロジックはGo側で一元管理する。

### ⑤ Go → Frontend
- 検索結果の返却
    - 類似度順に並んだユーザー一覧
    - 顔サムネイル
    - 最低限のプロフィール情報
    - 類似度指標 (相対スコア)
Frontedは結果を表示し、ユーザーは気になった相手に1日3件までメッセージを送ることができる。

# 4. シーケンス図

# 5. アーキテクチャ詳細 (クリーンアーキテクチャ & DDD の採用)


# 6. フロントエンド仕様 (Next.js+TypeScript)
## 主要ページ
- /singup, /login：認証画面。
- /profile：プロフィール編集 (名前、写真、公開範囲)
- /upload：顔写真アップロード
- /matches：類似ユーザー一覧 (スコア、サムネイル、チャットボタン)
- /chat/[userId]：チャット画面

## なぜNext.jsを採用するか、Reactとの比較

## TypeScriptを採用する理由

# 7. バックエンド仕様 (Go)
## GolangとPythonの比較

<!-- あとで詳細図で別に飛ばす -->
# 8. 顔認証/顔類似性サービス (Python)

<!-- あとで詳細図で別に飛ばす -->
# 9. API仕様書 (REST for frontend, gRPC for service間)

<!-- あとで詳細図で別に飛ばす -->
# 10. DBデータモデル

# 11.コミットメッセージについて(12/21〜)
コミットメッセージが適当だったのでいかを基準にします。
- add: 新しい機能
- fix: バグの修正 (ほぼないかも)
- docs: ドキュメントの変更
- update: 機能修正(機能削除も含める)
- update_file: go.sum,go.modのファイル変更
- remove: ファイルの削除
- refactor: コード改善(リファクタ)