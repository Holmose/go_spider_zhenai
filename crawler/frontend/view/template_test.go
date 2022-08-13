package view

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/frontend/model"
	common "PRO02/crawler/model"
	"os"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	template := template.Must(
		template.ParseFiles("template.html"),
	)
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1141336149",
		Type: "zhenai",
		Id:   "1141336149",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	out, err := os.Create("template_test.html")
	err = template.Execute(out, page)
	if err != nil {
		t.Error(err)
	}

}
