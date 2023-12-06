package util

import (
	"strings"
)

func EstimateReadingTime(str string) int {
	wordLength := len(strings.Fields(str))
	return wordLength / 200
}
