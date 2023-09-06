package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

// uma categoria pode chamar outra a mesma sendo um ponteiro "*"
/*
type Categoria struct{
	Nome string
	Pai *Categoria
}*/

func main() {
	p := Pessoa{
		Nome:      "Bruno",
		Sobrenome: "Maia",
		Idade:     34,
		Status:    true,
		cpf:       "000.000.000-00",
	}

	//Pode ser feito assim também na sequencia
	//p := Pessoa{"Bruno", "Maia", 34, true}

	fmt.Println(p)

	fmt.Println(p.Nome, p.Idade, p.cpf)

}
