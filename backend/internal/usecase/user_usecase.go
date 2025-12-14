package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"destinyface/internal/domain"
	"destinyface/internal/domain/entity"
	"destinyface/internal/domain/repository"
	"destinyface/internal/usecase/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseInterface interface {
	// 新規登録
	Register(ctx context.Context, input *dto.UserRegisterInput) (*dto.UserOutput, error)
	// ログイン
	Login(ctx context.Context, input *dto.UserLoginInput) (*dto.AuthTokenOutput, error)
	// プロフィール取得
	GetProfile(ctx context.Context, userID string) (*dto.UserOutput, error)
	// プロフィール更新
	UpdateProfile(ctx context.Context, userID string, input *dto.UserUpdateInput) (*dto.UserOutput, error)
}

type UserUseCase struct {
	userRepo repository.UserRepositoryInterface
}

// コンストラクタ
// func NewUserUseCase(userRepo repository.UserRepositoryInterface) UserUseCaseInterface {
// 	return &UserUseCase{
// 		userRepo: userRepo
// 	}
// }

// 新規登録
func (u *UserUseCase) Register(ctx context.Context, input *dto.UserRegisterInput) (*dto.UserOutput, error) {

	// いずれかの入力値が空かどうか
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return nil, domain.ErrInvalidInput
	}

	// Email形式の正規化チェック(詳細なバリデーション)

	// 入力値がすでにあるか重複しているかどうか

	// OKだったら、パスワードのハッシュかを行う


	// エンティティの生成

	// リポジトリを呼び出し、DBへ保存する

	// 成功レスポンスを返す
}

// ログイン
func (u *UserUseCase) Login() {
	// 401

	// 403

	// 404

	// 400
}

// プロフィール取得
func (u *UserUseCase) GetProfile() {
	
}

// プロフィール更新
func (u *UserUseCase) UpdateProfile() {
	
}

