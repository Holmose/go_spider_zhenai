package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	// WorkerChan 询问Scheduler 给自己那个channel, 是chan request 还是chan chan request
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run 指针类型的接收者，用于改变 Scheduler
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	// 将通道传入Scheduler
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//itemCount := 0
	for {
		// 获取输出通道的内容
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got item #%d: %#v", itemCount, item)
			//itemCount++
			if isDuplicate(item.Url) {
				continue
			}
			func() { e.ItemChan <- item }()
		}

		// 将request 送入Scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

// 去重 空间换取时间
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
