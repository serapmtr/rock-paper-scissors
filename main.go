package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name  string
	Score int
	Chose string
}

type Computer struct {
	Chose string
}

func Contains(array []string, search string) (int, bool) {
	for index, a := range array {
		if a == search {
			return index, true
		}
	}
	return 0, false
}

func ReadOptions() []string {
	var optionsArray []string
	var option string

	fmt.Scanf("%s", option)

	optionsArray = strings.Split(option, ",")

	return optionsArray
}

func RemoveOption(options []string, option string) []string {
	index, ok := Contains(options, option)

	if ok {
		return append(options[:index], options[index+1:]...)
	}

	return options
}

func Read() string {
	var userInput string

	fmt.Scanf("%s", &userInput)

	return userInput
}

func ReadName() string {
	var userName string

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s\n", &userName)
	fmt.Println("Hello, ", userName)
	return userName

}

func ReadTxt() []string {
	var lines []string
	file, err := os.Open("rating.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// i=0, username=Tom
func ShowScore(lines []string, username string) int {
	index, ok := Contains(lines, username)
	if ok {
		fmt.Println("Your rating:", lines[index+1])
		score, _ := strconv.Atoi(lines[index+1])
		return score
	} else {
		fmt.Println("Your rating: 0")
		return 0
	}
}

func main() {
	var user User
	var computer Computer

	username := ReadName()
	lines := ReadTxt()

	user.Score = ShowScore(lines, username)
	for {

		computerChose := ReadOptions()

		random := rand.Intn(len(computerChose))

		userInput := Read()

		user.Chose = userInput
		user.Name = username

		switch userInput {
		case "paper":
			switch random {
			case 0:
				computer.Chose = computerChose[0]
				fmt.Println("There is a draw (", computer.Chose, ")")
				user.Score += 50

			case 1:
				computer.Chose = computerChose[1]
				fmt.Println("Sorry, but the computer chose", computer.Chose)

			case 2:
				computer.Chose = computerChose[2]
				fmt.Println("Well done. The computer chose", computer.Chose, "and failed")
				user.Score += 100
			}
		case "scissors":
			switch random {
			case 0:
				computer.Chose = computerChose[0]
				fmt.Println("Well done. The computer chose", computer.Chose, "and failed")
				user.Score += 100
			case 1:
				computer.Chose = computerChose[1]
				fmt.Println("There is a draw (", computer.Chose, ")")
				user.Score += 50
			case 2:
				computer.Chose = computerChose[2]
				fmt.Println("Sorry, but the computer chose", computer.Chose)
			}
		case "rock":
			switch random {
			case 0:
				computer.Chose = computerChose[0]
				fmt.Println("Sorry, but the computer chose", computer.Chose)
			case 1:
				computer.Chose = computerChose[1]
				fmt.Println("Well done. The computer chose", computer.Chose, "and failed")
				user.Score += 100
			case 2:
				computer.Chose = computerChose[2]
				fmt.Println("There is a draw (", computer.Chose, ")")
				user.Score += 50
			}
		case "!rating":
			fmt.Println("Your score:", user.Score)
		case "!exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid input")
		}
	}

}
