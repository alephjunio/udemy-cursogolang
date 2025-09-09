package matematica

import (
	"testing"
)

const defaultError = "Valor recebido %g esse foi o valor esperado %g."

func TestMedia(t *testing.T) {
	valorEsperado := 4.0
	media := Media(2.0, 4.0, 6.0)
	if valorEsperado != media {
		t.Errorf(defaultError, media, valorEsperado)
	}
}
