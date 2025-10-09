package users

import "sync"

type CacheStore map[string]User
type Conn struct {
	cache CacheStore
	mu    *sync.Mutex
}

func NewConn() *Conn {
	return &Conn{
		cache: make(CacheStore, 100),
		mu:    new(sync.Mutex),
	}
}

func CreateUser() {

}
