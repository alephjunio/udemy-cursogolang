package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Conectando ao banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goapp")
	if err != nil {
		panic(err)
	}
	// Fechando a conexão com o banco de dados
	defer db.Close()

	// Atualizando os registros
	stmt, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	// Executando as atualizações
	stmt.Exec("Valesca", 1)
	stmt.Exec("Jhonsom", 2)
	stmt.Exec("Travis", 3)

	// Deletando os registros
	stmt, err = db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		panic(err)
	}

	// Executando as deleções
	stmt.Exec(3)

}
