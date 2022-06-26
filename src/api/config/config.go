package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PostgresConnection = ""
)

// Carregar() vai inicializar as variáveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	Port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		Port = 9000
	}

	PostgresConnection = fmt.Sprintf("host=localhost port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		Port,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

}
