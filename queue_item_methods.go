package main

import "fmt"

type QueItem struct {
	priority Priority
	content  string
	next     *QueItem
}

func NewQueueItem(priority Priority, content string, nextItem *QueItem) QueItem {
	return QueItem{
		priority: priority,
		content:  content,
		next:     nextItem,
	}
}

func (q *QueItem) String() string {
	if q.content == "" || q.priority == 0 {
		return ""
	}

	return fmt.Sprintf("{%d: %s}", q.priority, q.content)
}
