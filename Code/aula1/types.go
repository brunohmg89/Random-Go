package main

import "fmt"

var (
	nome string
	n1   int
	n2   int32
)

func main() {
	var b, f, s = true, 2.3, "Olá"
	fmt.Println("Váriavel Boolean:", b, "- Váriavel Float:", f, "- Váriavel String:", s)

	var total int
	total++
	fmt.Println("Total:", total)

	var mensagem2 string
	fmt.Println("Total:", mensagem2)

	mensagem := "Aula 03"
	fmt.Println(mensagem)

	nome = "Bruno Henrique"
	fmt.Println("Hello", nome, "!")

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
