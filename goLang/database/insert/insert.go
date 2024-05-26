package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=db_go sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Prepare statement simila o que se tem no java
	// para o mysql é ? e para o postgres é $1
	stmt, errPrep := db.Prepare(" INSERT INTO usuarios(nome) values($1)")
	if errPrep != nil {
		log.Fatal(errPrep)
	}

	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
