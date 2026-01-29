package repository

import (
	"context"
	"io"
)

type FileStorageInterface interface {
	// Uploadは画像データを保存し、S3からアクセス可能なURLを返す。
	// fileNameは、"users/{userID}/profile.jpg"のような固定パスを想定する。
	Upload(ctx context.Context, data io.Reader, fileName string) (string, error)
}