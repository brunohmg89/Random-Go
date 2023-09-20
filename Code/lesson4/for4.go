package main

import "fmt"

func main() {
	//exemplo de for
	nome := []string{"Bruno", "Daniel", "Maria", "João"}
	var k int
	for k < len(nome) {
		fmt.Println(nome[k])
		k++
	}
}
