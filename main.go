package main

import (
	"fmt"
	"math/rand"

	"github.com/serapmtr/rock-paper-scissors/options"
	"github.com/serapmtr/rock-paper-scissors/read"
)

type User struct {
	Name  string
	Score int
	Chose string
}

type Computer struct {
	Chose string
}

func main() {
	var user User
	var computer Computer

	username := read.ReadName()

	computerChose := options.ReadOptions()
	fmt.Println("Okay, let's start")
	for {

		random := rand.Intn(len(computerChose))

		userInput := read.Read()

		user.Chose = userInput
		user.Name = username

		computer.Chose = computerChose[random]

		removed := options.RemoveOption(computerChose, user.Chose)

		beat, beaten := options.BeatAndBeatenOptions(removed)

		_, okBeat := options.Contains(beat, computer.Chose)
		_, okBeaten := options.Contains(beaten, computer.Chose)

		switch userInput {
		case "!rating":
			fmt.Println("Your score:", user.Score)
		case "!exit":
			fmt.Println("Bye!")
		default:
			if okBeat {
				fmt.Println("Sorry, but the computer chose", computer.Chose)
			} else if okBeaten {
				fmt.Println("Well done. The computer chose", computer.Chose, "and failed")
				user.Score += 100
			} else if computer.Chose == userInput {
				fmt.Println("There is a draw (", computer.Chose, ")")
				user.Score += 50
			} else {
				fmt.Println("Invalid input")
			}
		}
	}

}
