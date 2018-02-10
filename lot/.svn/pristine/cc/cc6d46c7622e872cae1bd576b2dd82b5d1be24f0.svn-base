package lutils

import "sync"

//一个带Close功能的WaitGroup
type WaitGroupPlus struct {
	wg    sync.WaitGroup
	count int
	sync.Mutex
}

func (this *WaitGroupPlus) Add(delta int) {
	this.wg.Add(delta)
	defer this.Unlock()
	this.Lock()
	this.count += delta
}
func (this *WaitGroupPlus) Wait() {
	this.wg.Wait()
}
func (this *WaitGroupPlus) Done() {
	this.wg.Done()
	defer this.Unlock()
	this.Lock()
	this.count --
}
func (this *WaitGroupPlus) Close() {
	this.Lock()
	defer this.Unlock()
	this.wg.Add(-this.count)
}
