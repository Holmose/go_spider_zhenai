package persist

import (
	"PRO02/crawler/model"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {

	expected := model.Profile{
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
		Car:        ""}

	id, err := save(expected)
	if err != nil {
		t.Error(err)
	}
	// 让测试不依赖于外界，代码启动一个docker go client
	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.160.142:9200"),
		elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
	t.Logf("%v", actual)

	// 删除所有数据
	/// POST 请求 192.168.160.142:9200/dating_profile/zhenai/_delete_by_query?pretty

	// 获取数据
	// GET 192.168.160.142:9200/dating_profile/zhenai/_search?q=Gender:女士 Age:(<30)&pretty&size=10
}
