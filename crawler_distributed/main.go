package main

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/scheduler"
	"PRO02/crawler/zhenai/parser"
	"PRO02/crawler_distributed/config"
	"PRO02/crawler_distributed/persist/client"
	"fmt"
	"log"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		log.Panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
