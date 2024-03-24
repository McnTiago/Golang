package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Parceiros struct {
	ID       string    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Empresa  string    `json:"empresa,omitempty"`
	Endereco *Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Cidade string `json:"cidade,omitempty"`
	Estado string `json:"estado,omitempty"`
}

var parceiro []Parceiros

func GetParceiros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(parceiro)
}

func GetParceiro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range parceiro {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Parceiros{})
}

func CreateParceiros(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var parceiros Parceiros
	_ = json.NewDecoder(r.Body).Decode(&parceiros)
	parceiros.ID = params["id"]
	parceiro = append(parceiro, parceiros)
	json.NewEncoder(w).Encode(parceiro)
}

func DeleteParceiros(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range parceiro {
		if item.ID == params["id"] {
			parceiro = append(parceiro[:index], parceiro[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(parceiro)
	}
}

func main() {
	router := mux.NewRouter()
	parceiro = append(parceiro, Parceiros{ID: "1", Nome: "Maicon", Empresa: "AçaiMix", Endereco: &Endereco{Cidade: "Lambari", Estado: "MG"}})
	parceiro = append(parceiro, Parceiros{ID: "2", Nome: "Luiz", Empresa: "Quasar", Endereco: &Endereco{Cidade: "Atibaia", Estado: "SP"}})
	router.HandleFunc("/contato", GetParceiros).Methods("GET")
	router.HandleFunc("/contato/{id}", GetParceiro).Methods("GET")
	router.HandleFunc("contato/{id}", DeleteParceiros).Methods("DELETE")
	router.HandleFunc("contato/{id}", CreateParceiros).Methods("POST")
	log.Fatal(http.ListenAndServe(":3100", router))
}
