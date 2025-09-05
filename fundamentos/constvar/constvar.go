package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float64 = 3.14159265358979323846
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	println(e, f)

	g, h, i := 2, false, "epa!"
	println(g, h, i)

}
