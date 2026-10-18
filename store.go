package main

type Store struct {
	tasks     map[int]*Task
	nextID    int
	processor Processor
}

func NewStore(p Processor) *Store {
	return &Store{
		tasks:     make(map[int]*Task),
		nextID:    0,
		processor: p,
	}
}

func (s *Store) Create(input string) (*Task, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}
	s.nextID = s.nextID + 1
	t := &Task{ID: s.nextID, Input: input, Status: StatusPending}
	s.tasks[t.ID] = t
	return t, nil
}

func (s *Store) Get(id int) (*Task, error) {
	t, ok := s.tasks[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (s *Store) Process(id int) error {
	t, err := s.Get(id)
	if err != nil {
		return err
	}
	err = t.transition(StatusProcessing)
	if err != nil {
		return err
	}
	res, err := s.processor.Process(t.Input)
	if err != nil {
		if terr := t.transition(StatusFailed); terr != nil {
			return terr
		}
		return err
	}
	t.Result = &res
	return t.transition(StatusCompleted)
}
