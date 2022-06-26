package main

import (
	"Q2Bank/src/api/config"
	"Q2Bank/src/rest"
	"Q2Bank/src/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/transaction", rest.TransactionHandler)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, "Conectado no servidor!")
	})

	http.Handle("/", router)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)

	}
	config.Carregar()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
