package global

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
func (this *SafeErr) Error() (temp string) {
	this.RLock()
	temp = this.str
	this.RUnlock()
	return
}
func (this *SafeErr) IsValid() bool {
	var temp string
	this.RLock()
	temp = this.str
	this.RUnlock()
	return temp != ""
}

func (this *SafeErr) Push(str string) {
	this.Lock()
	defer this.Unlock()
	this.str = str
}
