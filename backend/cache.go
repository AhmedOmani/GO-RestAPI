package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client;
var ctx = context.Background();
const cacheDuration = 15 * time.Minute;

func InitRedis() {
	err := godotenv.Load();
	if err != nil {
		log.Println("Warning: Could not load .env file for Redis config");
	}

	redisAddr 	  := os.Getenv("REDIS_ADDR");
	redisPassword := os.Getenv("REDIS_PASSWORD");
	redisDBStr    := os.Getenv("REDIS_DB");
	redisDB , err := strconv.Atoi(redisDBStr);

	rdb = redis.NewClient(&redis.Options{
		Addr:		redisAddr,
		Password: 	redisPassword,
		DB: 		redisDB,
	});

	_ , err = rdb.Ping(ctx).Result();
	if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
	
	log.Println("Redis connection established successfully.");
}

//Generate redis key for user email.
func getUserCacheKey(email string) string {
	return "user:email:" + email;
}

//Stores user object in redis cache.
func SetUserInCache(user *User) error {
	key := getUserCacheKey(user.Email);

	jsonData , err := json.Marshal(user);
	if err != nil {
        log.Printf("Error marshaling user data for caching (key: %s): %v", key, err)
        return err
    }

	err = rdb.Set(ctx , key , jsonData , cacheDuration).Err();
	if err != nil {
        log.Printf("Error setting user in Redis cache (key: %s): %v", key, err)
        return err
    }

    log.Printf("User data cached successfully for key: %s (TTL: %v)", key, cacheDuration);
	return nil
}

//Retrieve the user cached in redis.
func GetUserInCache(email string) (*User, error) {
	key := getUserCacheKey(email);
	val , err := rdb.Get(ctx , key).Result();

	if err == redis.Nil {
		log.Printf("Cache MISS for key: %s" , key);
		return nil , nil ; 
	} else if err != nil {
		log.Printf("Error getting user from Redis cache (key: %s): %v", key, err);
        return nil, err // Actual Redis error
	}

	log.Printf("Cache HIT for key: %s" , key);
	user := &User{};
	err = json.Unmarshal([]byte(val) , user);
	
	if err != nil {
		log.Printf("Error unmarshaling cached user data (key: %s): %v", key, err);
		return nil, err
	}

	return user, nil ;
}

//Explicitly removes a user from cache ,
//Even if we not using in current flow but it is important in some cases.
func InvalidateUserCache(email string) error {
	key := getUserCacheKey(email);
	err := rdb.Del(ctx , key).Err();
	if err != nil && err != redis.Nil {
		log.Printf("Error deleting user cache key %s: %v", key, err)
		return err
	}
	log.Printf("Invalidated cache for key: %s successfully" , key);
	return nil ;
}