package lutils

import "sync"

type SafeErr struct {
	str string
	sync.RWMutex
}

func NewSafeError(strs ...string) *SafeErr {
	if len(strs) > 0 {
		return &SafeErr{str: strs[0]}
	}
	return &SafeErr{}
}
func (this *SafeErr) Error() string {
	this.RLock()
	defer this.RUnlock()
	return this.str
}
func (this *SafeErr) IsValid() bool {
	this.RLock()
	defer this.RUnlock()
	return this.str != ""
}

func (this *SafeErr) Push(str string) {
	this.Lock()
	defer this.Unlock()
	this.str = str
}