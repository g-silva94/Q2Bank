package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Usuario struct {
	ID        int64
	Nome      string
	Sobrenome string
	Email     string
	CPFCNPJ   string
	Senha     string
	Saldo     float64
	Tipo      string
}

type Transacao struct {
	ID        int64
	Valor     float64
	IDOrigem  int64
	IDDestino int64
	DateTime  time.Time
}

func MakeTransaction(payer Usuario, payee Usuario, transaction Transacao) error {

	conn, err := conectionDB()
	if err != nil {
		return fmt.Errorf("teste1: %v", err)
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	queryTransaction := fmt.Sprintf("INSERT INTO Transaction (Valor, IDOrigem, IDDestino, DateTime) VALUES (%0.2f, %d, %d, NOW())", transaction.Valor, transaction.IDOrigem, transaction.IDDestino)

	resultTransaction, err := tx.Exec(queryTransaction)
	if err != nil {
		tx.Rollback()
		return err
	}
	rowTransaction, err := resultTransaction.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Nenhuma transacao inserida")
	}
	if rowTransaction == 0 {
		tx.Rollback()
		return fmt.Errorf("Nenhuma transacao inserida")
	}

	querySaque := fmt.Sprintf("UPDATE User set Saldo = (Saldo-%v) WHERE ID = %v", transaction.Valor, transaction.IDOrigem)

	resultSaque, err := tx.Exec(querySaque)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowSaque, err := resultSaque.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Nenhum pagador encontrado")
	}
	if rowSaque == 0 {
		tx.Rollback()
		return fmt.Errorf("Nenhum pagador encontrado")
	}

	queryDeposito := fmt.Sprintf("UPDATE User set Saldo = (Saldo+%v) WHERE ID = %v", transaction.Valor, transaction.IDDestino)

	resultDeposito, err := tx.Exec(queryDeposito)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowDeposito, err := resultDeposito.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Nenhum recebedor encontrado")
	}
	if rowDeposito == 0 {
		tx.Rollback()
		return fmt.Errorf("Nenhum recebedor encontrado")
	}

	return tx.Commit()
}

func GetUser(id int64) (Usuario, error) {

	var user Usuario

	conn, err := conectionDB()
	if err != nil {
		return user, fmt.Errorf("Caiu aqui: %v", err)
	}
	defer conn.Close()

	rows := conn.QueryRow("SELECT ID, Nome, Sobrenome, Email, CPFCNPJ, Senha, Saldo, Tipo FROM User WHERE ID=%v", id)
	rows.Scan(&user.ID, &user.Nome, &user.Sobrenome, &user.Email, &user.CPFCNPJ, &user.Senha, &user.Saldo, &user.Tipo)

	return user, nil

}

func conectionDB() (*sql.DB, error) {

	var (
		err error
	)

	// postgresConnection := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
	// 	os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"),
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB"),
	// )

	// "postgres", "host=localhost port=15432 user=postgres password=master dbname=postgresql sslmode=disable"

	db, err := sql.Open("postgres", "host=0.0.0.0 port=5432 user=admin password=admin dbname=q2Teste sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
