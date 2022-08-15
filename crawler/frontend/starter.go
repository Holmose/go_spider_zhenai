package main

import (
	"PRO02/crawler/frontend/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view/")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"crawler/frontend/view/template.html"))
	fmt.Println("Frontend Server Listening 8888 ...")
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Panic(err)
	}

}
