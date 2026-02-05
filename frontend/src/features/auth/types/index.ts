/**
 * ユーザーの進行ステータス
 */
export type UserStatus = 'PENDING_NAME' | 'PENDING_IMAGE' | 'ACTIVE';

/**
 * 1: 新規登録 入力 (UserRegisterInput)
 */
export interface UserRegisterInput {
  email: string;
  password: string;
}

/**
 * 2: 名前設定 入力 (UserSetupNameInput)
 */
export interface UserSetupNameInput {
  name: string;
}

/**
 * 3: 画像設定 入力 (UserSetupImageInput)
 */
export interface UserSetupImageInput {
  profile_image_url: string;
}

/**
 * ログイン 入力 (UserLoginInput)
 */
export interface UserLoginInput {
  email: string;
  password: string;
}

/**
 * 汎用ユーザー出力 (UserOutput)
 * Register, GetProfile, UpdateProfile などのレスポンスに使用
 */
export interface UserResponse {
  id: string;
  name: string;
  email: string;
  profile_image_url: string;
  status: UserStatus;
  token?: string; // omitempty 対応
  created_at: string;
}

/**
 * ログイン成功時 出力 (AuthTokenOutput)
 */
export interface AuthTokenResponse {
  token: string;
  status: UserStatus;
}