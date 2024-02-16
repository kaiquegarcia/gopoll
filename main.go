package main

import (
	"fmt"
	"gopoll/ask"
	"time"
)

func main() {
	name, err := ask.String("What's your name?")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = ask.YesNo("Are you good?")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Whatever...")
	time.Sleep(2 * time.Second)

	pass, err := ask.Password("%s, what's the password?", name)
	if err != nil {
		fmt.Println(err)
		return
	}

	optIndex, err := ask.ChooseOption([]string{"ok", "i'll tell them right now!"}, "Shhh! Don't tell them the password is '%s'...", pass)
	if err != nil {
		fmt.Println(err)
		return
	}

	if optIndex == 0 {
		fmt.Println("Good... bye!")
		return
	} else {
		fmt.Println("Ok then... Go ahead...")
		time.Sleep(2 * time.Second)
	}

	indexes, err := ask.ChooseMultipleOptions(
		[]string{"hit by car", "hit by train", "hit by crazy bird"},
		"But first, what of the following options would you accept to die?",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(indexes) == 3 {
		fmt.Println("all of them? are you crazy?")
	} else if len(indexes) == 0 {
		fmt.Println("you learned well. bye")
	} else {
		fmt.Println("why did you choose something?? do you wanna die??")
	}
}
