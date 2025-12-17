package dto

import "time"

// -- 入力データ --

// 新規登録用、そのままのパスワードを受け取る
type UserRegisterInput struct {
	Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"` 
    Password string `json:"password" validate:"required,min=8"`
}

// ログイン用
type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 顔写真登録用

// プロフィール更新用、更新したい項目のみ受け取る
type UserUpdateInput struct {
	Name string  `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
}

// -- 出力データ --

// ユーザー情報を返す用
type UserOutput struct {
	ID string  `json:"id"`
	Name string  `json:"name"`
	Email string  `json:"email"`
	ProfileImageURL string `json:"profile_image_url"`
	CreatedAt time.Time `json:"created_at"`
}

// ログイン成功時にトークンを返す用
type AuthTokenOutput struct {
	Token string `json:"token"`
}