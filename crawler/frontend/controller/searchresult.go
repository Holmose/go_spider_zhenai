package controller

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/frontend/model"
	"PRO02/crawler/frontend/view"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// TODO
// support paging
// and start page

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.160.142:9200"),
		elastic.SetSniff(false))
	if err != nil {
		log.Panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSarchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男士&from=20
func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
			rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from

	for _, v := range resp.Each(
		reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		result.Items = append(result.Items, item)
	}
	result.PrevFrom =
		result.Start - len(result.Items)
	result.NextFrom =
		result.Start + len(result.Items)
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*:)`)
	return re.ReplaceAllString(q, "Payload.$1")
}
