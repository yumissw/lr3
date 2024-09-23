package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var m [5]int

	for i := 0; i < len(m); i++ {
		m[i] = rand.Intn(50)
		fmt.Print(m[i], " ")
	}

}
