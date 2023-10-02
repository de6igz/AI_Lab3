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
	//// Отправка запросов к базе знаний
	//query := "character(X)."
	//result, err := kb.Query(query)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Вывод результатов запроса
	//fmt.Println("Друзья Наруто:")
	//for result.Next() {
	//	var s struct {
	//		Friend string `prolog:"X"`
	//	}
	//	//fmt.Printf("%v", result)
	//	if err := result.Scan(&s); err != nil {
	//		panic(err)
	//	}
	////	fmt.Printf("Who = %s\n", s.Friend) // ==> Who = socrates
	//}
}
