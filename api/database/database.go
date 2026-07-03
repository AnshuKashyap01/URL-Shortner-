package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	client *redis.Client
	once   sync.Once
)

func CreateClient(dbNo int) *redis.Client {

	// Create Redis client only on the first call.
	once.Do(func() {

		client = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("DB_ADDR"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DB:       0,
		})

		_, err := client.Ping(Ctx).Result()
		if err != nil {
			log.Fatal("Redis connection failed:", err)
		}

		log.Println("Connected to Redis")
	})

	return client
}

// Application Starts
// ↓
// Create Client Once
// ↓
// Request 1
// ↓
// Request 2
// ↓
// Request 3
// ↓
// ...
