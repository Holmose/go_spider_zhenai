package main

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/scheduler"
	"PRO02/crawler/zhenai/parser"
	"PRO02/crawler_distributed/config"
	itemsaver "PRO02/crawler_distributed/persist/client"
	worker "PRO02/crawler_distributed/worker/client"
	"fmt"
	"log"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		log.Panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		log.Panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      5,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			"ParseCityList"),
	})

}
