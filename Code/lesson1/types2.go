package main

import "fmt"

var (
	nome string
	n1   int
	n2   int32
)

func main() {
	//Quando uma váriavel é declarada ela precisa ser utilizada
	/*var (
		b1 float32
		b2 float64
	)*/

	//Várivel pode ser criada dessa forma
	// var mensagem string = "Aula 03"

	//Podemos juntar as váriaveis de mesmo tipo e também declarar na mesma linha utilizando
	//a mesma sequencia utilizando ,
	//var a, b, c int32
	var b, f, s = true, 2.3, "Olá"
	fmt.Println("Váriavel Boolean:", b, "- Váriavel Float:", f, "- Váriavel String:", s)

	//int tem seu valor padrão igual a 0
	var total int
	total++
	fmt.Println("Total:", total)

	//string tem seu valor padrão igual a "" vazio.
	var mensagem2 string
	fmt.Println("Total:", mensagem2)

	//Ou podemos omitir utilizando :=
	mensagem := "Aula 03"
	fmt.Println(mensagem)

	nome = "Bruno Henrique"
	fmt.Println("Hello", nome, "!")

	//troca de valores
	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
