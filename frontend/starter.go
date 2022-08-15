package main

import (
	"PRO02/frontend/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("frontend/view/")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"frontend/view/template.tmpl"))
	fmt.Println("Frontend Server Listening 8888 ...")
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Panic(err)
	}

}
