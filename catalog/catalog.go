package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

type Products struct {
	Products []Product `json:"products"`
}

var productsURL string

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func main() {
	//instancia as rotas
	r := mux.NewRouter()

	//setando as rotas, quem acessar / ele ativa a funcao ListProducts
	r.HandleFunc("/", ListProducts)

	r.HandleFunc("/products/{id}", ShowProduct)

	//seta a porta
	http.ListenAndServe(":8080", r)

}

func loadProducts() []Product {
	response, err := http.Get(productsURL + "/products")
	if err != nil {
		fmt.Println("Erro de HTTP LoadProducts " + productsURL)
	}

	data, _ := ioutil.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)

	return products.Products
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	//carrega os produtos atraves da func especificada
	products := loadProducts()

	//carrega o tamplate
	t := template.Must(template.ParseFiles("templates/catalog.html"))

	//executa o template colocando as variaveis do Form products nele
	t.Execute(w, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	//Vars retorna as variáveis de rota para a solicitação atual
	vars := mux.Vars(r)

	//monta a url completa
	response, err := http.Get(productsURL + "/product/" + vars["id"])
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	//le todos os dados vindos da requisicao com a url montada acima
	data, _ := ioutil.ReadAll(response.Body)

	//instancia a variavel de tipo Product
	var product Product

	/*pega o []byte que esta salvo em data e tranforma em json, serializando-o utilizando a struct
	Product como referencia e salvando na variavel products obs: o & faz referencia ao endereco de memoria
	da variavel e modifica diretamente
	**/
	json.Unmarshal(data, &product)

	//instacia a variavel com o template salvo na pasta templates
	t := template.Must(template.ParseFiles("templates/view.html"))

	//executa o template colocando as variaveis do Form products nele
	t.Execute(w, product)
}
