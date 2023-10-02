package main

import (
	"AII_lab3/internal/provider"
	"fmt"
)

func main() {
	var userInput string
	// Создание новой базы знаний Prolog
	base, _ := provider.MakeNewBase()
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}
	fmt.Println(userInput)
	provider.ShowResult(
		base,
		userInput,
	)

}
