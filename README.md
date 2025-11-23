# 顔マッチングサービス — 仕様書
## 概要
このサービスは、オーストラリアの「ペットと里親の顔マッチング」アプリの着想を得て、人間版として実装するものです。人はこの世にドッペルゲンガーがいると言われます。ユーザーがアップロードした顔画像をもとに、「似ている他のユーザー」を検出・提示し、簡易チャットでつながれる Web アプリケーションです。
顔特徴量の抽出には Python の ML ライブラリ（例: insightface / facenet-pytorch）を使用する。最初はローカル CPU、必要に応じて GPU に移行。
顔特徴の抽出サービス（Python）とアプリケーションロジック（Go）は gRPC で接続する。
フロントエンドは Next.js + TypeScript を想定（Web メイン）。将来 Flutter / Swift に移植する予定。

# 1.要件サマリ
## 機能 (MVP)
- ユーザー登録/ログイン
- 顔画像のアップロード
- 顔特徴量(embedding)の生成、保存
- 類似ユーザーの検索
- 類似度スコアの表示
- 類似ユーザーとのチャット(WebSocket)

# 2. 全体図
```mermaid
flowchart LR

%% ======================
%% User / Frontend
%% ======================
subgraph "Frontend (Next.js)"
    UI[ユーザー UI<br>画像アップロード / 類似ユーザー一覧]
end

%% ======================
%% Backend API (Go)
%% ======================
subgraph "API Layer (Go)"
    AUTH[Auth Handler<br>(JWT/OAuth)]
    API[API Gateway<br>(HTTP / gRPC)]
    MATCH[Matcher Service<br>(類似度計算・検索)]
end

%% ======================
%% Embedding Server (Python)
%% ======================
subgraph "Embedding Layer (Python)"
    PY[Face Embedding Server<br>(Face Detection/Alignment/Embedding)]
end

%% ======================
%% Database
%% ======================
subgraph "Database Layer"
    DB[(PostgreSQL + pgvector)]
end

%% ======================
%% External API
%% ======================
subgraph "External Services"
    GCP[Google Cloud Vision API<br>(※検討中: 顔検出)]
end

%% ======================
%% Frontend → Backend
%% ======================
UI -->|画像アップロード| AUTH
AUTH -->|JWT検証| API
API -->|画像データ送信 (gRPC)| PY

%% ======================
%% Python Embedding Server
%% ======================
PY -->|Embedding生成(512次元)| API

%% ======================
%% Go側：類似度マッチング
%% ======================
API -->|Embedding受取| MATCH
MATCH -->|pgvector検索| DB
MATCH -->|類似ユーザー結果| API
API -->|レスポンス返却| UI

%% ======================
%% Optional: Vision API
%% ======================
PY -->|顔検出API呼び出し| GCP
```

# 2-add. アーキテクチャ詳細 (クリーンアーキテクチャ & DDD の採用)

# 3. 処理の流れ

## ① Frontend → Go API
Next.jsから画像をアップロードする

## ② Go → Python (gRPC)
画像をGo APIからPythonサーバへ送る。
Pythonは以下のものを担当する。
- 顔検出（Vision API）
- 顔前処理 (アライメント)
- 512次元 embedding 抽出

## ③ Python → Go (gRPC)
PythonからembeddingをGoに返す。

## ④ 類似度検索をGoが行う。
- embedding を受け取り
- pgvector で類似検索
- コサイン類似度でスコア算出する
- 類似ユーザー一覧を返却する

## ⑤ Go → Frontend
結果 (似ているユーザー)を返す。

# 4. フロントエンド仕様 (Next.js+TypeScript)
## 主要ページ
- /singup, /login
- /profile：プロフィール編集 (名前、写真、公開範囲)
- /upload：顔写真アップロード
- /matches：類似ユーザー一覧 (スコア、サムネイル、チャットボタン)
- /chat/[userId]：チャット画面

## なぜNext.jsを採用するか、Reactとの比較

## TypeScriptを採用する理由

# 5. バックエンド仕様 (Go)

# 6. 顔認証/顔類似性サービス (Python)

# 7. API仕様書 (REST for frontend, gRPC for service間)

# 8. DBデータモデル




