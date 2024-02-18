package main

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

type Lock struct {
	key        string
	value      string
	expiration time.Duration
	mux        sync.Mutex
	isLocked   bool
	count      int
}

func NewLock(key string, expiration time.Duration) *Lock {
	return &Lock{
		key:        key,
		value:      uuid.New().String(),
		expiration: expiration,
	}
}

func (l *Lock) Lock() (bool, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if l.isLocked {
		l.count++
		return true, nil
	}

	ok, err := client.SetNX(l.key, l.value, l.expiration).Result()
	if err != nil || !ok {
		return false, err
	}

	l.isLocked = true
	l.count++

	go l.renew()
	return true, nil
}

func (l *Lock) Unlock() error {
	l.mux.Lock()
	defer l.mux.Unlock()

	if !l.isLocked {
		return nil
	}

	l.count--

	if l.count > 0 {
		return nil
	}

	val, err := client.Get(l.key).Result()
	if err != nil {
		return err
	}

	if val != l.value {
		return nil
	}

	l.isLocked = false
	return client.Del(l.key).Err()
}

func (l *Lock) renew() {
	ticker := time.NewTicker(l.expiration / 2)
	for range ticker.C {
		l.mux.Lock()

		if !l.isLocked {
			ticker.Stop()
			l.mux.Unlock()
			break
		}

		client.Expire(l.key, l.expiration)
		l.mux.Unlock()
	}
}

func main() {
	lock := NewLock("mykey", 10*time.Second)
	locked, err := lock.Lock()
	if err != nil {
		panic(err)
	}

	if !locked {
		return
	}

	defer lock.Unlock()
}
