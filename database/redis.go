package database

import (
	"os"

	"github.com/go-redis/redis"
)

type RedisDatabase struct {
	client *redis.Client
}

/*
Connnect database server and
Create NewClient

You must open the db port
before connect db
> redis-server
> redis-cli
*/
func InitializeRedis() (Database, error) {
	dbAddr := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     dbAddr,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, &CreateDatabaseError{}
	}
	return &RedisDatabase{client: client}, nil
}

/*
Create, Update
*/
func (r *RedisDatabase) Set(key string, value string) (string, error) {
	result, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		return generateError("set", err)
	}
	return result, nil
}

/*
Read by key
*/
func (r *RedisDatabase) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return generateError("get", err)
	}
	return value, nil
}

/*
Delete value contain this key
*/
func (r *RedisDatabase) Delete(key string) (string, error) {
	_, err := r.client.Del(key).Result()
	if err != nil {
		return generateError("delete", err)
	}
	return key, nil
}

func generateError(operation string, err error) (string, error) {
	if err == redis.Nil {
		return operation, &OperationError{operation: operation}
	}

	return operation, &DownError{}
}
