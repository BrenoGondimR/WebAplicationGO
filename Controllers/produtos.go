package Controllers

import (
	"log"
	"mymodule/Models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("Templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := Models.SearchProdutos()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvert, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco")
		}
		quantidadeConvert, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade")
		}
		Models.CriarNewProducts(nome, descricao, precoConvert, quantidadeConvert)
		http.Redirect(w, r, "/", 301)
	}
}
