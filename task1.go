/*
Project: Go Algorithms and Concurrency
Description: Находит наибольший общий префикс в заданном массиве строк.
Author: Никита Мякишев
License: GPLv3
History:
  - [03.04.2025 17:02]: Initial task1
*/

package main

import (
	"fmt"
)

// Функция для поиска наибольшего общего префикса
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(prefix) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// Вспомогательная функция для проверки префикса
func startsWith(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func main() {
	words := []string{"flower", "flow", "flight"}
	fmt.Println("Наибольший общий префикс:", longestCommonPrefix(words))
}
