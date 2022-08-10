package parser

import (
	"PRO02/crawler/engine"
	"regexp"
)

const cityRe = "<a href=.(http://album.zhenai.com/u/[\\d]+).[^>]+>([^<]+)</a>" +
	"[^\u6027\u522b]+>性别：</span>([^<]+)</td>" +
	"[^\u5c45\u4f4f\u5730]+>居住地：</span>([\u4e00-\u9fa5]+)</td>" +
	"[^\u5e74\u9f84]+>年龄：</span>([\\d]+)</td>" +
	"[^\u6708\u85aa|\u5b66\u5386]+>[月|学].*?[薪|历]：</span>([^<]+)</td>" +
	"[^\u5a5a\u51b5]+>婚况：</span>([\u4e00-\u9fa5]+)</td>" +
	"[^\u8eab\u9ad8]+>身.*?高：</span>([\\d]+)</td>" +
	"[^=]*?class=.introduce.>([^<]+)</div>"

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		gender := string(m[3])
		place := string(m[4])
		age := string(m[5])
		salaryoredu := string(m[6])
		marry := string(m[7])
		height := string(m[8])
		introduce := string(m[9])

		items := map[string]string{
			"User":        name,
			"gender":      gender,
			"place":       place,
			"age":         age,
			"salaryoredu": salaryoredu,
			"marry":       marry,
			"height":      height,
			"introduce":   introduce,
		}

		result.Items = append(
			result.Items, items)

		//result.Request = append(
		//	result.Request, engine.Request{
		//		Url:        string(m[1]),
		//		ParserFunc: ParseProfile,
		//	})
	}

	return result

}
