package rest

import (
	"Q2Bank/src/service"
	"Q2Bank/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Metodo nao permitido!")
		return
	}
	var transac service.TransactionRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao ler o Body")
		return
	}

	if err := json.Unmarshal([]byte(body), &transac); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao executar Unmarshall do JSON")
		return
	}
	err = service.MakeTransaction(transac)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, err.Error())
		return
	}

	fmt.Printf("Transação feita com sucesso!")
	utils.RespondWithJSON(w, http.StatusOK, transac)
	return

}
