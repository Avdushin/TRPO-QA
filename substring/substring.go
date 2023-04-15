package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Введите первую строку: ")
	fmt.Scanln(&str1)
	fmt.Print("Введите вторую строку: ")
	fmt.Scanln(&str2)

	count := strings.Count(str1, str2)

	fmt.Printf("Количество вхождений второй строки в первую: %d\n", count)
}
