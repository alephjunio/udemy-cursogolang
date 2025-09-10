package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goapp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	stmt.Exec("Fulano", "fulano@example.com", "123456")
	stmt.Exec("Ciclano", "ciclano@example.com", "123456")
	stmt.Exec("Beltrano", "beltrano@example.com", "123456")

	res, _ := stmt.Exec("Alef", "alef@example.com", "123456")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)

}
