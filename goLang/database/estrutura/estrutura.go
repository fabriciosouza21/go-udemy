package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func createDatabaseIfNotExist(db *sql.DB, dbName string) error {
	// Verifica se o banco de dados existe
	var name string
	query := "SELECT datname FROM pg_catalog.pg_database WHERE datname = $1"
	err := db.QueryRow(query, dbName).Scan(&name)
	if err == sql.ErrNoRows {
		// O banco de dados não existe, então cria
		_, err := db.Exec("CREATE DATABASE " + dbName) // Uso simples, mas cuidado com injeção de SQL
		if err != nil {
			return err
		}
		log.Println("Database created")
	} else if err != nil {
		return err
	} else {
		log.Println("Database already exists")
	}
	return nil
}

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
	createDatabaseIfNotExist(db, "db_go")
	exec(db, "drop table if exists usuarios")
	// não está criando a tabela
	// se precisar fabricio do futuro descubra o motivo
	exec(db, `CREATE TABLE usuarios (
		id serial primary key,
		nome varchar(80)
	)`)
}
