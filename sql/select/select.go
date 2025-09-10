package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id       int
	name     string
	email    string
	password string
}

func main() {
	// Conectando ao banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goapp")
	if err != nil {
		panic(err)
	}
	// Fechando a conexão com o banco de dados
	defer db.Close()

	// Executando a consulta
	rows, err := db.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterando sobre os resultados
	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name, &u.email, &u.password)
		fmt.Println(u)
	}

}
