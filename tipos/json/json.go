package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Em Go, interfaces são tipos abstratos que definem um conjunto de métodos.
// Uma interface define um comportamento esperado, mas não sua implementação.

// Exemplos de uso de interfaces:
// 1. Uma interface 'Writer' pode ter um método 'Write()'
// 2. Qualquer tipo que implemente esse método automaticamente satisfaz a interface
// 3. Não é necessário declarar explicitamente que um tipo implementa uma interface

// As interfaces em Go são implementadas implicitamente
// Se um tipo possui todos os métodos que uma interface especifica,
// ele automaticamente implementa essa interface

// Neste exemplo, a struct produto poderia implementar interfaces como:
// - json.Marshaler para personalizar a serialização para JSON
// - fmt.Stringer para personalizar como o objeto é convertido para string

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// json.Marshal e json.Unmarshal usam interfaces internamente
	// para determinar como serializar/deserializar os objetos

	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 2500,
		Tags:  []string{"Promoção", "Eletrônico"},
	}
	p2 := produto{
		ID:    2,
		Nome:  "Mouse",
		Preco: 50,
		Tags:  []string{"Promoção", "Eletrônico"},
	}
	produtos := []produto{p1, p2}
	jsonItens, err := json.Marshal(produtos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonItens))

	var p3 produto
	jsonString := ` {"id":3,"nome":"Tablet","preco":1200,"Tags":["Promoção","Eletrônico", "Importado"]}`
	err = json.Unmarshal([]byte(jsonString), &p3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p3)
}
