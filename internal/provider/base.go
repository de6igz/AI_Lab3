package provider

import (
	"fmt"
	"github.com/ichiban/prolog"
	"os"
	"strconv"
)

const base = `
		% Факты о персонажах
		character(naruto).
		character(sasuke).
		character(sakura).
		character(kakashi).
		character(itachi).
		character(orochimaru).
		character(hinata).
		character(shikamaru).
		character(rock_lee).
		character(gaara).
		character(tsunade).
		character(jiraiya).
		character(sarutobi).
		character(minato).
		character(kushina).
		character(neji).
		character(ino).
		character(shino).
		character(kiba).

		% Предикаты о техниках
		jutsu(naruto, rasengan).
		jutsu(sasuke, chidori).
		jutsu(kakashi, sharingan).
		jutsu(itachi, amaterasu).
		jutsu(orochimaru, snake_summoning).
		jutsu(hinata, byakugan).
		jutsu(gaara, sand_control).
		jutsu(tsunade, medical_jutsu).
		jutsu(jiraiya, toad_summoning).
		jutsu(minato, flying_thunder_god).

		% Предикаты о деревнях
		village(naruto, konoha).
		village(sasuke, konoha).
		village(sakura, konoha).
		village(kakashi, konoha).
		village(itachi, konoha).
		village(orochimaru, sound).
		village(hinata, konoha).
		village(shikamaru, konoha).
		village(rock_lee, konoha).
		village(gaara, sand).
		village(tsunade, konoha).
		village(jiraiya, konoha).
		village(sarutobi, konoha).
		village(minato, konoha).
		village(kushina, konoha).
		village(neji, konoha).
		village(ino, konoha).
		village(shino, konoha).
		village(kiba, konoha).

		% Отношение стран и деревень
		world(konoha, fire_country).

		% Отношения между персонажами
		friend(naruto, sasuke).
		friend(naruto, sakura).
		teacher(kakashi, naruto).
		teacher(jiraiya, naruto).
		teacher(sarutobi, kakashi).
		teacher(orochimaru, sasuke).
		mentor(sasuke, orochimaru).
		mentor(naruto, jiraiya).
		sibling(itachi, sasuke).
		sibling(hinata, neji).
		love(naruto, hinata).
		love(minato, kushina).

		% Факты о врагах
		enemy(sasuke, itachi).
		enemy(itachi, sasuke).

		% Правила
		is_student(X, Y) :- teacher(Y, X).
		is_mentor(X, Y) :- mentor(X, Y).
		is_sibling(X, Y) :- sibling(X, Y); sibling(Y, X).
		is_friends_with(X, Y) :- friend(X, Y); friend(Y, X).
		is_enemy(X, Y) :- enemy(X, Y); enemy(Y, X).
	`

func MakeNewBase() (*prolog.Interpreter, error) {
	// Создаем новый интерпретатор Prolog.
	p := prolog.New(os.Stdin, os.Stdout) // Or prolog.New(nil, nil) if you don't need user_input/user_output.

	// Загружаем программу Prolog.
	if err := p.Exec(base); err != nil {
		return nil, err
	}
	fmt.Printf("1: Покажи всех пероснажей\n" +
		"2: Покажи братьев/сестер\n" +
		"3: Покажи персонажей и их джутсу\n" +
		"4: Покажи деревни и персонажей, которые живут в них\n" +
		"5: Покажи персонажей, которые состоят в отношениях\n" +
		"6: Покажи учителей и их учеников\n")
	return p, nil
}

func ShowResult(base *prolog.Interpreter, variant string) {

	num, err := strconv.Atoi(variant)
	if err != nil {
		_ = fmt.Errorf("вы ввели не число")
		return
	}
	switch num {
	case 1:
		result, err2 := base.Query("character(X).")
		if err2 != nil {
			return
		}
		for result.Next() {
			var s struct {
				Friend string `prolog:"X"`
			}
			if err := result.Scan(&s); err != nil {
				panic(err)
			}
			fmt.Printf("Персонаж = %s\n", s.Friend)
		}
	case 2:
		result, err2 := base.Query("sibling(X,Y).")
		if err2 != nil {
			return
		}
		for result.Next() {
			var s struct {
				CharacterOne string `prolog:"X"`
				CharacterTwo string `prolog:"Y"`
			}
			if err := result.Scan(&s); err != nil {
				panic(err)
			}
			fmt.Printf("%s и %s родственники\n", s.CharacterOne, s.CharacterTwo)
		}
	case 3:
		result, err2 := base.Query("jutsu(X,Y).")
		if err2 != nil {
			return
		}
		for result.Next() {
			var s struct {
				Character string `prolog:"X"`
				Jutsu     string `prolog:"Y"`
			}
			if err := result.Scan(&s); err != nil {
				panic(err)
			}
			fmt.Printf("%s имеет джутсу %s\n", s.Character, s.Jutsu)
		}
	case 4:
		result, err2 := base.Query("village(X,Y).")
		if err2 != nil {
			return
		}
		for result.Next() {
			var s struct {
				Character string `prolog:"X"`
				Village   string `prolog:"Y"`
			}
			if err := result.Scan(&s); err != nil {
				panic(err)
			}
			fmt.Printf("%s из деревни %s\n", s.Character, s.Village)
		}
	case 5:
		result, err2 := base.Query("love(X,Y).")
		if err2 != nil {
			return
		}
		for result.Next() {
			var s struct {
				CharacterOne string `prolog:"X"`
				CharacterTwo string `prolog:"Y"`
			}
			if err := result.Scan(&s); err != nil {
				panic(err)
			}
			fmt.Printf("%s в отношениях с %s\n", s.CharacterOne, s.CharacterTwo)
		}
	case 6:
		result, err2 := base.Query("is_student(X,Y).")
		if err2 != nil {
			return
		}
		for result.Next() {
			var s struct {
				Student string `prolog:"X"`
				Teacher string `prolog:"Y"`
			}
			if err := result.Scan(&s); err != nil {
				panic(err)
			}
			fmt.Printf("%s является учеником %s\n", s.Student, s.Teacher)
		}
	default:
		fmt.Printf("Такого варианта нет")

	}
}
