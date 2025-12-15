package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"destinyface/internal/domain/errors"
	"destinyface/internal/domain/entity"
	"destinyface/internal/domain/repository"
	"destinyface/internal/usecase/dto"

	"github.com/go-playground/validator/v10"
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
	validator *validator.Validate
}

// コンストラクタ
func NewUserUseCase(userRepo repository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepo: userRepo,
		validator: validator.New(),
	}
}

// 新規登録
func (u *UserUseCase) Register(ctx context.Context, input *dto.UserRegisterInput) (*dto.UserOutput, error) {

	// DTOタグに定義されたルールでバリデーションを実行
	if err := u.validator.Struct(input); err != nil {
		return nil, domain.ErrInvalidInput
	}

	// Emailの重複チェック
	_, err := u.userRepo.FindByEmail(ctx, input.Email)

	// nilだったら、ユーザーが重複しているということ
	if err == nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("FindByEmailの処理に失敗しました: %w", err)
	}

	// パスワードのハッシュ化を行う
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, fmt.Errorf("パスワードのハッシュ化に失敗しました： %w", err)
	}

	// エンティティの生成
	now := time.Now()
	newUser := &entity.User{
		ID:           uuid.New().String(),
        Email:        input.Email,
        PasswordHash: string(hashedPassword),
        Name:         input.Name,
        CreatedAt:    now,
        UpdatedAt:    now,
	}

	// リポジトリを呼び出し、DBへ保存する
	if createdErr := u.userRepo.Create(ctx, newUser); createdErr != nil {
		return nil, fmt.Errorf("DBへの保存に失敗しました: %w", createdErr)
	}

	// 成功レスポンスを返す
	return &dto.UserOutput{
		ID:    newUser.ID,
        Name:  newUser.Name,
        Email: newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}, nil
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

