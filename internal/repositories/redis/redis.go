package redis

import (
	"context"
	"fmt"
	_ "io"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	cfg "music-stream-service/internal/config"
	"music-stream-service/service/repository"
)

type RedisStorage struct {
	repository.RedisStorage
	DB *redis.Client
	context.Context
}

func NewRedisStorage(cfg cfg.RedisConfig) *RedisStorage {
	db := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		Username:     cfg.User,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to init redis: %v", err)
	}

	return &RedisStorage{DB: db, Context: context.Background()}
}

func (rs *RedisStorage) Set(key string, value interface{}, expiration time.Duration) error {
	err := rs.DB.Set(rs.Context, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set data, error: %s", err.Error())
	}
	return nil
}

func (rs *RedisStorage) Get(key string) (string, error) {
	val, err := rs.DB.Get(rs.Context, key).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", fmt.Errorf("failed to get value, error: %v\n", err)
	}
	return val, nil
}
