package main

import (
	"errors"
	"testing"
)

func TestCreate_OK(t *testing.T) {
	store := NewStore(TextAnalyzer{})
	task, err := store.Create("hello world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.ID == 0 {
		t.Errorf("ID must not be 0")
	}
	if task.Status != StatusPending {
		t.Errorf("status = %v, want pending", task.Status)
	}
	if task.Result != nil {
		t.Errorf("result must be nil before processing")
	}
}

func TestCreate_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"spaces", "   "},
		{"tabs and newlines", "\n\t"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore(TextAnalyzer{})
			_, err := store.Create(tc.input)
			if !errors.Is(err, ErrEmptyInput) {
				t.Errorf("err = %v, want ErrEmptyInput", err)
			}
		})
	}
}

func TestGet_OK(t *testing.T) {
	store := NewStore(TextAnalyzer{})
	created, _ := store.Create("hello")
	got, err := store.Get(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != created.ID || got.Input != "hello" {
		t.Errorf("got wrong task: %+v", got)
	}
}

func TestGet_NotFound(t *testing.T) {
	store := NewStore(TextAnalyzer{})
	_, err := store.Get(42)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestProcess_OK(t *testing.T) {
	store := NewStore(TextAnalyzer{})
	task, _ := store.Create("Hello Go world")
	err := store.Process(task.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Status != StatusCompleted {
		t.Errorf("status = %v, want completed", task.Status)
	}
	if task.Result == nil {
		t.Fatalf("result is nil")
	}
	if task.Result.Words != 3 || task.Result.Runes != 14 || task.Result.Lines != 1 {
		t.Errorf("wrong result: %+v", *task.Result)
	}
}

func TestTextAnalyzer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Result
	}{
		{"latin", "Hello Go world", Result{Words: 3, Runes: 14, Lines: 1}},
		{"cyrillic", "Привет мир", Result{Words: 2, Runes: 10, Lines: 1}},
		{"multiline", "one\ntwo\nthree", Result{Words: 3, Runes: 13, Lines: 3}},
		{"extra spaces", "  a   b  ", Result{Words: 2, Runes: 9, Lines: 1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := TextAnalyzer{}.Process(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestProcess_Twice(t *testing.T) {
	store := NewStore(TextAnalyzer{})
	task, _ := store.Create("hello")
	store.Process(task.ID)
	err := store.Process(task.ID)
	if !errors.Is(err, ErrInvalidTransition) {
		t.Errorf("err = %v, want ErrInvalidTransition", err)
	}
	if task.Status != StatusCompleted {
		t.Errorf("status changed to %v, want completed", task.Status)
	}
}

func TestMultipleTasks(t *testing.T) {
	store := NewStore(TextAnalyzer{})
	t1, _ := store.Create("hello world")
	t2, _ := store.Create("a b c")
	if t1.ID == t2.ID {
		t.Fatalf("IDs must differ")
	}
	store.Process(t1.ID)
	store.Process(t2.ID)
	if t1.Result.Words != 2 || t1.Result.Runes != 11 {
		t.Errorf("task 1 wrong result: %+v", *t1.Result)
	}
	if t2.Result.Words != 3 || t2.Result.Runes != 5 {
		t.Errorf("task 2 wrong result: %+v", *t2.Result)
	}
}
