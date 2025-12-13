package repository

import (
	"context"
	"destinyface/internal/domain/entity"
)

// ユーザーデータの永続化を抽象化する
// Create (新規登録、新しいユーザー情報をDBに挿入する)
// Read (ログイン時にユーザーが入力したものが正しいかの確認)
// Update (プロフィールの更新)
// Delete (アカウントの削除)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}