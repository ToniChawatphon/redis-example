package app

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

// Redis struct
type Redis struct {
	rdb      *redis.Client
	RedisUrl string
	Password string
	DB       int
}

// Connect makes connection
func (r *Redis) Connect() {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     r.RedisUrl,
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
	})
}

// Set key and value
func (r *Redis) Set(key string, value string) string {
	res, err := r.rdb.Set(context.Background(), key, value, 0).Result()
	if err != nil {
		log.Fatal(err)
	}

	return res
}

// Get value from key
func (r *Redis) Get(key string) string {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Fatal(err)
	}

	return val
}

// HSet hash key and value
func (r *Redis) HSet(key string, data interface{}) int64 {
	res, err := r.rdb.HSet(context.Background(), key, "json", hashStruct(data)).Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	return res
}

// HGet value from hash key
func (r *Redis) HGet(id string) string {
	res, err := r.rdb.HGet(context.Background(), id, "json").Result()
	if err != nil {
		log.Println(err)
	}

	return res
}

// HGetAll value from hash key
func (r *Redis) HGetAll(key string) map[string]string {
	res, err := r.rdb.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	return res
}

// hashStruct converts struct to byte
func hashStruct(data interface{}) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err.Error())
	}

	return d
}
