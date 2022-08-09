package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		log.Panic("起始页打开失败！")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panic("Error: status code", resp.Status)
	}
	// 推测编码
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := io.ReadAll(utf8Reader)
	if err != nil {
		log.Panic("获取页面数据失败！")
	}
	// 解码Unicode
	resp_str, _ := zhToUnicode([]byte(all))
	//fmt.Println(string(resp_str))
	printCityList(resp_str)

}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("City: %s, URL: %s\n", match[2], match[1])
	}
	fmt.Println("Matches found: %d\n", len(matches))

}

// Unicode解码
func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
