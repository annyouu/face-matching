# API仕様書(V2.0)

1. 認証・アカウント系 (/auth)
- POST /auth/register: 新規会員登録。セッションを作成し、Cookieをセットする。
- POST /auth/login: ログイン。セッションを作成し、Cookieをセットする。
- POST /auth/logout: ログアウト。サーバー側のセッション（Redis）を削除し、Cookieを破棄する。

2. ユーザープロフィール・セットアップ系 (/users)
- GET /users/me: 自分のプロフィール情報の取得。
- PATCH /users/me: ニックネームやステータスなどの基本情報更新。
- PATCH /users/setup/name: 新規登録直後のニックネーム設定。ステータスを PENDING_IMAGE 等へ進行させる。
- PATCH /users/setup/image: 新規登録直後の顔画像登録。
    - 挙動: Goが画像を受け取り → Python(gRPC)でベクトル化 → pgvectorへ保存 → ステータスを COMPLETED に更新。

3. 顔類似検索系 (/matches)
- POST /matches/search: 検索用画像のアップロードと類似検索の実行。
    - リクエスト: multipart/form-data (画像ファイル)。
    - レスポンス: 類似度スコア付きユーザーリスト。

- GET /matches/users/{userId}: 検索結果から特定のユーザーの詳細（プロフィール、公開画像等）を取得。

4. メッセージ・制限系 (/messages)
- GET /messages/rooms: チャットルーム一覧の取得（最新メッセージと相手情報含む）。
- GET /messages/rooms/{userId}: 特定の相手とのメッセージ履歴取得。
- POST /messages: メッセージの送信。
- GET /messages/limit: 本日の残り送信可能枠（0〜3）の取得。

5. 画像配信系 (/images)
- GET /images/{fileName}: サーバーまたはストレージに保存された画像を表示。
    - 注意：S3等の外部ストレージ利用時は、各APIのレスポンス（profile_image_url 等）に署名付きURLを直接含める運用を考える必要あり。