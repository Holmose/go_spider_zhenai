package model

import "PRO02/crawler/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
