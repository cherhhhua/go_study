package main

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	//初始化list
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	v := stack.list.Back()
	if v != nil {
		stack.list.Remove(v)
		return v.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	v := stack.list.Back()
	if v != nil {
		return v.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
