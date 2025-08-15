package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=5698 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) // Encapsula todos os templates e retorna dois valores o template e o erro, template.ParseGlob() caminho de onde estão os templates, "templates/*.index.html", *.html, pega todos os arquivos em html

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index) //(qual requisição você quer atender, quem atenderá essa requisição)

	http.ListenAndServe(":8000", nil) //ouça e sirva, porta, não passa nem uma informação

}

func index(w http.ResponseWriter, r *http.Request) { // (Mostrar a resposta, apontando para o requerimento) Estrutura padrão
	produtos := []Produto{
		{"Camiseta", "Azul", 25, 2},
		{"Notebook", "O melhor do mercado", 2500, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos) //(Quem consegue escrever, quem ele vai exibir, quer passar alguma informação)

}
