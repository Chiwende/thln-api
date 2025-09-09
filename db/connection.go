package db

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// 	"qinsights.com/thln/initializers"
// )

// // RedisConfig holds Redis configuration
// type RedisConfig struct {
// 	Host     string
// 	Port     string
// 	Password string
// 	DB       int
// }

// // DefaultRedisConfig returns default Redis configuration from environment variables
// func DefaultRedisConfig() *RedisConfig {
// 	return &RedisConfig{
// 		Host:     initializers.GetEnv("REDIS_HOST", "localhost"),
// 		Port:     initializers.GetEnv("REDIS_PORT", "6379"),
// 		Password: initializers.GetEnv("REDIS_PASSWORD", ""),
// 		DB:       0, // Default Redis database
// 	}
// }

// // ConnectRedis establishes a connection to Redis
// func ConnectRedis(config *RedisConfig) (*redis.Client, error) {
// 	if config == nil {
// 		config = DefaultRedisConfig()
// 	}

// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
// 		Password: config.Password,
// 		DB:       config.DB,
// 	})

// 	// Test the connection
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
// 	}

// 	log.Println("Successfully connected to Redis")
// 	return rdb, nil
// }

// // CloseRedis closes the Redis connection
// func CloseRedis(rdb *redis.Client) error {
// 	if rdb != nil {
// 		return rdb.Close()
// 	}
// 	return nil
// }
