package main

import "fmt"

// Métodos
// Métodos são funções que são associadas a um tipo
// Isso significa que você pode definir métodos para um tipo e chamá-los diretamente
// Nesse caso, o método getNomeCompleto é associado ao tipo pessoa
// Isso significa que você pode chamar p.getNomeCompleto() para obter o nome completo da pessoa
// O método setNomeCompleto é associado ao tipo *pessoa
// Isso significa que você pode chamar p.setNomeCompleto("João Santos") para definir o nome completo da pessoa
// O método setNomeCompleto é associado ao tipo *pessoa
// Isso significa que você pode chamar p.setNomeCompleto("João Santos") para definir o nome completo da pessoa

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	p.nome, p.sobrenome = nomeCompleto, ""
}

func main() {
	p := pessoa{
		nome:      "Alef",
		sobrenome: "Silva",
	}

	fmt.Println(p.getNomeCompleto())

	p.setNomeCompleto("João Santos")

	fmt.Println(p.getNomeCompleto())
}
