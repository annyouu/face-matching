<!-- ER図について考える -->

## データベース設計 (PostgreSQL + pgvector) (ベクトル検索)

アプリケーションのコアとなる6つのテーブルと、そのリレーションを示します。

### 1. ER図 (実体関連図)

多対多の関係（Likes）と、チャット機能の構造（Conversations, Messages）が核となります。

```mermaid
erDiagram
    USERS {
        uuid id PK
        string name
        string email
        string password_hash
    }
    FACES {
        uuid id PK
        uuid user_id FK "誰の顔か"
        vector embedding "顔特徴量"
    }
    LIKES {
        uuid id PK
        uuid sender_id FK "メッセージを送った人"
        uuid receiver_id FK "メッセージを送られた人"
        text message_content
    }
    DAILY_LIMITS {
        uuid user_id FK, PK
        date date PK "日付ごとの制限"
        int count "送信済み人数"
    }
    CONVERSATIONS {
        uuid id PK
        uuid user_a_id FK
        uuid user_b_id FK
    }
    MESSAGES {
        uuid id PK
        uuid conversation_id FK "どの会話に属するか"
        uuid sender_id FK "送信者"
        text content "メッセージ本文"
    }

    USERS ||--o{ FACES : "has (1対多)"
    USERS ||--o{ LIKES : "sends/receives (多対多)"
    USERS ||--o{ DAILY_LIMITS : "managed by (1対多)"
    CONVERSATIONS ||--o{ MESSAGES : "contains (1対多)"
    
    % USERS と CONVERSATIONS の関係 (多対多を Likes/Conversationsで実現)
    USERS ||--o{ CONVERSATIONS : "participates (多対多)"
    LIKES }o--o{ CONVERSATIONS : "initiates (1対1 or 1対多)"
```

### 2. テーブル構成と役割

|   | テーブル名 | 主な役割 | リレーションの性質 | 特記事項 |
| :--- | :--- | :--- | :--- | :--- |
| **1** | **Users** | ユーザーの基本情報を管理する | 1対多 | 特になし |
| **2** | **Faces** | **顔特徴量 (VECTOR(512))** を保存。ユーザーデータから顔データを分離（正規化）を意識。 | 多対1 (Users) | pgvector使用 |
| **3** | **Likes** | **中間テーブル**としての役割。 | 多対多 (Users) | `(sender_id, receiver_id)` のユニーク性確保 |
| **4** | **Conversations**| **チャットルームID**を管理する。 | 多対多 (Users) | MessagesテーブルのFK |
| **5** | **Messages** | チャット履歴（トランザクション）を記録。 | 多対1 (Conversations) | 特になし |
| **6** | **Daily_Limits**| **1日3人制限**のカウントを管理する。 | 多対1 (Users) | `(user_id, date)` の複合主キー |

### 3. テーブル作成 SQL (DDL)

プロジェクトの初期セットアップに使用するSQL定義です。`pgvector` の拡張機能を有効化する必要があります。

```sql
-- 1. pgvector 拡張機能の有効化
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS vector;

-- 2. Users テーブル
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 3. Faces テーブル (コア機能)
CREATE TABLE faces (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    embedding VECTOR(512) NOT NULL, -- 512次元の顔特徴量
    is_primary BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
-- 検索速度向上のため HNSW インデックスを適用推奨 (大規模データ向け)
-- CREATE INDEX ON faces USING hnsw (embedding vector_cosine_ops); 


-- 4. Likes テーブル (中間テーブル/初手メッセージ記録)
-- sender_id と receiver_id の組み合わせは一意 (重複メッセージ防止)
CREATE TABLE likes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender_id UUID REFERENCES users(id) ON DELETE CASCADE,
    receiver_id UUID REFERENCES users(id) ON DELETE CASCADE,
    message_content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE (sender_id, receiver_id)
);

-- 5. Daily_Limits テーブル (制限管理)
CREATE TABLE daily_limits (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    count INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, date) -- user_id と date の組み合わせで一意
);

-- 6. Conversations テーブル (チャットルーム)
CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_a_id UUID REFERENCES users(id) ON DELETE RESTRICT,
    user_b_id UUID REFERENCES users(id) ON DELETE RESTRICT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 7. Messages テーブル (チャット履歴)
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID REFERENCES conversations(id) ON DELETE CASCADE,
    sender_id UUID REFERENCES users(id) ON DELETE RESTRICT,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);