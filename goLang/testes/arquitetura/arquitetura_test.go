package arquitetura

import (
	"runtime"
	"testing"
)

// go test -v  para ver o logger
func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	t.Fail()
}
