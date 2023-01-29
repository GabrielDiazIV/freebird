package domain

import (
	"Freebird/app/datastore/data"
	"sync"
)

type UserPool struct {
	mu    sync.RWMutex
	users map[chan<- *data.Bird]bool
}

func NewPool() *UserPool {
	return &UserPool{
		mu:    sync.RWMutex{},
		users: make(map[chan<- *data.Bird]bool),
	}

}

func (u *UserPool) UserChans() []chan<- *data.Bird {
	u.mu.RLock()
	defer u.mu.Unlock()

	userChans := make([]chan<- *data.Bird, len(u.users))

	i := 0
	for k := range u.users {
		userChans[i] = k
		i++
	}

	return userChans
}

func (u *UserPool) Subscibe(userChan chan<- *data.Bird) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.users[userChan] = true
}

func (u *UserPool) Remove(userChan chan<- *data.Bird) {
	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.users, userChan)
}
