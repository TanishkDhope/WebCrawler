package main

import (
	"sync"
)

type Stack[T any] struct{
	data []T
	len int
	mu sync.Mutex
}


func (s *Stack[T]) push(val T){
	s.data=append(s.data,val)
	s.len++
}

func (s *Stack[T]) pop() T{
	if s.len==0{
		var zero T
		return zero
	}
	last:=s.data[s.len-1]
	s.data=s.data[:s.len-1]
	s.len--
	return last
}
