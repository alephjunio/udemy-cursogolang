package main

import "fmt"

// Structs Aninhadas
// Structs aninhadas são structs que são declaradas dentro de outra struct
// Isso é útil quando você deseja agrupar dados relacionados em uma única unidade
// Nesse caso, o struct item é aninhado dentro do struct pedido

type item struct {
	nome       string
	preco      float64
	quantidade int
}

type pedido struct {
	clienteID string
	nome      string
	items     []item
}

func (p *pedido) total() float64 {
	total := 0.0
	for _, item := range p.items {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		clienteID: "12345678900",
		nome:      "Alef Silva",
		items: []item{
			{
				nome:       "Notebook",
				preco:      2500,
				quantidade: 1,
			},
			{
				nome:       "Mouse",
				preco:      50,
				quantidade: 2,
			},
		},
	}
	fmt.Printf("Cliente: %s\n", pedido.clienteID)
	fmt.Printf("Nome: %s\n", pedido.nome)
	fmt.Printf("Itens:\n")
	for _, item := range pedido.items {
		fmt.Printf("- %s: R$ %.2f x %d = R$ %.2f\n", item.nome, item.preco, item.quantidade, item.preco*float64(item.quantidade))
	}
	fmt.Printf("Valor Total do Pedido: R$ %.2f\n", pedido.total())
}
