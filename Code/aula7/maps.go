package main

import "fmt"

func main() {
	idades := make(map[string]uint8)
	idades["Bruno"] = 34
	idades["Tamara"] = 33
	idades["Murilo"] = 6

	//validando chave
	val, ok := idades["Bruno"]
	fmt.Println(val, ok)

	fmt.Println(idades)
}
