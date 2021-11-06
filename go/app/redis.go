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
		Password: r.Password,
		DB:       r.DB,
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

// HSet set key, field and value
func (r *Redis) HSet(key string, field string, data interface{}) int64 {
	res, err := r.rdb.HSet(context.Background(), key, field, hashStruct(data)).Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	return res
}

// HGet field, value from key
func (r *Redis) HGet(id, field string) string {
	res, err := r.rdb.HGet(context.Background(), id, field).Result()
	if err != nil {
		log.Println(err)
	}

	return res
}

// HGetAll field, value from hash key
func (r *Redis) HGetAll(key string) map[string]string {
	res, err := r.rdb.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	return res
}

// Append add item into list based on key
func (r *Redis) Append(key string, values []string) error {
	_, err := r.rdb.RPush(context.Background(), key, values).Result()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (r *Redis) GetList(key string) []string {
	res, err := r.rdb.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (r *Redis) AddUnique(key string, field []string) int64 {
	res, err := r.rdb.SAdd(context.Background(), key, field).Result()
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (r *Redis) GetUniqueSet(key string) []string {
	res, err := r.rdb.SMembers(context.Background(), key).Result()
	if err != nil {
		log.Fatal(err)
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
