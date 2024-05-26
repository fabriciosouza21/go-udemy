package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

// TestIndex verifica se a função Index está correta
// t.Logf é usado para imprimir a mensagem de log
// go test -v - para executar os teste do terminal com verbose
func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct {
		texto    string
		part     string
		esperado int
	}{
		{"Curso de Go", "Go", 9},
		{"", "", 0},
		{"Curso de Go", "go", -1},
		{"fabricio", "f", 0},
	}

	for _, teste := range testes {
		t.Logf("Massa de teste: %v", teste)
		atual := strings.Index(teste.texto, teste.part)
		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.part, teste.esperado, atual)
		}
	}

}
