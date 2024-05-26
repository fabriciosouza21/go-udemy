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

	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id,nome) values($1, $2)")
	stmt.Exec(2000, "Bia")
	stmt.Exec(3000, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // chave duplicada

	// go não tratar o erro, então é necessário fazer o rollback
	// se não fizer o rollback, o erro será ignorado
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	errTx := tx.Commit()
	if errTx != nil {
		log.Fatal(err)
	}

}
