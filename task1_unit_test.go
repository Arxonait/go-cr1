package main

import (
	"testing"
)

// Unit-тест для проверки longestCommonPrefix
func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"a"}, "a"},
		{[]string{"abc", "abc", "abc"}, "abc"},
		{[]string{"", "b"}, ""},
	}

	for _, testCase := range testCases {
		result := longestCommonPrefix(testCase.input)
		if result != testCase.expected {
			t.Errorf("Для входных данных %v ожидалось %s, но получено %s", testCase.input, testCase.expected, result)
		}
	}
}
