package main

import (
	"log"

	"github.com/ToniChawatphon/redis-example/app"
	uuid "github.com/satori/go.uuid"
)

type Model struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// SetGetExample
func SetGetExample(r app.Redis) {
	log.Println("===== Example 1: SET GET =====")
	var key, value string
	key, value = "1", "golang-redis"
	r.Set(key, value)
	res := r.Get(key)
	log.Printf("GET value '%v'", res)
}

func SetGetJsonExample(r app.Redis) {
	log.Println("===== Example 2: SET GET JSON =====")
	key := "2"
	model := &Model{
		Id:   uuid.NewV4().String(),
		Name: "Hi Golang",
	}
	r.SetJson(key, model)
	res := r.Get(key)
	log.Printf("GET json value %v", res)
}

// GetAllKeysExample
func GetKeysExample(r app.Redis) {
	log.Println("===== Example 3: Get KEYS =====")
	key1 := "1"
	res1 := r.GetKeys(key1)
	log.Printf("Get KEYS %v: %v", key1, res1)

	// get all existing keys
	key2 := "*"
	res2 := r.GetKeys(key2)
	log.Printf("Get KEYS %v: %v", key2, res2)
}

// HsetSGetExample
func HsetSGetExample(r app.Redis) {
	log.Println("===== Example 4: HGet HSet =====")
	var key, field1, field2 string
	key = "3"
	field1 = "animal"
	field2 = "food"

	r.HSet(key, field1, "cat")
	r.HSet(key, field2, "noodles")
	res1 := r.HGet(key, field1)
	res2 := r.HGet(key, field2)
	log.Printf("HGet field: '%v', value: '%v'", field1, res1)
	log.Printf("HGet field: '%v', value: '%v'", field2, res2)
}

// StructToJsonExample using hset and hget
func HSetHGetJsonExample(r app.Redis) {
	// variables
	key := "x01x11x1"
	field := "json"

	model := &Model{
		Id:   uuid.NewV4().String(),
		Name: "Hi Golang",
	}

	log.Println("===== Example 5: HSet HGet Json  =====")
	r.HSetJson(key, field, model)
	res := r.HGet(key, field)
	log.Printf("GET field: '%v', value: %v", field, res)
}

// HGetAllExample get all field & value from key
func HGetExample(r app.Redis) {
	log.Println("===== Example 6: HGet all =====")
	key := "x01x11x1"
	log.Printf("Get all field & value from key: '%v'", key)
	hashes := r.HGetAll(key)
	for h, v := range hashes {
		log.Println(h, v)
	}
}

// AddItemIntoListExample item can be duplicated
func AppendItemIntoListExample(r app.Redis) {
	log.Println("===== Example 7: Append item into list =====")
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
	log.Println("===== Example 8: Add items into set =====")
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

	// set get
	SetGetExample(r)
	SetGetJsonExample(r)
	GetKeysExample(r)

	// hset hget
	HsetSGetExample(r)
	HSetHGetJsonExample(r)
	HGetExample(r)

	// list & set
	AppendItemIntoListExample(r)
	AddItemIntoSetExample(r)
}
