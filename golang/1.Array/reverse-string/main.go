package main

import "fmt"

// "strings"

func reverserString(words string) {
	// O(n)

	if words == "" || len(words) == 1 {
		return
	}

	runes := []rune(words)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println("========================================")
	fmt.Printf("%+v\n", string(runes))
	fmt.Println("========================================")
}

func main() {

	word := "I am Asad"

	reverserString(word)

}
