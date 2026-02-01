package usecase

import (
	"context"
	"io"
	"destinyface/internal/domain/repository"
	"destinyface/internal/usecase/dto"
)

type ProfileUseCaseInterface interface {
	SetupProfileImage(ctx context.Context, userID string, file io.Reader) (*dto.UserOutput, error)
}

type profileUseCase struct {
	userRepo repository.UserRepositoryInterface
	fileRepo repository.FileStorageInterface
}

// コンストラクタで初期化処理
func NewProfileUseCase(
	ur repository.UserRepositoryInterface,
	fr repository.FileStorageInterface,
) ProfileUseCaseInterface {
	return &profileUseCase{
		userRepo: ur,
		fileRepo: fr,
	}
}

func (u *profileUseCase) SetupProfileImage(ctx context.Context, userID string, file io.Reader) (*dto.UserOutput, error) {
	// 画像を送ったユーザー名を取得
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// ローカルストレージへ保存する
	uploadPath := user.GenerateProfilePath()
	imageURL, err := u.fileRepo.Upload(ctx, file, uploadPath)
	if err != nil {
		return nil, err
	}

	// Entityのルールを適用し、画像保存できたので、ステータスをACTIVEにする
	if err := user.CompleteProfileSetup(imageURL); err != nil {
		return nil, err
	}

	// DBを更新する
	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	// レスポンスを返す
	return &dto.UserOutput{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		ProfileImageURL: user.ProfileImageURL,
		Status: user.Status,
		CreatedAt: user.CreatedAt,
	}, nil
}