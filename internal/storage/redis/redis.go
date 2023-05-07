package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oowhyy/passwordbot/internal/storage"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisStorage struct {
	Client     *redis.Client
	ExpireTime time.Duration
}

func (rs *RedisStorage) CompositeKey(username string, service string) string {
	return fmt.Sprintf("{%s}%s", username, service)
}

func NewRedisStorage(addr string, password string, db int, expire int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("unable to connect to redis: ", err)
	}
	return &RedisStorage{Client: client, ExpireTime: time.Minute * time.Duration(expire)}
}

func (rs *RedisStorage) Set(username string, item *storage.Item) error {
	key := rs.CompositeKey(username, item.Service)
	if err := rs.Client.Set(ctx, key, item.Password, rs.ExpireTime).Err(); err != nil {
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
