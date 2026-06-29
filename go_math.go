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
	var range_from int
	flag.IntVar(&range_from, "from", 1, "range left border")
	flag.IntVar(&range_from, "f", 1, "range left border (shorthand)")

	var range_to int
	flag.IntVar(&range_to, "to", 9, "range right border")
	flag.IntVar(&range_to, "t", 9, "range right border (shorthand)")

	var limit int
	flag.IntVar(&limit, "limit", 20, "number of questions")
	flag.IntVar(&limit, "l", 20, "number of questions (shorthand)")

	flag.Parse()
	fmt.Println("Learn range from", range_from, "to", range_to, "with limit", limit)

	pairs := make_pairs(range_from, range_to)
	rand.Shuffle(len(pairs), func(i, j int) {
		pairs[i], pairs[j] = pairs[j], pairs[i]
	})
	limit = min(len(pairs), limit)
	pairs = pairs[:limit]

	for {
		incorrect_answers := learn_pairs(pairs)
		ilen := len(incorrect_answers)

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
			pairs = incorrect_answers
		} else {
			break
		}
	}
	fmt.Printf("%sDone%s\n", Green, Reset)
}

func make_pairs(range_from, range_to int) []Pair {
	diff := range_to - range_from
	pairs := make([]Pair, 0, diff*diff)
	for x := range_from; x <= range_to; x++ {
		for y := range_from; y <= range_to; y++ {
			pairs = append(pairs, Pair{x, y})
		}
	}
	return pairs
}

func learn_pairs(pairs []Pair) []Pair {
	var incorrect []Pair = []Pair{}

	for _, pair := range pairs {
		if !learn_pair(pair) {
			incorrect = append(incorrect, pair)
		}
	}
	return incorrect
}

func learn_pair(pair Pair) bool {
	correct_answer := pair.x * pair.y
	fmt.Printf("%d x %d = ", pair.x, pair.y)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	int_answer, err := strconv.Atoi(answer)

	if err == nil && int_answer == correct_answer {
		fmt.Printf("%sCorrect%s\n", Green, Reset)
		return true
	}
	fmt.Printf("%sCorrect answer is %d%s\n", Red, correct_answer, Reset)
	return false
}
