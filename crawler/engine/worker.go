package engine

import (
	"PRO02/crawler/fetcher"
	"log"
)

// Worker ：Fetcher和 Parser合并
func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
