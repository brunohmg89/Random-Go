package main

import "fmt"

func main() {
	//exemplo de for
	nome := []string{"Bruno", "Daniel", "Maria", "João"}
	for k, nome := range nome {
		fmt.Println(k, nome)
	}

	for _, nome := range nome {
		fmt.Println(nome)
	}
}
