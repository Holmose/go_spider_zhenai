package parser

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

var (
	profileRe = regexp.MustCompile("<a href=.(http://album.zhenai.com/u/[\\d]+).[^>]+>([^<]+)</a>" +
		"[^\u6027\u522b]+>性别：</span>([^<]+)</td>" +
		"[^\u5c45\u4f4f\u5730]+>居住地：</span>([\u4e00-\u9fa5]+)</td>" +
		"[^\u5e74\u9f84]+>年龄：</span>([\\d]+)</td>" +
		"[^\u6708\u85aa|\u5b66\u5386]+>[月|学].*?[薪|历]：</span>([^<]+)</td>" +
		"[^\u5a5a\u51b5]+>婚况：</span>([\u4e00-\u9fa5]+)</td>" +
		"[^\u8eab\u9ad8]+>身.*?高：</span>([\\d]+)</td>" +
		"[^=]*?class=.introduce.>([^<]+)</div>")

	cityUrlRe = regexp.MustCompile(`<a.+?href="(http://www.zhenai.com/zhenghun/[^"]+?)">[^<]+</a>`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		gender := string(m[3])
		place := string(m[4])
		age, err := strconv.Atoi(string(m[5]))
		if err != nil {
			age = -1
		}
		salaryoredu := string(m[6])
		marry := string(m[7])
		height, err := strconv.Atoi(string(m[8]))
		if err != nil {
			height = -1
		}
		introduce := string(m[9])

		item := model.Profile{
			Name:      name,
			Gender:    gender,
			Age:       age,
			Height:    height,
			Marriage:  marry,
			Place:     place,
			Introduce: introduce,
		}

		if strings.Contains(salaryoredu, "元") {
			item.Income = salaryoredu
		} else {
			item.Education = salaryoredu
		}

		result.Items = append(
			result.Items, item)

		//result.Requests = append(
		//	result.Requests, engine.Requests{
		//		Url:        string(m[1]),
		//		ParserFunc: ParseProfile,
		//	})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result

}
