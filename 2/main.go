package main

import (
	"fmt"
	mathutils "mnm/package"
)

func main() {
	var x int

	fmt.Print("введите число: ")
	fmt.Scan(&x)

	fmt.Println(mathutils.Factorial(x))

}
