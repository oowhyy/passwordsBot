package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/oowhyy/passwordbot/internal/storage"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisStorage struct {
	Client         *redis.Client
	ExpireDuration time.Duration
}

func (rs *RedisStorage) CompositeKey(username string, service string) string {
	return fmt.Sprintf("{%s}%s", username, service)
}

func NewRedisStorage(addr string, password string, db int, expire int) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.New("unable to connect to redis")
	}
	return &RedisStorage{Client: client, ExpireDuration: time.Second * time.Duration(expire)}, nil
}

func (rs *RedisStorage) Set(username string, item *storage.Item) error {
	key := rs.CompositeKey(username, item.Service)
	if err := rs.Client.Set(ctx, key, item.Password, rs.ExpireDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (rs *RedisStorage) Get(username, service string) (*storage.Item, error) {
	key := rs.CompositeKey(username, service)
	res, err := rs.Client.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}
	item := &storage.Item{
		Service:  service,
		Password: res,
	}
	return item, nil
}

func (rs *RedisStorage) Delete(username string, service string) (int64, error) {
	key := rs.CompositeKey(username, service)
	res, err := rs.Client.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (rs *RedisStorage) Expire() time.Duration {
	return rs.ExpireDuration
}
