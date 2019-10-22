package cache

import (
	"github.com/go-redis/redis"
	"log"
	"vkBot/conversation"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6380",
})

func GetDialogState(authorId string) conversation.State {
	state, err := redisClient.Get(authorId).Int()
	if err != nil {
		log.Panic("Error while conversion redis value by key ", authorId, err)
	}
	return conversation.State(state)
}

func SetDialogState(key string, value int) {
	_, err := redisClient.Set(key, value, 0).Result()
	if err != nil {
		log.Panic("Error while setting value to redis cache by key", key, err)
	}
}
