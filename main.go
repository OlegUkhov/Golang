package main

import "fmt"

func main() {
	store := NewStore(TextAnalyzer{})

	t1, _ := store.Create("Hello Go world")
	t2, _ := store.Create("Привет мир")

	store.Process(t1.ID)
	store.Process(t2.ID)

	fmt.Println(t1.ID, t1.Status, *t1.Result)
	fmt.Println(t2.ID, t2.Status, *t2.Result)

	_, err := store.Create("   ")
	fmt.Println("empty input:", err)

	err = store.Process(t1.ID)
	fmt.Println("process twice:", err)

	_, err = store.Get(999)
	fmt.Println("unknown id:", err)
}
