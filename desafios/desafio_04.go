package main

import "fmt"

type Aluno struct {
	Matricula string
	Nome      string
	Serie     string
}

func buscarAluno(matricula string, alunos []Aluno) Aluno {
	for _, aluno := range alunos {
		if aluno.Matricula == matricula {
			return aluno
		}
	}
	return Aluno{}
}

func main() {

	alunos := []Aluno{
		{
			Matricula: "2024001",
			Nome:      "João Silva",
			Serie:     "5A",
		},
		{
			Matricula: "2024002",
			Nome:      "Maria Santos",
			Serie:     "6B",
		},
	}

	aluno := buscarAluno("2024002", alunos)
	fmt.Println(aluno)

}
