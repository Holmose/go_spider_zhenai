package parser

import (
	"PRO02/crawler/engine"
	"fmt"
	"regexp"
)

const profileMoreRe = "[\u4e00-\u9fa5]+"

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	re := regexp.MustCompile(profileMoreRe)
	fmt.Println(string(contents))
	matches := re.FindAllString(string(contents), -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		fmt.Print("profile: ", match)
	}
	return result
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(
	contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (
	name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
