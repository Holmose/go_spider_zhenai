package main

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/scheduler"
	"PRO02/crawler/zhenai/parser"
	itemsaver "PRO02/crawler_distributed/persist/client"
	"PRO02/crawler_distributed/rpcsupport"
	worker "PRO02/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		log.Panic(err)
	}

	pool := createClientPool(
		strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)
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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
