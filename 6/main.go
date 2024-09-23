package main

import (
	"fmt"
)

func main() {

	s := []string{"hello", "bye", "morning"}

	l := ""

	fmt.Print("все строки: ")

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}

	for _, a := range s {
		if len(a) > len(l) {
			l = a
		}
	}

	fmt.Println("\nсамая длинная строка: ", l)

}
