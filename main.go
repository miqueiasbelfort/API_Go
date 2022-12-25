package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers) // Chamando uma função na rota /users
	fmt.Println("The server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil)) // Adicionando a porta 8080, caso não der certo para o programa com o Fatal
}

type User struct {
	ID   int    `json:"id"` // A string `` serve para mudar o nome da chave
	Name string `json:"name"`
}

func getUsers(res http.ResponseWriter, req *http.Request) {

	// Se for usado um metodo diferente de GET retorn um status code de erro
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("Centent-Type", "application/json") //Configurando o header

	// Mandando um JSON no req
	json.NewEncoder(res).Encode([]User{
		{
			ID:   1,
			Name: "Miqueias Belfort",
		},
		{
			ID:   2,
			Name: "Amanda Silva",
		},
	})
}
