package main

import (
	"cr1/task1"
	"fmt"
)

func main() {
	// Задача 1: Найти наибольший общий префикс в заданном массиве строк
	words := []string{"flower", "flow", "flight"}
	fmt.Println("Наибольший общий префикс:", task1.LongestCommonPrefix(words))
}
