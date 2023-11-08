package main

import (
	"github.com/matzew/kodev/pkg/template"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", template.TemplateHandler)
	http.ListenAndServe(":8080", m)
}
