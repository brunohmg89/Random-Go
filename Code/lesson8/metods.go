package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

// uma categoria pode chamar outra a mesma sendo um ponteiro "*"
type Categoria struct {
	Nome string
	Pai  *Categoria
}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}

func main() {
	p := Pessoa{"Bruno", "Maia", 34, true, "000.000.000-00"}
	fmt.Println(p)

	cat := Categoria{Nome: "Categoria 1"}
	cat.SetPai(&Categoria{Nome: "Pai"})
	if !cat.HasParent() {
		fmt.Println("não tem pai")
	}

}
