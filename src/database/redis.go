package database

import (
	"github.com/go-redis/redis"
)

type redisDatabase struct {
	client *redis.Client
}

func createRedisDatabase() (Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, &CreateDatabaseError{}
	}
	return &redisDatabase{client: client}, nil
}

func (r *redisDatabase) RecordDB(key string, value string) (string, error) {
	_, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		return generateError("RecordDB", err)
	}
	return key, nil
}

func (r *redisDatabase) GetRecordDB(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return generateError("GetRecordDB", err)
	}
	return value, nil
}

func generateError(operation string, err error) (string, error) {
	if err == redis.Nil {
		return "", &OperationError{operation}
	}
	return "", &DownError{}
}
