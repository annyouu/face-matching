package repository

import (
	"context"
	"destinyface/internal/domain/entity"
)

// ユーザーデータの永続化を抽象化する
// Create (新規登録、新しいユーザー情報をDBに挿入する)
// ブランチ確認
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}