package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Inicia a transação
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goapp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepara a transação
	transacao, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// Prepara o statement
	stmt, _ := transacao.Prepare("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)")
	// Executa o statement
	stmt.Exec(2000, "Teste de Transação", "teste@example.com", "123456")
	stmt.Exec(2001, "Teste de Transação", "teste@example.com", "123456")
	// Executa o statement com erro
	_, err = stmt.Exec(1, "Gera erro de Rollback", "rollback@example.com", "123456") // Gera erro de Rollback
	// Se houver erro, rollback
	if err != nil {
		transacao.Rollback()
		log.Fatal(err)
	}

	// Se não houver erro, commit
	transacao.Commit()
}
