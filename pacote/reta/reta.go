package main

import "math"

// Iniciando com letra maiúscula é público, visivel apenas dentro e fora do pacote
// Iniciando com letra minúscula é privado, visivel apenas dentro do pacote
// Exemplo
// Ponto - gerará algo publico
// ponto - gerará algo privado

// Ponto - Representa um ponto no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

// catetos - Calcula os catetos de um triângulo retângulo
// a - Ponto a
// b - Ponto b
// cx - Cateto x
// cy - Cateto y
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia - Calcula a distância entre dois pontos
// a - Ponto a
// b - Ponto b
// return - Distância entre os pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
