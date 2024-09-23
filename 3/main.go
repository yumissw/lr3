package main

import (
	"bufio"
	"fmt"
	stringutils "mnm/package1"
	"os"
)

func main() {
	fmt.Print("введите строку: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("перевернутая строка: ", stringutils.Reverse(s))
}
