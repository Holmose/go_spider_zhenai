package parser

import (
	"PRO02/crawler/engine"
	"fmt"
	"regexp"
)

const profileMoreRe = "[\u4e00-\u9fa5]+"

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(profileMoreRe)
	fmt.Println(string(contents))
	matches := re.FindAllString(string(contents), -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		fmt.Print("profile: ", match)
	}
	return result
}
