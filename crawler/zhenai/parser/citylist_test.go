package parser

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//fmt.Println(string(contents))
	contents, err := os.ReadFile("citylist_test_data.html")
	if err != nil {
		log.Panic(err)
	}
	result := ParseCityList(contents, "http://www.zhenai.com/zhenghun")
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	//expectedCities := []string{
	//	"阿坝",
	//	"阿克苏",
	//	"阿拉善盟",
	//}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d", resultSize, len(result.Requests))
	}

	fmt.Printf("%q\n", result.Items)

	//if len(result.Items) != resultSize {
	//	t.Errorf("item should have %d "+
	//		"item; but had %d", resultSize, len(result.Items))
	//}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}
	//for i, url := range expectedCities {
	//	if result.Items[i].Url != url {
	//		t.Errorf("expected item #%d: %s; but "+
	//			"was %s",
	//			i, url, result.Items[i])
	//	}
	//}
}
