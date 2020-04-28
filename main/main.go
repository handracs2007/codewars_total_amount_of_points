// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Points(games []string) int {
	var score = 0

	for _, game := range games {
		// Separate the score by :
		var scores = strings.Split(game, ":")
		var score1, _ = strconv.Atoi(scores[0])
		var score2, _ = strconv.Atoi(scores[1])

		if score1 > score2 {
			score += 3
		} else if score1 == score2 {
			score += 1
		}
	}

	return score
}

func main() {
	fmt.Println(Points([]string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"}))
}
