package usecase

import (
	"context"
	stdErrors "errors"
	"fmt"
	"time"

	"destinyface/internal/domain/entity"
	"destinyface/internal/domain/repository"
	appErrors "destinyface/internal/errors"
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

	// 確認フェーズ
	SetupName(ctx context.Context, userID string, input *dto.UserSetupNameInput) (*dto.UserOutput, error)
}

type UserUseCase struct {
	userRepo repository.UserRepositoryInterface
	sessionRepo repository.SessionRepositoryInterface
	validator *validator.Validate
}

// コンストラクタ
func NewUserUseCase(userRepo repository.UserRepositoryInterface, sessionRepo repository.SessionRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepo: userRepo,
		sessionRepo: sessionRepo,
		validator: validator.New(),
	}
}

// 新規登録
func (u *UserUseCase) Register(ctx context.Context, input *dto.UserRegisterInput) (*dto.UserOutput, error) {

	// DTOタグに定義されたルールでバリデーションを実行
	if err := u.validator.Struct(input); err != nil {
		return nil, appErrors.ErrInvalidInput
	}

	// Emailの重複チェック
	_, err := u.userRepo.FindByEmail(ctx, input.Email)

	// nilだったら、ユーザーが重複しているということ
	if err == nil {
		return nil, appErrors.ErrEmailAlreadyExists
	}

	if !stdErrors.Is(err, appErrors.ErrNotFound) {
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
		Status: "PENDING_NAME",
        CreatedAt:    now,
        UpdatedAt:    now,
	}

	// リポジトリを呼び出し、DBへ保存する
	if createdErr := u.userRepo.Create(ctx, newUser); createdErr != nil {
		return nil, fmt.Errorf("DBへの保存に失敗しました: %w", createdErr)
	}

	// JWTトークン生成の代わりに、Redisセッションを作成
	sessionID, err := u.sessionRepo.CreateSession(ctx, newUser.ID)
	if err != nil {
		return nil, fmt.Errorf("セッション生成に失敗しました: %w", err)
	}

	// 成功レスポンスを返す
	return &dto.UserOutput{
		ID:    newUser.ID,
        Email: newUser.Email,
		Status: newUser.Status,
		Token: sessionID,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (u *UserUseCase) SetupName(ctx context.Context, userID string, input *dto.UserSetupNameInput) (*dto.UserOutput, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 状態チェック
	if user.Status != "PENDING_NAME" {
		return nil, fmt.Errorf("invalid status for name setup: %s", user.Status)
	}

	user.Name = input.Name
	user.Status = "PENDING_IMAGE"
	user.UpdatedAt = time.Now()

	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserOutput{
		ID: user.ID,
		Name: user.Name,
		Status: user.Status,
	}, nil
}

// ログイン
func (u *UserUseCase) Login(ctx context.Context, input *dto.UserLoginInput) (*dto.AuthTokenOutput, error) {
	// 400 バリデーションチェック (400 Bad Request 相当)
	if err := u.validator.Struct(input); err != nil {
		return nil, appErrors.ErrInvalidInput
	}

	// 401 ユーザーの存在確認 (401 Unauthorized 相当)
	user, err := u.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("userRepo.FindByEmailに失敗しました: %w", err)
	}

	// 401 パスワードの照合 (401 Unauthorized 相当)
	// 入力されたパスワードとDBに保存されているハッシュ値を比較する
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, appErrors.ErrInvalidCredentials
	}

	// 認証成功した場合：JWTのトークンを発行の代わりに、Redisセッション保存へ
	sessionID, err := u.sessionRepo.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("セッションの生成に失敗しました: %w", err)
	}

	return &dto.AuthTokenOutput{
		Token: sessionID,
		Status: user.Status,
	}, nil
}

// プロフィール取得
func (u *UserUseCase) GetProfile(ctx context.Context, userID string) (*dto.UserOutput, error) {
	// ユーザーのIDの形式チェック
	if userID == "" {
		return nil, appErrors.ErrInvalidInput
	}

	// リポジトリからユーザー取得
	user, err := u.userRepo.FindByID(ctx, userID)

	// 500
	if err != nil {
		return nil, fmt.Errorf("userRepoのFindByIDは失敗しました: %w", err)
	}

	// DTOにして返す
	return &dto.UserOutput{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		ProfileImageURL: user.ProfileImageURL,
		CreatedAt: user.CreatedAt,
	}, nil
}

// プロフィール更新
func (u *UserUseCase) UpdateProfile(ctx context.Context, userID string, input *dto.UserUpdateInput) (*dto.UserOutput, error) {
	// バリデーションチェック
	if err := u.validator.Struct(input); err != nil {
		return nil, appErrors.ErrInvalidInput
	}

	// 更新対象のユーザーがいるかどうか確認
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("userRepoのFindByIDは失敗しました: %w", err)
	}

	// エンティティの値を更新する
	// 入力があった項目のみ更新する
	if input.Name != "" {
		user.Name = input.Name
	}

	if input.ProfileImageURL != "" {
		user.ProfileImageURL = input.ProfileImageURL
	}

	user.UpdatedAt = time.Now()

	// リポジトリで更新
	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("userRepo.Updateに失敗しました: %w", err)
	}

	// 更新後の情報を返す
	return &dto.UserOutput{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		ProfileImageURL: user.ProfileImageURL,
		CreatedAt: user.CreatedAt,
	}, nil
}

