package entity

import (
	"time"
	"errors"
)

const (
	StatusPendingName = "PENDING_NAME"
	StatusPendingImage = "PENDING_IMAGE"
	StatusActive = "ACTIVE"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	PasswordHash  string `json:"password_hash"`
	Name string `json:"name"`
	ProfileImageURL string
	Status string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// コンストラクタ
// Userオブジェクト生成時のルールをあらかじめ書く
func NewUser(id, email string) (*User, error) {
	if id == "" {
		return nil, errors.New("IDは必須です")
	}
	if email == "" {
		return nil, errors.New("メールアドレスは必須です")
	}

	return &User{
		ID: id,
		Email: email,
		Status: StatusPendingName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// 名前設定時のバリデーションとステータス移行を管理する
func (u *User) CompleteName(name string) error {
	if u.Status != StatusPendingName {
		return errors.New("名前を設定できる状況ではないです")
	}
	if name == "" {
		return errors.New("名前は必須です")
	}

	u.Name = name
	u.Status = StatusPendingImage
	// ここで更新する必要ってあるかな？
	u.UpdatedAt = time.Now()
	return nil
}

// 画像登録時のバリデーションとステータス遷移を管理する
func (u *User) CompleteProfileSetup(imageURL string) error {
	if u.Status != StatusPendingImage {
		return errors.New("画像を設定できる状態ではない")
	}

	if imageURL == "" {
		return errors.New("画像URLは必須です")
	}

	u.ProfileImageURL = imageURL
	u.Status = StatusActive
	u.UpdatedAt = time.Now()
	return nil
}

// S3の保存先パスを生成する
// これをEntityが持つことで、保存先ルールがアプリ全体で統一される
func (u *User) GenerateProfilePath() string {
	return "users/" + u.ID + "/profile.jpg"
}