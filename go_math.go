package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
)

type Pair struct {
	x int
	y int
}

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

	incorrect_answers := learn_pairs(pairs)
	fmt.Println("Incorrect answers:", incorrect_answers)
}

func make_pairs(range_from, range_to int) []Pair {
	pairs := []Pair{}
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

	var answer string
	fmt.Scanln(&answer)
	int_answer, err := strconv.Atoi(answer)

	if err == nil && int_answer == correct_answer {
		fmt.Println("Correct")
		return true
	}
	fmt.Println("Correct answer is", correct_answer)
	return false
}
