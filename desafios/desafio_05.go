// "ID1:05-03-2025,ID2:01-02-2025,ID3:05-03-2025,ID4:01-02-2025,ID5:04-07-2025,ID6:04-07-2025,ID7:20-05-2025,ID8:01-02-2025,ID9:05-03-2025,ID10:01-02-2025";
// resposta: ["01-02-2025", "05-03-2025", "04-07-2025"]
/*
// Dado uma lista de datas de vencimentos de mensalidade, retornar em order as datas que mais aparecem para que menos aparecem
// Retornando uma data de cada uma dessas mesalidades, sendo as três que mais aparecem.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	datasVencimentosFrequencia := make(map[string]int)

	input := "ID1:05-03-2025,ID2:01-02-2025,ID3:05-03-2025,ID4:01-02-2025,ID5:04-07-2025,ID6:04-07-2025,ID7:20-05-2025,ID8:01-02-2025,ID9:05-03-2025,ID10:01-02-2025"

	dataParse := strings.Split(input, ",")

	for _, data := range dataParse {
		dataValor := strings.Split(data, ":")
		data := dataValor[1]
		datasVencimentosFrequencia[data]++
	}

	type dataCount struct {
		data  string
		count int
	}

	var sortedDatas []dataCount

	for data, count := range datasVencimentosFrequencia {
		sortedDatas = append(sortedDatas, dataCount{data: data, count: count})
	}

	sort.Slice(sortedDatas, func(i, j int) bool {
		return sortedDatas[i].count > sortedDatas[j].count
	})

	var resulFinal []string
	for i := 0; i < 3 && i < len(sortedDatas); i++ {
		resulFinal = append(resulFinal, sortedDatas[i].data)
	}

	fmt.Println(resulFinal)

}
