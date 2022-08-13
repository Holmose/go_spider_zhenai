package main

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/persist"
	"PRO02/crawler/scheduler"
	"PRO02/crawler/zhenai/parser"
	"log"
)

func main() {
	itemChan, err := persist.ItemSaver(
		"dating_profile")
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
