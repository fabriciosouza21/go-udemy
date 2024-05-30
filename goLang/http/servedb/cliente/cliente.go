package cliente

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// struct com tags para json
// os nomes dos campos devem começar com letra maiúscula
// para serem exportados
// caso contrário, não serão exportados
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler lida com requisições para o recurso de usuários
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")

	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuariosTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("Método não suportado")
	}

}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=db_go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var u Usuario
	// QueryRow retorna apenas uma linha
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = $1", id).Scan(&u.ID, &u.Nome)

	// retorna json em bytes
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuariosTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=db_go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, errRow := db.Query("SELECT id, nome FROM usuarios")

	if errRow != nil {
		log.Fatal(errRow)
	}

	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, u)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
