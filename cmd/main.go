package main

import (
	"fmt"
)

func main() {
	var answer int
	// var l int = 1
	_, input := fmt.Scanf("%d", &answer)

	if input == nil {
		answer = answer * 200
	} else {
		fmt.Printf("somethine wrong")
	}

	fmt.Printf("USD to LKR %d", answer)
}
