package commands

import "fmt"

func isPresent[T comparable](confirm T, inv []T) int {
	var idx int = -1
	for i, e := range inv {
		if e == confirm {
			idx = i
		}
	}
	return idx
}

func removeElem[T any](idx int, s []T) []T {
	s[idx] = s[len(s)-1]
	fmt.Println(s[:len(s)-1])
	return s[:len(s)-1]
}
