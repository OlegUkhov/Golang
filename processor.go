package main

import (
	"strings"
	"unicode/utf8"
)

type Processor interface {
	Process(input string) (Result, error)
}

func validate(input string) error {
	if strings.TrimSpace(input) == "" {
		return ErrEmptyInput
	}
	return nil
}

type TextAnalyzer struct{}

func (TextAnalyzer) Process(input string) (Result, error) {
	err := validate(input)
	if err != nil {
		return Result{}, err
	}
	words := len(strings.Fields(input))
	runes := utf8.RuneCountInString(input)
	lines := strings.Count(input, "\n") + 1
	return Result{Words: words, Runes: runes, Lines: lines}, nil
}
