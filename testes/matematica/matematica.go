package matematica

import (
	"fmt"
	"strconv"
)

func Media(valores ...float64) float64 {
	total := 0.0
	for _, valor := range valores {
		total += valor
	}

	media := total / float64(len(valores))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
