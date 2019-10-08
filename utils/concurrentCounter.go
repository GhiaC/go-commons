package utils

import "sync"

type ConcurrentCounter struct {
	sync.Mutex
	Count int
}

func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{
		Count: 0,
	}
}

func (u *ConcurrentCounter) Increase() {
	u.Lock()
	defer u.Unlock()
	u.Count += 1
}

func (u *ConcurrentCounter) Get() int {
	u.Lock()
	defer u.Unlock()
	return u.Count
}
