package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func (stu Student) Talk(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	peo := Student{}
	peoref := &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	fmt.Println(peo.Talk(think))
	fmt.Println(peoref.Speak(think))
	fmt.Println(peoref.Talk(think))
}
