package users

import (
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CacheStore map[string]User
type Conn struct {
	cache CacheStore
	mu    *sync.RWMutex
}

func NewConn() *Conn {
	return &Conn{
		cache: make(CacheStore, 100),
		mu:    new(sync.RWMutex),
	}
}

func (c *Conn) CreateUser(n NewUser) (User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	u := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(hashedPass),
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[u.Email] = u

	return u, nil

}
