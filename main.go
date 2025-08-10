package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) // Encapsula todos os templates e retorna dois valores o template e o erro, template.ParseGlob() caminho de onde estão os templates, "templates/*.index.html", *.html, pega todos os arquivos em html

func main() {
	http.HandleFunc("/", index) //(qual requisição você quer atender, quem atenderá essa requisição)

	http.ListenAndServe(":8000", nil) //ouça e sirva, porta, não passa nem uma informação

}

func index(w http.ResponseWriter, r *http.Request) { // (Mostrar a resposta, apontando para o requerimento) Estrutura padrão
	temp.ExecuteTemplate(w, "Index", nil) //(Quem consegue escrever, quem ele vai exibir, quer passar alguma informação)

}
