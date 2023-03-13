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

	fmt.Scanf("%s", &option)

	optionsArray = strings.Split(option, ",")

	return optionsArray
}

func RemoveOption(options []string, option string) []string {
	index, ok := Contains(options, option)

	if ok {
		return append(options[:index], options[index+1:]...)
	}

	return nil
}

func BeatAndBeatenOptions(removed []string) ([]string, []string) {
	var beat []string
	var beaten []string
	for i := 0; i < len(removed)/2; i++ {
		beat = append(beat, removed[i])
	}
	for i := len(removed) / 2; i < len(removed); i++ {
		beaten = append(beaten, removed[i])
	}
	return beat, beaten
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

	computerChose := ReadOptions()
	fmt.Println("Okay, let's start")
	for {

		random := rand.Intn(len(computerChose))

		userInput := Read()

		user.Chose = userInput
		user.Name = username

		computer.Chose = computerChose[random]

		removed := RemoveOption(computerChose, user.Chose)

		beat, beaten := BeatAndBeatenOptions(removed)

		_, okBeat := Contains(beat, computer.Chose)
		_, okBeaten := Contains(beaten, computer.Chose)

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
