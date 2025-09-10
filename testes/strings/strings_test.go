package strings

import (
	"strings"
	"testing"
)

const messageError = "%s (parte: %s)  - Indices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	expected := []struct {
		text   string
		part   string
		indice int
	}{
		{"Hello, World!", "World", 7},
		{"Hello, World!", "Hello", 0},
		{"Hello, World!", "!", 12},
	}

	for _, test := range expected {
		t.Logf("Teste: %s, Parte: %s, Indice: %d", test.text, test.part, test.indice)
		actual := strings.Index(test.text, test.part)
		if test.indice != actual {
			t.Errorf(messageError, "Index", test.part, test.indice, actual)
		}
	}

}
