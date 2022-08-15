package main

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/model"
	"PRO02/crawler_distributed/config"
	"PRO02/crawler_distributed/rpcsupport"
	"log"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		log.Panic(err)
	}

	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1141336149",
		Type: "zhenai",
		Id:   "1141336149",
		Payload: model.Profile{
			Name:       "Dingsir_",
			Gender:     "男士",
			Age:        31,
			Height:     174,
			Income:     "8001-12000元",
			Marriage:   "未婚",
			Education:  "",
			Occupation: "",
			Place:      "新疆阿克苏",
			Introduce:  "我是一个比较简单的人，没太多的心眼和套路，我希望我的另一半也是个简单点的人。",
			House:      "",
			Car:        ""},
	}

	// Call Save
	result := ""
	err = client.Call(
		config.ItemSaverRpc,
		item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}

}
