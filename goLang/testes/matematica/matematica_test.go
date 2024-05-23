package matematica

import "testing"

const erroPadrao = "valor esperado %v, mas o resultado encontrado foi %v"

// go test - para executar os teste do terminal
// go teste ./... - para executar todos os teste em subdiretorios
func TestMedia(t *testing.T) {
	valorEsparado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsparado {
		t.Errorf(erroPadrao, valorEsparado, valor)
	}
}
