package main

import "encoding/json"

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)

	// retorna um slice de bytes
	// converte o slice para string
	println(string(p1Json))

	// json to struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	// converte a string para slice de bytes
	// passar a referência do objeto para a função
	// é retornado a struct preenchida
	json.Unmarshal([]byte(jsonString), &p2)

	println(p2.Tags[1])
}
