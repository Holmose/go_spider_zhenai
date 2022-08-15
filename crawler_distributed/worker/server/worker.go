package main

import (
	"PRO02/crawler_distributed/rpcsupport"
	"PRO02/crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Worker Server must specify a port")
		return
	}
	log.Printf("Starting Worker Server ...")
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
