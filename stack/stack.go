package stack

import (
	"errors"
)

type node struct {
	prev  *node
	value interface{}
}

type Stack struct {
	size    int
	current *node
}

func NewStack() *Stack {

	return &Stack{0, nil}
}

func (self *Stack) Push(item interface{}) {

	self.current = &node{self.current, item}
	self.size++
}

func (self *Stack) Top() (interface{}, error) {

	if self.size == 0 {
		return nil, errors.New("Empty Stack.")
	}

	return self.current.value, nil
}

func (self *Stack) Pop() (interface{}, error) {

	if self.size == 0 {
		return nil, errors.New("Empty Stack.")
	}

	ret := self.current.value
	self.current = self.current.prev

	self.size--
	return ret, nil
}
