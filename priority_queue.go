package main

import "fmt"

type PriorityQueue struct {
	head *QueItem
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{}
}

func (p *PriorityQueue) Insert(new *QueItem) {
	// don't append 'nil' element
	if new == nil || new == (&QueItem{}) {
		return
	}

	// if queue is empty set new head
	if p.IsEmpty() {
		p.head = new
		return
	}

	// set new head if 'new.value' is smaller
	if p.head.priority >= new.priority {
		temp := p.head
		p.head = new
		p.head.next = temp
		return
	}

	current := p.head

	// walk queue until fitting position is found
	for {

		if current.next == nil {
			break
		}

		if current.next.priority < new.priority {
			current = current.next
			continue
		}

		if current.next.priority >= new.priority {
			break
		}
	}

	temp := current.next
	current.next = new
	current.next.next = temp
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.head == nil
}

func (p *PriorityQueue) Peek() *QueItem {
	return p.head
}

func (p *PriorityQueue) PopHighestPriority() *QueItem {
	popItem := p.head
	nextItem := popItem.next

	if nextItem == nil {
		p.head = nil
	}

	p.head = nextItem
	popItem.next = nil

	return popItem
}

func (p *PriorityQueue) Length() int {
	i := 0
	current := p.head
	for {
		if current == nil {
			break
		}

		current = current.next
		i++
	}

	return i
}

func (p *PriorityQueue) String() string {
	if p.IsEmpty() {
		return "-> none"
	}

	var s string

	item := p.head
	for {
		if item == nil {
			break
		}
		s = fmt.Sprintf("%s -> %s", s, item.String())
		item = item.next
	}

	return s
}
