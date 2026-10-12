package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

type settings struct{ from, to, count int }

var defaultSettings = settings{from: 1, to: 9, count: 20}

func (s settings) validate() error {
	if s.from < 1 || s.from > 9 || s.to < 1 || s.to > 9 {
		return errors.New("Factors must be whole numbers from 1 to 9")
	}
	if s.from > s.to {
		return errors.New("Minimum factor must not exceed maximum factor")
	}
	if s.count < 1 || s.count > 100 {
		return errors.New("Question count must be a whole number from 1 to 100")
	}
	return nil
}

func parseSettings(from, to, count string) (settings, error) {
	var parsed [3]int
	for i, value := range []string{from, to, count} {
		n, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			return settings{}, errors.New("Settings must be whole numbers")
		}
		parsed[i] = n
	}
	s := settings{parsed[0], parsed[1], parsed[2]}
	return s, s.validate()
}

func (s settings) available() int { return (s.to - s.from + 1) * (s.to - s.from + 1) }

type question struct{ x, y int }

func (q question) product() int { return q.x * q.y }

type phase int

const (
	questionPhase phase = iota
	feedbackPhase
	resultsPhase
)

type practice struct {
	pairs    []question
	index    int
	state    phase
	answer   int
	correct  int
	mistakes []question
	round    int
	rng      *rand.Rand
}

func newPractice(s settings, rng *rand.Rand) (*practice, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	pairs := make([]question, 0, s.available())
	for x := s.from; x <= s.to; x++ {
		for y := s.from; y <= s.to; y++ {
			pairs = append(pairs, question{x, y})
		}
	}
	rng.Shuffle(len(pairs), func(i, j int) { pairs[i], pairs[j] = pairs[j], pairs[i] })
	return &practice{pairs: pairs[:min(s.count, len(pairs))], rng: rng}, nil
}

func (p *practice) current() question { return p.pairs[p.index] }

func (p *practice) submit(text string) error {
	if p.state != questionPhase {
		return errors.New("This question has already been checked")
	}
	answer, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return errors.New("Enter a whole number")
	}
	p.answer = answer
	if answer == p.current().product() {
		p.correct++
	} else {
		p.mistakes = append(p.mistakes, p.current())
	}
	p.state = feedbackPhase
	return nil
}

func (p *practice) next() bool {
	if p.state != feedbackPhase {
		return false
	}
	p.index++
	p.answer = 0
	if p.index == len(p.pairs) {
		p.state = resultsPhase
	} else {
		p.state = questionPhase
	}
	return true
}

func (p *practice) retry() bool {
	if p.state != resultsPhase || len(p.mistakes) == 0 {
		return false
	}
	p.pairs = append([]question(nil), p.mistakes...)
	p.rng.Shuffle(len(p.pairs), func(i, j int) { p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i] })
	p.index, p.correct, p.answer = 0, 0, 0
	p.mistakes = nil
	p.round++
	p.state = questionPhase
	return true
}
