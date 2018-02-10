package global

import "sync"

//定时任务的开关
type Switch struct {
	state    int64 //1开启2结束
	stopChan chan interface{}
	sync.Mutex
}

func NewSwitch() *Switch {
	return &Switch{state: 2, stopChan: make(chan interface{})}
}

func (m *Switch) IsOpen() int64 {
	m.Lock()
	defer m.Unlock()
	return m.state
}

//开启定时任务,并传入任务的方法
func (m *Switch) Open(work func(stopChan chan interface{})) {
	if m.state == 2 {
		m.Lock()
		defer m.Unlock()
		if m.state == 2 {
			m.state = 1
			if m.stopChan == nil {
				panic("error:struct Switch.stopChan not init")
			}
			work(m.stopChan)
		}
	}
}

//关闭定时任务
func (m *Switch) Close() {
	if m.state == 1 {
		m.Lock()
		defer m.Unlock()
		if m.state == 1 {
			m.state = 2
			if m.stopChan != nil {
				m.stopChan <- 0
			}
		}
	}
}
