package main

import "fmt"

func main() {
	//exemplo de for com array
	nome := []string{"Bruno", "Daniel", "Maria", "João"}
	for j := 0; j < len(nome); j++ {
		fmt.Println(nome[j])
	}
}
