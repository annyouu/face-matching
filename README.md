# Destiny Face 顔マッチングサービス — 仕様書

## 概要
従来のアプリケーションは、年収、年齢、身長などの表面的なスペックや条件による判断が先行し、ユーザーに心理的な負担を与えてきました。
本サービス「Destiny Face」は、新しいアプローチをします。
ユーザーの顔をAIが解析し、「芸能人・系統」および「雰囲気（Vibe）」といったAIタグとして資産化します。
これにより、ユーザーは自分の顔を好きと言ってくれる異性（＝自身のタイプを検索したユーザー）と効率的にマッチングし、
ストレスフリーな繋がりを目指します。

# 1.要件サマリ
AIが顔の「系統」と「雰囲気」を診断し、その顔を「タイプ」として探している異性とマッチングさせる、需要と供給を一致させるアプリ。

## 機能 (MVP)
- 認証機能: ユーザー登録およびログイン機能（JWT/OAuthを使用）。
- 自分の顔登録とAI診断:
    - 自分の顔画像をアップロードし、PythonのMLサービスで顔特徴量 (Embedding) を生成・保存。
    - MLが顔の特徴を解析し、自動でAIタグを付与。

- AIタグ検索:
    - ユーザーは、AIタグを利用して、自分の好みの異性を検索、閲覧することができる。

- タイプ検索：
    - 探したい顔の系統を示す情報（芸能人の名前、または自分の元恋人の画像）をアップロード。
    - アップロードされた画像に似た顔を検索、閲覧できる。

- コミュニケーション：
    - 1日3人にメッセージを送れるようにし、リアルタイムチャットが可能。有料会員は制限なし。

# 2. 全体図
```mermaid
flowchart LR

%% Frontend
subgraph FE["Frontend (Next.js)"]
    LOGIN["ログイン / 新規登録 UI"]
    UI["ユーザー UI 画像アップロード / 類似ユーザー一覧"]
end

%% Backend
subgraph API_Layer["API Layer (Go)"]
    AUTH["Auth Handler (JWT/OAuth)"]
    UPLOAD["Upload Handler"]
    MATCH["Matcher Service (類似度計算・検索)"]
end

%% Embedding Server
subgraph PY_Layer["Embedding Layer (Python)"]
    PY["Face Embedding Server (Face Detection/Alignment/Embedding)"]
end

%% Database
subgraph DB_Layer["Database Layer"]
    DB[(PostgreSQL + pgvector)]
end

%% External
subgraph EXT["External Services"]
    GCP["Google Cloud Vision API"]
end

%% 認証フロー
LOGIN -->|認証リクエスト| AUTH
AUTH -->|JWT発行| LOGIN

%% 画像アップロードフロー
UI -->|画像アップロード + JWT| UPLOAD
UPLOAD -->|gRPCで画像パス送信| PY
PY -->|gRPCでEmbedding生成| UPLOAD

UPLOAD -->|Embeddingを渡す| MATCH
MATCH -->|pgvector検索| DB
MATCH -->|類似ユーザー結果| UPLOAD
UPLOAD -->|レスポンス返却| UI

%% 外部 API 呼び出し
PY -->|顔検出 API 呼び出し| GCP
```

# 3. ディレクトリ構成
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

<!-- websocketによるチャット機能をfrontendのflowchart LR追加する -->

# 4. 処理の流れ
### ① Frontend → Go API
Next.jsから画像をアップロードする

### ② Go → Python (gRPC)
画像をGo APIからPythonサーバへ送る。
Pythonは以下のものを担当する。
- 顔検出（Vision API）
- 顔前処理 (アライメント)
- 512次元 embedding 抽出
- AIタグ分類ロジックを実行。

### ③ Python → Go (gRPC)
Embeddingと付与されたAIタグをGoに返す。

### ④ 類似度検索をGoが行う
- EmbeddingとAIタグをDB（pgvector）に保存。

### ⑤ Go → Frontend
結果 (似ているユーザー)を返す。

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
