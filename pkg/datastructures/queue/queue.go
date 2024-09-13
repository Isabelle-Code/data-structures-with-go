package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements [10]int
	head     int
	tail     int
	size     int
}

func NewQueue() *Queue {
	return &Queue{
		head: -1,
		tail: -1,
		size: 0,
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Enqueue(element int) error {

	if q.size == 10 {
		return errors.New("Queue overflow!")
	}

	q.size++
	if q.tail < 0 {
		q.head = 0
		q.tail = 0
		q.elements[q.tail] = element
	} else {
		q.tail++

		if q.tail >= 10 {
			q.tail = 0
		}

		q.elements[q.tail] = element
	}

	return nil
}

func (q *Queue) Dequeue() (int, error) {

	if q.size == 0 {
		return 0, errors.New("Queue underflow!")
	}

	q.size--
	element := q.elements[q.head]

	q.head++
	if q.head >= 10 {
		q.head = 0
	}

	return element, nil
}

func (q *Queue) Peek() (int, error) {

	if q.size == 0 {
		return 0, errors.New("Queue underflow!")
	}

	return q.elements[q.head], nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Print() {

	flag := true
	loop := false
	for i := q.head; flag; {
		message := fmt.Sprintf("[%d]", q.elements[i])
		fmt.Print(message)

		i++

		if q.head <= q.tail {
			flag = i <= q.tail
		} else if q.head > q.tail {
			flag = i >= q.tail && !loop
			if i >= 10 && flag {
				i = 0
				loop = true
			}
		}
	}
}
