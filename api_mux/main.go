// This is a api with MUX
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("The server with mux is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter)) // Adicionando a porta 8080, caso não der certo para o programa com o Fatal

}

type User struct {
	ID   int    `json:"id"` // A string `` serve para mudar o nome da chave
	Name string `json:"name"`
}

func getUsers(res http.ResponseWriter, req *http.Request) {

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
