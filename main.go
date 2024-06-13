package main

import "fmt"

type Priority int

const (
	Highest Priority = iota + 1
	High
	Medium
	Low
	Lowest
)

func main() {
	// Create a new PriorityQueue
	q := NewPriorityQueue()

	// Create some QueueItems
	i1 := NewQueueItem(Highest, "learn for the test", nil)
	i2 := NewQueueItem(High, "finish Homework", nil)
	i3a := NewQueueItem(Medium, "do sports", nil)
	i3b := NewQueueItem(Medium, "eat properly", nil)
	i4 := NewQueueItem(Low, "tidy Up my room", nil)
	i5 := NewQueueItem(Lowest, "buy new game", nil)

	// Add QueueItems to the PriorityQueue
	q.Insert(&i5)
	q.Insert(&i2)
	q.Insert(&i1)
	q.Insert(&i3b)
	q.Insert(&i3a)
	q.Insert(&i4)

	fmt.Printf("The PriorityQue looks like this:\n%s\n", q.String())
	fmt.Printf("The PriorityQueue has a length of %d.\n", q.Length())

	doNow := q.PopHighestPriority()
	fmt.Printf("Task to do now: %s\n", doNow.String())
	fmt.Printf("The PriorityQue looks like this:\n%s\n", q.String())
	fmt.Printf("The PriorityQueue has a length of %d.\n", q.Length())
}
