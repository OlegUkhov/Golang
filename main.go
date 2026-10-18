package main

import "fmt"

func main() {
	store := NewStore(TextAnalyzer{})

	t1, err := store.Create("Hello Go world")
	if err != nil {
		fmt.Println("create error:", err)
		return
	}
	t2, err := store.Create("Привет мир")
	if err != nil {
		fmt.Println("create error:", err)
		return
	}

	if err := store.Process(t1.ID); err != nil {
		fmt.Println("process error:", err)
		return
	}
	if err := store.Process(t2.ID); err != nil {
		fmt.Println("process error:", err)
		return
	}

	fmt.Println(t1.ID, t1.Status, *t1.Result)
	fmt.Println(t2.ID, t2.Status, *t2.Result)

	_, err = store.Create("   ")
	fmt.Println("empty input:", err)

	err = store.Process(t1.ID)
	fmt.Println("process twice:", err)

	_, err = store.Get(999)
	fmt.Println("unknown id:", err)
}
