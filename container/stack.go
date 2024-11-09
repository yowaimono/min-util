package container

import (
	"errors"
)

// Stack 结构体
type Stack[T any] struct {
	data []T
}

// NewStack 创建一个新的 Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// Push 添加元素到栈顶
func (s *Stack[T]) Push(item T) *Stack[T] {
	s.data = append(s.data, item)
	return s
}

// Pop 移除并返回栈顶元素
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item, nil
}

// Peek 返回栈顶元素但不移除
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty 检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Len 返回栈的长度
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// Clear 清空栈
func (s *Stack[T]) Clear() *Stack[T] {
	s.data = make([]T, 0)
	return s
}

// ToSlice 转换为切片
func (s *Stack[T]) ToSlice() []T {
	return s.data
}

// ForEach 对每个元素应用函数
func (s *Stack[T]) ForEach(fn func(T)) *Stack[T] {
	for _, item := range s.data {
		fn(item)
	}
	return s
}

// Map 对每个元素应用函数并返回新的 Stack
func (s *Stack[T]) Map(fn func(T) T) *Stack[T] {
	result := NewStack[T]()
	for _, item := range s.data {
		result.Push(fn(item))
	}
	return result
}

// Filter 过滤元素并返回新的 Stack
func (s *Stack[T]) Filter(fn func(T) bool) *Stack[T] {
	result := NewStack[T]()
	for _, item := range s.data {
		if fn(item) {
			result.Push(item)
		}
	}
	return result
}

// Reduce 对元素进行归约
func (s *Stack[T]) Reduce(fn func(T, T) T, initial T) T {
	result := initial
	for _, item := range s.data {
		result = fn(result, item)
	}
	return result
}
