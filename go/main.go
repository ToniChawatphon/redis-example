package main

import (
	"encoding/json"
	"log"

	"github.com/ToniChawatphon/redis-example/app"
	uuid "github.com/satori/go.uuid"
)

// SetGetExample
func SetGetExample(r app.Redis) {
	var key, value string
	key, value = "1", "golang-redis"

	log.Println("===== Example 1: Set value =====")
	log.Printf("SET '%v: %v'", key, value)
	r.Set(key, value)
	res := r.Get(key)
	log.Printf("GET value '%v'", res)
}

// GetAllKeysExample
func GetAllKeysExample(r app.Redis) {
	log.Println("===== Example 2: Get all keys =====")
	key1 := "1"
	res1 := r.GetKeys(key1)
	log.Printf("Get KEYS %v: %v", key1, res1)

	// get all existing keys
	key2 := "*"
	res2 := r.GetKeys(key2)
	log.Printf("Get KEYS %v: %v", key2, res2)
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

	log.Println("===== Example 3: Struct to Json =====")
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
	log.Println("===== Example 4: HGet all =====")
	hashKey := "x01x11x1"
	log.Printf("List all from '%v' hkeys", hashKey)
	hashes := r.HGetAll(hashKey)
	for h, v := range hashes {
		log.Println(h, v)
	}
}

// AddItemIntoListExample item can be duplicated
func AppendItemIntoListExample(r app.Redis) {
	log.Println("===== Example 5: Append item into list =====")
	var coins []string
	var position string

	key := "coin"
	position = "right" // can be 'left' or 'right'
	coins = []string{"BTC", "SOL"}

	log.Printf("Append '%v'", coins)
	r.Append(position, key, coins)
	res := r.GetList(key)
	log.Printf("Get list %v", res)
}

// AddItemIntoSetExample no duplicate item in set
func AddItemIntoSetExample(r app.Redis) {
	log.Println("===== Example 6: Add items into set =====")
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
	GetAllKeysExample(r)

	StructToJsonExample(r)
	HGetAllExample(r)

	AppendItemIntoListExample(r)
	AddItemIntoSetExample(r)
}
