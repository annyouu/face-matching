package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"destinyface/internal/domain/repository"
)

type localStorage struct {
	baseDir string
}

// LocalStorageImplの実体を作成するコンストラクタ
func NewLocalStorage(dir string) repository.FileStorageInterface {
	// 保存先ディレクトリがなければ作っておく
	if _, err := os.Stat(dir); err != nil {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
	return &localStorage{
		baseDir: dir,
	}
}

func (s *localStorage) Upload(ctx context.Context, file io.Reader, path string) (string, error) {
	// 保存先のフルパスを作成 (例: uploads/users/123/profile.jpg)
	fullPath := filepath.Join(s.baseDir, path)

	// 親ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return "", err
	}

	// 保存用のファイルを作成
	out, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// io.Reader(file)の中身を、作成したファイル(out)にコピーする
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	// ローカルでのパスを返す
	return "/uploads/" + path, nil
}