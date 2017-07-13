package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

// KVStore stand for Key Value Store, it's an interface above the persitance engine
// (here redis).
type KVStore interface {
	Read(string) (string, bool)
	Write(string, []byte) bool
	PrintDebug()
}

// currentKVStore is the store singleton
var currentKVStore KVStore

// GetKVStore is the factory function to access the store
// It's kind of cool : if you are running the app in local
// (ie : you don't have redis installed on your developpement
// machine) it's going to use a dictionnary (in memory) as
// KeyValue Store to mock the redis and allow you to debug
func GetKVStore() KVStore {

	if currentKVStore == nil {

		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {

			// We don't have a redis url in the environment
			// so we assume we are in a local environment
			// instead of redis we will just use a Hash
			l := newLocalStore()
			currentKVStore = l

		} else {

			opt, err := redis.ParseURL(redisURL)
			if err != nil {
				log.Print("Can't parse to the redis URL")
			}

			client := redis.NewClient(opt)
			_, errPing := client.Ping().Result()

			if errPing != nil {
				// not cool
				log.Print("Can't connect to the redis URL")
			}

			r := redisStore{}
			r.redis = client
			currentKVStore = r
		}

	}
	return currentKVStore

}

type localStore struct {
	dic map[string]string
}

func newLocalStore() localStore {

	l := localStore{}
	l.dic = make(map[string]string)
	return l
}

func (l localStore) Read(k string) (string, bool) {

	r, ok := l.dic[k]
	return r, ok

}

func (l localStore) Write(k string, v []byte) bool {

	l.dic[k] = string(v[:])
	return true
}

func (l localStore) PrintDebug() {

	log.Println("================")
	for k, v := range l.dic {
		log.Println(k + " | " + v)
	}
}

type redisStore struct {
	redis *redis.Client
}

func (r redisStore) Read(k string) (string, bool) {
	val, err := r.redis.Get(k).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (r redisStore) Write(k string, v []byte) bool {

	err := r.redis.Set(k, v, 0).Err()
	if err != nil {
		return false
	}
	return true
}

func (r redisStore) PrintDebug() {

}
