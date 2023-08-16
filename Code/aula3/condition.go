package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++

	case "coroa":
		coroa++

	default:
		fmt.Println("caiu em pé")
	}
}

func main() {
	a, b := 15, 15

	if a > b {
		fmt.Println("a é a maior que b")
	} else if a < b {
		fmt.Println("a é a menor que b")
	} else {
		fmt.Println("a e b são iguais")
	}

	switch {
	case a > b:
		fmt.Println("a é a maior que b")

	case a < b:
		fmt.Println("a é a menor que b")

	default:
		fmt.Println("a e b são iguais")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}
