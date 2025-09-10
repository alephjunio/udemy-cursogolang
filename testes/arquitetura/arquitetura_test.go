package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	t.Parallel() // para que  teste rode de forma paralela com todos os demais que estão com essa assinatura
	if runtime.GOARCH == "amd64" {
		t.Skip("Teste não é necessário em arquitetura amd64")
	}
	t.Fail()
}
