package main

import (
	"PRO02/crawler_distributed/config"
	"PRO02/crawler_distributed/persist"
	"PRO02/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("ItemSaver Server must specify a port")
		return
	}
	log.Print("Starting ItemSaver Server ...")
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.160.142:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
