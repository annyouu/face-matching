package repository

import (
	"context"
	"io"
)

type FileStorageInterface interface {
	// データをアップロードして、アクセス用のURLを返す
	// file: 画像のバイナリ (io.Reader)
	// path: EntityのGenerateProfilePath()で生成したパス
	Upload(ctx context.Context, file io.Reader, path string) (string, error) 
}