package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"destinyface/internal/domain/entity"
	"destinyface/internal/domain/repository"
	"destinyface/internal/usecase/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInterface interface {
	// 新規登録
	Register(ctx context.Context, input *dto.UserRegisterInput) (*dto.UserOutput, error)
	// ログイン
	Login(ctx context.Context, input *dto.UserLoginInput) (*dto.AuthTokenOutput, error)
	// プロフィール取得
	GetProfile(ctx context.Context, userID string) (*dto.UserOutput, error)
	// プロフィール更新
	UpdateProfile(ctx context.Context, userID string, input *dto.UserUpdateInput) (*dto.UserOutput, error)
}

