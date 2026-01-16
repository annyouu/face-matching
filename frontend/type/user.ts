export type UserStatus = "PENDING_NAME" | "PENDING_IMAGE" | "ACTIVE";

/**
 * ユーザー情報の基本構造
 * バックエンドの dto.UserOutput と対応
 */
export interface UserResponse {
    id: string;
    email: string;
    name?: string;
    profile_image_url?: string;
    status: UserStatus;
    token?: string;
    created_at?: string;
}

/**
 * ログイン成功時のレスポンス
 * バックエンドの dto.AuthTokenOutput と対応
 */

export interface AuthTokenOutput {
    token: string;
    status: UserStatus;
}