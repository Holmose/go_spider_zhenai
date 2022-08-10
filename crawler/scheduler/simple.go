package scheduler

import "PRO02/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	// 使用groutine解除循环等待
	go func() { s.workerChan <- r }()
}

// ConfigureMasterWorkerChan 指针类型的接收者，用于改变 workerChan的指
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
