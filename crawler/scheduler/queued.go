package scheduler

import (
	"PRO02/crawler/engine"
	"log"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	// 每个work都有自己的channel
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	// worker的channel准备好了，送入request
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	log.Panic("Implement interface", c)
}
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			// 多个事情谁的channel中有数据，先做谁
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
