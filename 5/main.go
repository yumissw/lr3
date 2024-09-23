package main

import (
	"fmt"
)

func main() {

	m := []int{1, 2, 3}
	fmt.Println("массив: ", m)

	m = append(m, 4)
	fmt.Println("после добавления элемента: ", m)

	var n = 2
	m = append(m[:n], m[n+1:]...)

	fmt.Println("после удаления 3 элемента: ", m)
}
