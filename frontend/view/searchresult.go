package view

import (
	"PRO02/frontend/model"
	"html/template"
	"io"
	"strings"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSarchResultView(filename string) SearchResultView {
	add := func(left int, right int) int {
		return left + right
	}
	filenames := strings.Split(filename, "/")
	t := template.New(filenames[len(filenames)-1])
	// 自定义函数必须在解析模板之前
	t.Funcs(template.FuncMap{
		"add": add,
	})
	return SearchResultView{
		template: template.Must(
			t.ParseFiles(filename))}
}

func (s SearchResultView) Render(
	w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
