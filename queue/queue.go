package queue

import (
	"errors"
)

type node struct {
	next  *node
	value interface{}
}

type Queue struct {
	size  int
	front *node
	rear  *node
}

func NewQueue() *Queue {

	return &Queue{0, nil, nil}

}

//추후 수정하고 싶음
//O(1) 로 할 방법이 없을까 ?
//Linked List 자료형을 만들까 ?
func (self *Queue) Push(item interface{}) {

	if self.size == 0 {
		self.rear = &node{nil, item}
		self.front = self.rear

	} else {
		self.rear.next = &node{nil, item}
		self.rear = self.rear.next

	}
	self.size++
}

//rear element
func (self *Queue) Back() (interface{}, error) {

	if self.size == 0 {
		return nil, errors.New("Empty Queue.")
	}

	return self.rear.value, nil
}

//front element
func (self *Queue) Front() (interface{}, error) {

	if self.size == 0 {
		return nil, errors.New("Empty Queue.")
	}

	return self.front.value, nil

}
func (self *Queue) Pop() error {

	if self.size == 0 {
		return errors.New("Empty Queue.")
	}

	self.front = self.front.next

	self.size--
	return nil
}
