package main

import "net/http"

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Hello World!"))
	//})
	mux.HandleFunc("/", HomeHander)
	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello man!"))
	})

	http.ListenAndServe(":8082", mux2)
}

func HomeHander(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(b.title))
}
