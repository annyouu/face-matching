package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	
	"destinyface/internal/domain/entity"
	"destinyface/internal/domain/repository"
)

// DBへの実際の処理を記述
type usersRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &usersRepositoryImpl{
		db: db,
	}
}

// Create 新規ユーザーの登録を行う
func (r *usersRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	// INSERTクエリ
	query := `
	INSERT INTO users (id, email, password_hash, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	if user.CreatedAt.IsZero() {
        user.CreatedAt = time.Now()
    }
    if user.UpdatedAt.IsZero() {
        user.UpdatedAt = time.Now()
    }

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("userの作成に失敗しました: %w", err)
	}

	fmt.Println("userの作成に成功しました")
	return nil
}

// FindByID IDによるユーザー検索
func (r *usersRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.User, error) {
	user := &entity.User{}

	// クエリ実行
	query := `
	SELECT id, email, password_hash, name, created_at, updated_at
	FROM users
	WHERE id = $1
	`
	// QueryRowContextでクエリを実行し、行を取得
	row := r.db.QueryRowContext(ctx, query, id)

	// row.Scan()を使用して取得した値をエンティティのフィールドに入れる
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	// エラーチェックを行う
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("userが見つかりませんでした: %w", err)
	}
	
	return user, nil

}

// Update ユーザー情報の更新
func (r *usersRepositoryImpl) Update(ctx context.Context, user *entity.User) error {

	// Updateクエリを定義する
	query := `
	UPDATE users
	SET email = $2, password_hash = $3, name = $4, updated_at = $5
	WHERE id = $1
	`

	// user.UpdatedAtを現在の時刻に設定する
	user.UpdatedAt = time.Now()

	// ExecContextでクエリを実行する
	result, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("userの更新に失敗しました: %w", err)
	}

	// ResultからRowsAffected()を取得し、更新された行数を取得
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	
	return nil
}

// Delete ユーザーの削除
func (r *usersRepositoryImpl) Delete(ctx context.Context, id string) error {
	// DELETEクエリを定義
	query := `DELETE FROM users WHERE id = $1`

	// ExecContextでクエリを実行する
	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("userの削除に失敗しました: %w", err)
	}

	// ResultからRowsAffected()を取得し、削除された行数が0の場合、repository.ErrNotFound に変換して返す
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	
	return nil
}