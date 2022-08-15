package main

import (
	"PRO02/crawler_distributed/config"
	"PRO02/crawler_distributed/rpcsupport"
	"PRO02/crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Print("Starting Worker Server ... ")
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}
