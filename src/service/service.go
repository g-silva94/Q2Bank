package service

import (
	"Q2Bank/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TransactionRequest struct {
	Value float64 `json:"value"`
	Payer int64   `json:"payer"`
	Payee int64   `json:"payee"`
}

type MockJSON struct {
	Authorization bool `json:"authorization"`
}

func MakeTransaction(req TransactionRequest) error {

	var mock MockJSON
	urlMock := "https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31"

	resp, err := http.Get(urlMock)
	if err != nil {
		log.Printf("Erro ao realizar requisicao %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Falha ao realizar leitura do Body")
		return err
	}

	if err := json.Unmarshal(body, &mock); err != nil {
		log.Printf("Falha ao realizar Unmarshal do Body")
		return err
	}

	if !mock.Authorization {
		log.Printf("Você não está autorizado a realizar essa transacao")
		return fmt.Errorf("Você não está autorizado a realizar essa transacao")
	}

	payer, err := repository.GetUser(req.Payer)
	if err != nil {
		log.Printf("Falha ao buscar Payer %v", err)
		return err
	}

	payee, err := repository.GetUser(req.Payee)
	if err != nil {
		log.Printf("Falha ao buscar Payee %v", err)
		return err
	}

	if payer.Tipo == "lojista" {
		return fmt.Errorf("Lojista não está autorizado a realizar uma transferência.")
	}

	if payer.Saldo < req.Value {
		return fmt.Errorf("Você não possui saldo suficiente!")
	}

	transac := repository.Transacao{
		Valor:     req.Value,
		IDOrigem:  payer.ID,
		IDDestino: payee.ID,
		DateTime:  time.Now(),
	}

	err = repository.MakeTransaction(payer, payee, transac)
	if err != nil {
		return err
	}

	return nil
}
