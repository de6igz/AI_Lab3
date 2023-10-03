package main

import (
	"AII_lab3/internal/provider"
	"fmt"
	"github.com/ichiban/prolog"
)

func main() {
	//var userInput string
	// Создание новой базы знаний Prolog
	base, _ := provider.MakeNewBase()
	for {
		do(base)
	}
	//_, err := fmt.Scan(&userInput)
	//if err != nil {
	//	return
	//}
	//fmt.Println(userInput)
	//provider.ShowResult(
	//	base,
	//	userInput,
	//)

}

func do(base *prolog.Interpreter) {
	var userInput string

	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}
	fmt.Println(userInput)
	provider.ShowResult(
		base,
		userInput,
	)
	fmt.Printf("1: Покажи всех пероснажей\n" +
		"2: Покажи братьев/сестер\n" +
		"3: Покажи персонажей и их джутсу\n" +
		"4: Покажи деревни и персонажей, которые живут в них\n" +
		"5: Покажи персонажей, которые состоят в отношениях\n" +
		"6: Покажи учителей и их учеников\n")
}
