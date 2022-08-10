package parser

import (
	"PRO02/crawler/engine"
	"fmt"
	"regexp"
)

const profileRe = "[\u4e00-\u9fa5]+"

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(profileRe)
	fmt.Println(string(contents))
	matches := re.FindAllString(string(contents), -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		fmt.Print("profile: ", match)
	}
	return result
}
