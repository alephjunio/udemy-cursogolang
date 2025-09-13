package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// connectDB establishes a connection to the MySQL database.
func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goapp")
	if err != nil {
		panic(err)
	}

	return db
}

// UserHandler handles HTTP requests for user resources.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path by removing "/users/" prefix
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	if err != nil {
		// If ID is not a valid integer, set it to 0
		id = 0
	}

	switch {
	case r.Method == "GET" && id > 0:
		getUser(w, r, id)
	case r.Method == "GET":
		getUsers(w, r)
	case r.Method == "POST":
		createUser(w, r)
	case r.Method == "PUT":
		updateUser(w, r)
	case r.Method == "DELETE":
		deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Recurso não encontrado")
	}
}

// getUser retrieves a user by ID from the database.
func getUser(w http.ResponseWriter, r *http.Request, id int) {
	db := connectDB()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE id = ?", id)

	var u user
	row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

// getUsers retrieves all users from the database.
func getUsers(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email, password, created_at FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
		users = append(users, u)
	}

	json, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

// createUser creates a new user in the database.
func createUser(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	var u user
	json.NewDecoder(r.Body).Decode(&u)

	stmt, err := db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Usuário criado com sucesso")
}

// updateUser updates an existing user in the database.
func updateUser(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	if err != nil {
		panic(err)
	}

	var u user
	json.NewDecoder(r.Body).Decode(&u)

	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Password, id)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Usuário atualizado com sucesso")
}

// deleteUser deletes a user by ID from the database.
func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Usuário deletado com sucesso")
}
