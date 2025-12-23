package redis

import (
	"context"
    "time"
    "destinyface/internal/domain/repository"
    "github.com/google/uuid"
    "github.com/redis/go-redis/v9"
)

type sessionRepository struct {
	client *redis.Client
	expiry time.Duration
}

// コンストラクタ
func NewSessionRepository(client *redis.Client) repository.SessionRepositoryInterface {
	return &sessionRepository{
		client: client,
		expiry: 24 * time.Hour,
	}
}

func (r *sessionRepository) CreateSession(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()

	if err := r.client.Set(ctx, "session:"+sessionID, userID, r.expiry).Err(); err != nil {
		return "", err
	}
	return sessionID, nil
}

func (r *sessionRepository) GetUserID(ctx context.Context, sessionID string) (string, error) {
	userID, err := r.client.Get(ctx, "session:"+sessionID).Result()
	// キーがない
	if err == redis.Nil {
		return "", nil
	}
	// それ以外のエラー
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	if err := r.client.Del(ctx, "session:"+sessionID).Err(); err != nil {
		return err
	}
	return nil
}