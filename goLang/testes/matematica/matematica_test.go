package matematica

import "testing"

const erroPadrao = "valor esperado %v, mas o resultado encontrado foi %v"

// go test - para executar os teste do terminal
// go teste ./... - para executar todos os teste em subdiretorios
// convesão de nomeclatura de teste: Test + nome da função que será testada
// convesão de numeclatrua do arquivo de teste: nome do arquivo + _test.go
// *testing.T é um ponteiro de utilitário de teste
func TestMedia(t *testing.T) {
	valorEsparado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsparado {
		t.Errorf(erroPadrao, valorEsparado, valor)
	}
}
