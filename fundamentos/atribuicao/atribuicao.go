package main

import "fmt"

func main() {

	var b byte = 3 // declaração de variável - declara a variável b do tipo byte e atribui o valor 3
	fmt.Println(b)

	i := 3 // inferência de tipo - declara e inicializa a variável i com valor 3
	i += 3 // atribuição de adição - adiciona 3 ao i (i = i + 3)
	i -= 3 // atribuição de subtração - subtrai 3 do i (i = i - 3)
	i *= 2 // atribuição de multiplicação - multiplica i por 2 (i = i * 2)
	i /= 2 // atribuição de divisão - divide i por 2 (i = i / 2)
	i %= 2 // atribuição de módulo - define i como o resto da divisão de i por 2 (i = i % 2)
	fmt.Println(i)

	x, y := 1, 2 // inferência de tipo - declara e inicializa as variáveis x e y com valores 1 e 2, respectivamente
	fmt.Println(x, y)

	x, y = y, x // atribuição de múltipla variável - atribui o valor de y à x e o valor de x à y
	fmt.Println(x, y)

}
