package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("break")
		break
	}

	for k := 1; k < 5; k++ {
		fmt.Println(k)
	}
}
