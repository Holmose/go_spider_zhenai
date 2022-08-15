package main

import (
	"PRO02/crawler_distributed/config"
	"PRO02/crawler_distributed/rpcsupport"
	"PRO02/crawler_distributed/worker"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host,
		worker.CrawlService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != err {
		log.Panic(err)
	}
	req := worker.Request{
		Url: "http://www.zhenai.com/zhenghun/aba",
		Parser: worker.SerializedParser{
			Name: config.ParseCity,
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
