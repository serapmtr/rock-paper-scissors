package read

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
