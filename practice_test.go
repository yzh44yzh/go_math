package main

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestSettingsValidation(t *testing.T) {
	cases := []struct {
		name     string
		settings settings
		valid    bool
	}{
		{"defaults", defaultSettings, true}, {"one factor", settings{5, 5, 100}, true},
		{"zero factor", settings{0, 9, 20}, false}, {"high factor", settings{1, 10, 20}, false},
		{"reversed", settings{9, 1, 20}, false}, {"negative", settings{-1, 9, 20}, false},
		{"zero count", settings{1, 9, 0}, false}, {"high count", settings{1, 9, 101}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.settings.validate()
			if (err == nil) != c.valid {
				t.Fatalf("valid=%v, error=%v", c.valid, err)
			}
		})
	}
}

func TestQuestionSelection(t *testing.T) {
	for _, s := range []settings{{1, 9, 100}, {3, 4, 100}, {9, 9, 20}, {1, 9, 20}} {
		p, err := newPractice(s, rand.New(rand.NewSource(1)))
		if err != nil {
			t.Fatal(err)
		}
		if len(p.pairs) != min(s.count, s.available()) {
			t.Fatalf("wrong count for %+v: %d", s, len(p.pairs))
		}
		seen := map[question]bool{}
		for _, q := range p.pairs {
			if q.x < s.from || q.x > s.to || q.y < s.from || q.y > s.to || seen[q] {
				t.Fatalf("out-of-range or duplicate pair: %+v", q)
			}
			seen[q] = true
		}
		if s.from == 3 && s.to == 4 && (!seen[question{3, 4}] || !seen[question{4, 3}]) {
			t.Fatal("swapped factors must be separate questions")
		}
		other, _ := newPractice(s, rand.New(rand.NewSource(1)))
		if !reflect.DeepEqual(p.pairs, other.pairs) {
			t.Fatal("seeded selection is not reproducible")
		}
	}
	if _, err := newPractice(settings{9, 1, 20}, rand.New(rand.NewSource(1))); err == nil {
		t.Fatal("invalid settings accepted")
	}
}

func TestAnswerTransitions(t *testing.T) {
	p, _ := newPractice(settings{2, 2, 1}, rand.New(rand.NewSource(1)))
	if p.next() || p.retry() {
		t.Fatal("unanswered question advanced")
	}
	for _, input := range []string{"", " ", "x", "4.0", "99999999999999999999999999999"} {
		if p.submit(input) == nil {
			t.Fatalf("accepted %q", input)
		}
		if p.state != questionPhase || p.correct != 0 || len(p.mistakes) != 0 {
			t.Fatal("invalid input changed practice")
		}
	}
	if err := p.submit(" 4 \n"); err != nil {
		t.Fatal(err)
	}
	if p.state != feedbackPhase || p.index != 0 || p.correct != 1 {
		t.Fatal("feedback must wait for explicit advancement")
	}
	if p.submit("0") == nil || p.correct != 1 || len(p.mistakes) != 0 {
		t.Fatal("duplicate submission changed result")
	}
	if p.retry() {
		t.Fatal("retry allowed during feedback")
	}
	if !p.next() || p.state != resultsPhase || p.next() || p.retry() {
		t.Fatal("invalid completed-round transition")
	}
}

func TestRepeatedMistakeOnlyRetries(t *testing.T) {
	p, _ := newPractice(settings{1, 3, 9}, rand.New(rand.NewSource(4)))
	wanted := map[question]bool{}
	for p.state != resultsPhase {
		q := p.current()
		answer := q.product()
		if q.x != q.y {
			answer = -1
			wanted[q] = true
		}
		if err := p.submit(strconv.Itoa(answer)); err != nil {
			t.Fatal(err)
		}
		p.next()
	}
	if p.correct != 3 || len(p.mistakes) != 6 {
		t.Fatalf("wrong totals: correct=%d mistakes=%d", p.correct, len(p.mistakes))
	}
	if !p.retry() || p.round != 1 || p.correct != 0 || len(p.mistakes) != 0 {
		t.Fatal("retry did not reset round")
	}
	seen := map[question]bool{}
	for _, q := range p.pairs {
		if !wanted[q] || seen[q] {
			t.Fatalf("invalid retry pair: %+v", q)
		}
		seen[q] = true
	}
	if len(seen) != len(wanted) {
		t.Fatal("retry omitted mistakes")
	}
	remaining := p.current()
	for p.state != resultsPhase {
		answer := p.current().product()
		if p.index == 0 {
			answer = -1
		}
		p.submit(strconv.Itoa(answer))
		p.next()
	}
	if p.correct != 5 || len(p.mistakes) != 1 || !p.retry() || p.round != 2 {
		t.Fatal("second retry used wrong round")
	}
	if len(p.pairs) != 1 || p.current() != remaining {
		t.Fatal("second retry did not contain only the remaining mistake")
	}
	p.submit(strconv.Itoa(p.current().product()))
	p.next()
	if p.retry() {
		t.Fatal("all-correct round offered retry")
	}
}
