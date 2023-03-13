package options

import (
	"fmt"
	"strings"
)

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
