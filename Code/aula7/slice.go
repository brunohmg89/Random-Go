package main

import "fmt"

func main() {
	nomes := []string{"Bruno", "Henrique", "Maia", "Gonçalves"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Tamara")
	fmt.Println(nomes, len(nomes), cap(nomes))

	//iniciando um slice com tamanho e capacidade especifico
	nomes2 := make([]string, 10, 20)
	fmt.Println(nomes2, len(nomes2), cap(nomes2))
}
