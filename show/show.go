package show

import (
	"fmt"
	"strconv"

	"github.com/serapmtr/rock-paper-scissors/options"
)

func ShowScore(lines []string, username string) int {
	index, ok := options.Contains(lines, username)
	if ok {
		fmt.Println("Your rating:", lines[index+1])
		score, _ := strconv.Atoi(lines[index+1])
		return score
	} else {
		fmt.Println("Your rating: 0")
		return 0
	}
}
