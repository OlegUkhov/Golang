SIS Project checkpoint 1

Text analyzer
A go application. A task takes a text as input, analyzes it and stores the result: number of words, runes and lines 

To run:
```
go run .
go test ./...
```

Files
- errors.go - application errors
- task.go - task, Status, state transtitions
- processor.go - processor interface, textanalyzer, input validation
- store.go - in memory store, map of tasks
- main.go - demo
- task_test.go - unit tests

Task lifecycle
pending -> processing -> completed or failed

Allowed transitions:
- pending -> processing
- processing -> completed
- processing -> failed

Validation
Empty or space only text is rejected with ErrEmptyInput
