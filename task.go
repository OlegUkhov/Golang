package main

type Status int

const StatusPending Status = 1
const StatusProcessing Status = 2
const StatusCompleted Status = 3
const StatusFailed Status = 4

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusProcessing:
		return "processing"
	case StatusCompleted:
		return "completed"
	case StatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

type Result struct {
	Words int
	Runes int
	Lines int
}

type Task struct {
	ID     int
	Input  string
	Status Status
	Result *Result // check later
}

var allowedTransitions = map[Status][]Status{
	StatusPending:    {StatusProcessing},
	StatusProcessing: {StatusCompleted, StatusFailed},
}

func (t *Task) transition(to Status) error {
	for _, a := range allowedTransitions[t.Status] {
		if a == to {
			t.Status = to
			return nil
		}
	}
	return ErrInvalidTransition
}
