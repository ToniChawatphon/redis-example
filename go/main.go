package main

import (
	"encoding/json"
	"log"

	"github.com/ToniChawatphon/redis-example/app"
	uuid "github.com/satori/go.uuid"
)

// Model A json struct
type Model struct {
	Id   string
	Name string
}

// SetGetExample
func SetGetExample(r app.Redis) {
	var key1, value1 string
	key1, value1 = "1", "golang-redis"

	log.Println("===== Example 1 =====")
	r.Set(key1, value1)
	log.Printf("SET '%v: %v'", key1, value1)
	res := r.Get(key1)
	log.Printf("GET value '%v'", res)
}

// HsetHgetStructToJsonExample
func HsetHgetStructToJsonExample(r app.Redis) {
	type Model struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	hashKey := "x01x11x1"

	model := &Model{
		Id:   uuid.NewV4().String(),
		Name: "Hi Golang",
	}

	log.Println("===== Example 2 =====")
	log.Printf("HSET '%v':'%v'", hashKey, model)
	r.HSet(hashKey, model)
	res1 := r.HGet(hashKey)
	log.Printf("GET value %v", res1)

	newModel := &Model{}
	err := json.Unmarshal([]byte(res1), newModel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("result '%v'", newModel)
}

// HGetAllExample
func HGetAllExample(r app.Redis) {
	log.Println("===== Example 3 =====")
	hashKey := "x01x11x1"
	log.Printf("List all json from '%v'", hashKey)
	hashes := r.HGetAll(hashKey)
	for h, v := range hashes {
		log.Println(h, v)
	}
}

func main() {
	r := app.Redis{
		RedisUrl: "localhost:6379",
		Password: "",
		DB:       0,
	}

	r.Connect()
	SetGetExample(r)
	HsetHgetStructToJsonExample(r)
	HGetAllExample(r)
}
