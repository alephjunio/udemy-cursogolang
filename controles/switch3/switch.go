package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case string:
		return "String"
	case func():
		return "Função"
	case bool:
		return "Booleano"
	case float64, float32:
		return "Real"
	case complex64, complex128:
		return "Complexo"
	case []int:
		return "Slice de inteiros"
	case map[string]int:
		return "Mapa de strings para inteiros"
	case map[string]string:
		return "Mapa de strings para strings"
	case []string:
		return "Slice de strings"
	case nil:
		return "Nil"
	default:
		return "Tipo desconhecido"
	}
}

func main() {
	fmt.Println("Tipo de 10:", tipo(10))
	fmt.Println("String:", tipo("Hello"))
	fmt.Println("Função:", tipo(func() {}))
	fmt.Println("Booleano:", tipo(true))
	fmt.Println("Float:", tipo(3.14))
	fmt.Println("Complexo:", tipo(1+2i))
	fmt.Println("Slice de inteiros:", tipo([]int{1, 2, 3}))
	fmt.Println("Mapa de strings para inteiros:", tipo(map[string]int{"a": 1, "b": 2}))
	fmt.Println("Mapa de strings para strings:", tipo(map[string]string{"a": "A", "b": "B"}))
	fmt.Println("Slice de strings:", tipo([]string{"a", "b", "c"}))
	fmt.Println("Nil:", tipo(nil))
	fmt.Println("Tipo desconhecido:", tipo(struct{}{}))
}
