package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//estrutura para simular um form obs:o lang lida com orientacao a objetos dessa forma
type Product struct {

	/**
		Nessa estrutura estou serializando as variaveis com ocampo correspodente
		do JSON que é lido, ex: Uuid recebe o conteúdo da variavel de nome uuid do JSON
		obs: na variavel Price que é um float64 ela tenta ler uma string do json, o atributo string
		da expressao server para informar que a variavel vem como String e é pra converter pra float64
	**/
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

//Products é uma coleção da estrutura Product, semelhante a um array em java
type Products struct {
	//aparentemente no golang você informa o nome da variavel depois seu tipo
	Products []Product
}

func main() {
	//cria uma rota
	r := mux.NewRouter()
	//quem acesar o sub endereco /products vai habilitar a funcao que lista os produtos
	r.HandleFunc("/products", ListProducts)
	//quem acesar o sub endereco /product/"id do produto" vai habilitar a funcao que retorna o produto com aquele id
	r.HandleFunc("/product/{id}", getProductById)
	//essa funcao seta qual porta sera usada pelo servidor localhost
	http.ListenAndServe(":8081", r)

}

func loadData() []byte {

	jsonFile, err := os.Open("products.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()
	dados, err := ioutil.ReadAll(jsonFile)
	return dados
}

/**função pra listar os produtos, o atributo w é responsavel em informar qual sera a resposta da requisicao
	e a variave r é responsavel por receber a requisicao
**/
func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	//resposta em []byte
	w.Write([]byte(products))
}

/**função pra retornar o produto informado, o atributo w é responsavel em informar qual sera a resposta da requisicao
	e a variave r é responsavel por receber a requisicao
**/
func getProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//carrega os dados do json
	data := loadData()

	//declaracao da variavel de tipo Products(Array de Product)
	var products Products
	/*pega o []byte que esta salvo em data e tranforma em json, serializando-o utilizando a struct
	Product como referencia e salvando na variavel products obs: o & faz referencia ao endereco de memoria
	da variavel e modifica diretamente
	*/
	json.Unmarshal(data, &products)

	//laço para percorrer todos os produtos dentro de products
	for _, v := range products.Products {
		if v.Uuid == vars["id"] {
			//se ele achar o produto com o mesmo id ele salva em formato []byte na variavel product
			product, _ := json.Marshal(v)
			//escreve a resposta da requisicao
			w.Write([]byte(product))
		}

	}

}
