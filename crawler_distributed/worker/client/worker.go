package client

import (
	"PRO02/crawler/engine"
	"PRO02/crawler_distributed/config"
	"PRO02/crawler_distributed/rpcsupport"
	"PRO02/crawler_distributed/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(
		req engine.Request) (
		engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err = client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil

	}, nil
}
