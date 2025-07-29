// Redis setup, TTL triggers, pub/sub handlers
package redis


import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}
	log.Println("Connected to Redis")

	go watchChaosEvents()
}

func watchChaosEvents() {
	pubsub := Rdb.PSubscribe(Ctx, "__keyevent@0__:expired")
	for msg := range pubsub.Channel() {
		log.Println("Chaos Triggered:", msg.Payload)
		// TODO: trigger game mutation or chaos round in game state
	}
}

func SetChaosTTL(key string, ttl time.Duration) {
	Rdb.Set(Ctx, key, "chaos", ttl)
}