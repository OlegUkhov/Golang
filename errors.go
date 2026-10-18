package main

import "errors"

var ErrEmptyInput = errors.New("invalid input")
var ErrNotFound = errors.New("task is not found")
var ErrInvalidTransition = errors.New("invalid transition")
