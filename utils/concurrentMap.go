package utils

import "sync"

// ConcurrentStringMap is a concurrent storage for strings (emails, links etc.)
// that automatically removes duplicates. Very helpful for verifying output
// of Newsletter recipients or URL links in Newsletter HTML/Text.
type ConcurrentStringMap struct {
	sync.Mutex
	strings map[string]int
}

func NewConcurrentStringMap() *ConcurrentStringMap {
	return &ConcurrentStringMap{
		strings: map[string]int{},
	}
}

func (u *ConcurrentStringMap) Add(str string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.strings[str]; ok {
		u.strings[str] += 1
		return false
	}
	u.strings[str] = 0 // panic happens here
	return true
}

func (u *ConcurrentStringMap) Count() int {
	return len(u.strings)
}

func (u *ConcurrentStringMap) Get(str string) int {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.strings[str]; ok {
		return u.strings[str]
	}
	return 0
}

func (u *ConcurrentStringMap) Slice() []string {
	u.Lock()
	defer u.Unlock()
	slice := make([]string, 0, len(u.strings))
	for str, _ := range u.strings {
		slice = append(slice, str)
	}
	return slice
}
