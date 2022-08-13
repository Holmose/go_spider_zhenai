package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %#v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out
}

// 使用客户端存储数据 https://github.com/olivere/elastic
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.160.142:9200"),
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
