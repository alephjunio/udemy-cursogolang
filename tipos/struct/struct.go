package main

import "fmt"

// Produto representa um produto no estoque
// O que é uma struct?
// Struct é uma coleção de campos
// Cada campo tem um nome e um tipo
// Os campos são separados por vírgula
// Os campos são declarados na ordem em que serão utilizados
type Produto struct {
	Nome       string
	Preco      float64
	Desconto   float64
	Quantidade int
}

// PrecoComDesconto retorna o preço do produto com o desconto aplicado
func (p Produto) PrecoComDesconto() float64 {
	return p.Preco * (1 - p.Desconto)
}

// AumentarQuantidade aumenta a quantidade do produto em estoque
func (p *Produto) AumentarQuantidade(qtd int) {
	p.Quantidade += qtd
}

// ReduzirQuantidade reduz a quantidade do produto em estoque
func (p *Produto) ReduzirQuantidade(qtd int) {
	p.Quantidade -= qtd
}

func main() {

	produto := Produto{
		Nome:       "Notebook",
		Preco:      2500,
		Desconto:   0.1,
		Quantidade: 10,
	}

	fmt.Println("--------------------------------")
	fmt.Println("Nome:", produto.Nome)
	fmt.Println("Preço:", produto.Preco)
	fmt.Println("Desconto:", produto.Desconto)
	fmt.Println("Quantidade:", produto.Quantidade)
	fmt.Println("--------------------------------")
	fmt.Println("Preço com desconto:", produto.PrecoComDesconto())
	fmt.Println("Quantidade no Estoque:", produto.Quantidade)
	fmt.Println("--------------------------------")
	produto.AumentarQuantidade(5)
	fmt.Println("Quantidade no Estoque após aumento:", produto.Quantidade)
	fmt.Println("--------------------------------")
	produto.ReduzirQuantidade(3)
	fmt.Println("Quantidade no Estoque após redução:", produto.Quantidade)
}
