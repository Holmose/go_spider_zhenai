package persist

import (
	"PRO02/crawler/engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.160.142:9200"),
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %#v: %#v", itemCount, item, item.Payload)
			itemCount++
			err := save(client, item)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// 使用客户端存储数据 https://github.com/olivere/elastic
func save(client *elastic.Client, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
