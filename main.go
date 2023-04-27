package main

import (
	"mymodule/Models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("Templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := Models.SearchProdutos()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
