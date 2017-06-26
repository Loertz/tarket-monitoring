package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

type KVStore interface {
	Read(string) (string, bool)
	Write(string, []byte) bool
	PrintDebug()
}

var currentKVStore KVStore = nil

func GetKVStore() KVStore {

	if currentKVStore == nil {

		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {

			// We don't have a redis url in the environment
			// so we assume we are in a local environment
			// instead of redis we will just use a Hash
			l := NewLocalStore()
			currentKVStore = l

		} else {

			opt, err := redis.ParseURL("redis://:qwerty@localhost:6379/1")
			if err != nil {
				log.Print("Can't parse to the redis URL")
			}

			client := redis.NewClient(opt)
			_, errPing := client.Ping().Result()

			if errPing != nil {
				// not cool
				log.Print("Can't connect to the redis URL")
			}

			r := RedisStore{}
			r.redis = client
			currentKVStore = r
		}

	}
	return currentKVStore

}

type LocalStore struct {
	dic map[string]string
}

func NewLocalStore() LocalStore {

	l := LocalStore{}
	l.dic = make(map[string]string)
	return l
}

func (l LocalStore) Read(k string) (string, bool) {

	r, ok := l.dic[k]
	return r, ok

}

func (l LocalStore) Write(k string, v []byte) bool {

	l.dic[k] = string(v[:])
	return true
}

func (l LocalStore) PrintDebug() {

	log.Println("================")
	for k, v := range l.dic {
		log.Println(k + " | " + v)
	}
}

type RedisStore struct {
	redis *redis.Client
}

func (r RedisStore) Read(k string) (string, bool) {
	val, err := r.redis.Get(k).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (r RedisStore) Write(k string, v []byte) bool {

	err := r.redis.Set(k, v, 0).Err()
	if err != nil {
		return false
	}
	return true
}

func (r RedisStore) PrintDebug() {

}
