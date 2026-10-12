package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Blue  = "\033[34m"
)

type Pair struct {
	x int
	y int
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var rangeFrom int
	flag.IntVar(&rangeFrom, "from", 1, "range left border")
	flag.IntVar(&rangeFrom, "f", 1, "range left border (shorthand)")

	var rangeTo int
	flag.IntVar(&rangeTo, "to", 9, "range right border")
	flag.IntVar(&rangeTo, "t", 9, "range right border (shorthand)")

	var limit int
	flag.IntVar(&limit, "limit", 20, "number of questions")
	flag.IntVar(&limit, "l", 20, "number of questions (shorthand)")

	flag.Parse()
	fmt.Println("Learn range from", rangeFrom, "to", rangeTo, "with limit", limit)

	pairs := makePairs(rangeFrom, rangeTo)
	rand.Shuffle(len(pairs), func(i, j int) {
		pairs[i], pairs[j] = pairs[j], pairs[i]
	})
	limit = min(len(pairs), limit)
	pairs = pairs[:limit]

	for {
		incorrectAnswers := learnPairs(pairs)
		ilen := len(incorrectAnswers)

		if ilen == 0 {
			break
		}

		var msg string
		if ilen == 1 {
			msg = "%sThere is %d incorrect answer%s, repeat? y/n: "
		} else {
			msg = "%sThere are %d incorrect answers%s, repeat? y/n: "
		}
		fmt.Printf(msg, Red, ilen, Reset)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, "y") {
			pairs = incorrectAnswers
		} else {
			break
		}
	}
	fmt.Printf("%sDone%s\n", Green, Reset)
}

func makePairs(rangeFrom, rangeTo int) []Pair {
	diff := rangeTo - rangeFrom
	pairs := make([]Pair, 0, diff*diff)
	for x := rangeFrom; x <= rangeTo; x++ {
		for y := rangeFrom; y <= rangeTo; y++ {
			pairs = append(pairs, Pair{x, y})
		}
	}
	return pairs
}

func learnPairs(pairs []Pair) []Pair {
	var incorrect []Pair = []Pair{}

	for _, pair := range pairs {
		if !learnPair(pair) {
			incorrect = append(incorrect, pair)
		}
	}
	return incorrect
}

func learnPair(pair Pair) bool {
	correctAnswer := pair.x * pair.y
	fmt.Printf("%d x %d = ", pair.x, pair.y)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	intAnswer, err := strconv.Atoi(answer)

	if err == nil && intAnswer == correctAnswer {
		fmt.Printf("%sCorrect%s\n", Green, Reset)
		return true
	}
	fmt.Printf("%sCorrect answer is %d%s\n", Red, correctAnswer, Reset)
	return false
}
