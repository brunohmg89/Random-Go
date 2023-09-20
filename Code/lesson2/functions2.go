package main

import (
	"fmt"
	"strconv"
)

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = a + i

	return
}

func main() {
	total, err := convertAndSum(10, "23")
	fmt.Println("Total = ", total, err)
}
