package main

import (
	"flag"
	"fmt"
)

func main() {
	var range_from int
	flag.IntVar(&range_from, "from", 1, "range left border")
	flag.IntVar(&range_from, "f", 1, "range left border (shorthand)")

	var range_to int
	flag.IntVar(&range_to, "to", 9, "range right border")
	flag.IntVar(&range_to, "t", 9, "range right border (shorthand)")

	var limit int
	flag.IntVar(&limit, "limit", 10, "number of questions")
	flag.IntVar(&limit, "l", 10, "number of questions (shorthand)")

	flag.Parse()
	fmt.Println("from", range_from, "to", range_to, "limit", limit)
}
