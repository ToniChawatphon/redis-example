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

	log.Println("===== Example 1: Set value =====")
	log.Printf("SET '%v: %v'", key1, value1)
	r.Set(key1, value1)
	res := r.Get(key1)
	log.Printf("GET value '%v'", res)
}

// StructToJsonExample using hset and hget
func StructToJsonExample(r app.Redis) {

	// json struct
	type Model struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	// variables
	hashKey := "x01x11x1"
	field1 := "json1"
	field2 := "json2"

	model := &Model{
		Id:   uuid.NewV4().String(),
		Name: "Hi Golang",
	}

	log.Println("===== Example 2: Struct to Json =====")
	log.Printf("HSET key: '%v', field: '%v ", hashKey, field1)
	log.Printf("HSET key: '%v', field: '%v ", hashKey, field2)
	r.HSet(hashKey, field1, model)
	r.HSet(hashKey, field2, model)
	res1 := r.HGet(hashKey, field1)
	res2 := r.HGet(hashKey, field2)
	log.Printf("GET field '%v', value %v", field1, res1)
	log.Printf("GET field '%v', value %v", field2, res2)

	// convert json back to struct
	newModel1 := &Model{}
	err := json.Unmarshal([]byte(res1), newModel1)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// convert json back to struct
	newModel2 := &Model{}
	err = json.Unmarshal([]byte(res2), newModel2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("struct result '%v'", newModel1)
	log.Printf("struct result '%v'", newModel2)
}

// HGetAllExample get all field & value from key
func HGetAllExample(r app.Redis) {
	log.Println("===== Example 3: HGet all =====")
	hashKey := "x01x11x1"
	log.Printf("List all json from '%v'", hashKey)
	hashes := r.HGetAll(hashKey)
	for h, v := range hashes {
		log.Println(h, v)
	}
}

// AddItemIntoListExample item can be duplicated
func AppendItemIntoListExample(r app.Redis) {
	log.Println("===== Example 4: Append item into list =====")
	var coins []string
	key := "coin"
	coins = []string{"BTC", "ETH", "ADA"}

	log.Printf("Append '%v'", coins)
	r.Append(key, coins)
	res := r.GetList(key)
	log.Printf("Get list %v", res)
}

// AddItemIntoSetExample no duplicate item in set
func AddItemIntoSetExample(r app.Redis) {
	log.Println("===== Example 5: Add items into set =====")
	var blockchain []string
	key := "blockchain"
	blockchain = []string{"ETH", "BNB", "ETH"}
	log.Printf("Add items %v", blockchain)
	r.AddUnique(key, blockchain)
	res := r.GetUniqueSet(key)
	log.Printf("Get all set %v", res)
}

func main() {
	r := app.Redis{
		RedisUrl: "localhost:6379",
		Password: "",
		DB:       0,
	}

	r.Connect()
	SetGetExample(r)
	StructToJsonExample(r)
	HGetAllExample(r)
	AppendItemIntoListExample(r)
	AddItemIntoSetExample(r)
}
