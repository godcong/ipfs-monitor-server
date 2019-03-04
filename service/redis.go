package service

import (
	"github.com/go-redis/redis"
)

// RedisQueueIndex ...
const RedisQueueIndex = 1

// RedisKeyStoreIndex ...
const RedisKeyStoreIndex = 2

// RedisIPFSServiceIndex ...
const RedisIPFSServiceIndex = 3

// RedisKeyNameIPFSPins ...
const RedisKeyNameIPFSPins = "ipfs.pins"

// RedisKeyNameIPFSSwarmAddress ...
const RedisKeyNameIPFSSwarmAddress = "ipfs.swarm.address"

var store *redis.Client

// NewKeyStore ...
func NewKeyStore() *redis.Client {
	return newRedisWithDB(RedisKeyStoreIndex)
}

// NewRedisQueue ...
func NewRedisQueue() *redis.Client {
	return newRedisWithDB(RedisQueueIndex)
}

// NewIPFSServiceRedis ...
func NewIPFSServiceRedis() *redis.Client {
	return newRedisWithDB(RedisIPFSServiceIndex)
}

// newRedisWithDB ...
func newRedisWithDB(idx int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       idx, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

// InitRedis ...
func InitRedis(c *redis.Client) {
	store = c
}
