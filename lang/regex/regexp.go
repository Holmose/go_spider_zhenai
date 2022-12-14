package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com
email is abc@def.org
email2 is  kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([\w]+)@([\w]+)(\.[\w.]+)`)
	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}
}
