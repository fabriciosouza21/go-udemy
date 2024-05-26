package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

// https://aprendagolang.com.br/como-conectar-e-fazer-crud-em-um-banco-postgresql/
func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	exec(db, "drop database if exists db_go")
	exec(db, "create database  db_go")
	exec(db, "drop table if exists usuarios")
	exec(db, `CREATE TABLE usuarios (
		id serial primary key,
		nome varchar(80)
	)`)
}
