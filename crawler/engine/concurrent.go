package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

// Run 指针类型的接收者，用于改变 Scheduler
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	// 将通道传入Scheduler
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		// 获取输出通道的内容
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		// 将request 送入Scheduler
		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(out chan ParseResult, s Scheduler) {
	// 每个Worker自己的channel
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
