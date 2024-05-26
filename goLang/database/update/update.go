package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=db_go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// sempre fechar a conexão
	// verifique se há erro antes de tentar fechar

	defer db.Close()

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = $1 WHERE id = $2")
	stmt.Exec("Uógiton", 1)
	stmt.Exec("valesca", 2)

	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = $1")
	stmt2.Exec(3)
}
