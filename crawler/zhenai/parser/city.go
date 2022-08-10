package parser

import (
	"PRO02/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)"[^>]+>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(
			result.Items, "User "+name)
		result.Request = append(
			result.Request, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseProfile,
			})
	}
	return result

}
