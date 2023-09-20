package main

import "fmt"

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pj.cnpj)
}

func show(d Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	pf := PessoaFisica{
		Pessoa{Nome: "Bruno", Idade: 34, Status: true},
		"Maia",
		"000.000.000-00",
	}

	show(pf)
	fmt.Println()

	pj := PessoaJuridica{
		Pessoa{Nome: "Brunos", Idade: 1, Status: true},
		"Bruno LTDA",
		"00.000.000/0000-00",
	}

	show(pj)

}
