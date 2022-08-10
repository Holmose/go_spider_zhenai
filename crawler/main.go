package main

import (
	"PRO02/crawler/engine"
	"PRO02/crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/aba",
	//	ParserFunc: parser.ParseCity,
	//})

}
