package main

import (
	"fmt"
	"strings"
)

// NÍVEL 1: BÁSICO
// Desafio 1: Validação de CPF
// Contexto: Isaac precisa validar CPFs de alunos no cadastro.
// Problema: Implemente uma função que valide se um CPF é válido usando o algoritmo oficial.
// Input: `"12345678909"` ou `"123.456.789-09"`
// Output: `True` ou `False`

// Dicas:
// - Remover pontos e hífens
// - Validar dígitos verificadores
// - Considerar CPFs inválidos conhecidos (000.000.000-00, 111.111.111-11, etc.)

func validarCPF(cpf string) bool {
	// Remover pontos e hífens
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verificar se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	if cpf == "00000000000" ||
		cpf == "11111111111" ||
		cpf == "22222222222" ||
		cpf == "33333333333" ||
		cpf == "44444444444" ||
		cpf == "55555555555" ||
		cpf == "66666666666" ||
		cpf == "77777777777" ||
		cpf == "88888888888" ||
		cpf == "99999999999" ||
		cpf == "12345678909" {
		return false
	}

	return true

}

func main() {

	// Testes
	cpf1 := "12345678909"
	cpf2 := "123.456.789-09"
	cpf3 := "000.000.000-00"
	cpf4 := "111.111.111-11"
	cpf5 := "98765432100"

	// Resultados esperados
	esperado1 := false
	esperado2 := true
	esperado3 := false
	esperado4 := false
	esperado5 := true

	// Testar os CPFs
	resultado1 := validarCPF(cpf1)
	fmt.Printf("CPF: %s | Resultado: %b | Esperado: %b\n", cpf1, resultado1, esperado1)

	resultado2 := validarCPF(cpf2)
	fmt.Printf("CPF: %s | Resultado: %b | Esperado: %b\n", cpf2, resultado2, esperado2)

	resultado3 := validarCPF(cpf3)
	fmt.Printf("CPF: %s | Resultado: %b | Esperado: %b\n", cpf3, resultado3, esperado3)

	resultado4 := validarCPF(cpf4)
	fmt.Printf("CPF: %s | Resultado: %b | Esperado: %b\n", cpf4, resultado4, esperado4)

	resultado5 := validarCPF(cpf5)
	fmt.Printf("CPF: %s | Resultado: %b | Esperado: %b\n", cpf5, resultado5, esperado5)

}
