package util

import (
	"math"
	"strings"
)

func EstimateReadingTime(str string) int {
	wordLength := len(strings.Fields(str))
	return int(math.Round(float64(wordLength) / 200.0))
}
