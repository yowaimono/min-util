package container

import (
	"errors"
	"fmt"
)

// Deque 是一个泛型双端队列
type Deque[T any] struct {
	items []T
}

// NewDeque 创建一个新的 Deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{items: make([]T, 0)}
}

// PushFront 在队列的前面插入一个元素
func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// PushBack 在队列的后面插入一个元素
func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

// PopFront 从队列的前面移除一个元素并返回它
func (d *Deque[T]) PopFront() (T, error) {
	if len(d.items) == 0 {
		var zero T
		return zero, errors.New("deque is empty")
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, nil
}

// PopBack 从队列的后面移除一个元素并返回它
func (d *Deque[T]) PopBack() (T, error) {
	if len(d.items) == 0 {
		var zero T
		return zero, errors.New("deque is empty")
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, nil
}

// Front 返回队列前面的元素但不移除它
func (d *Deque[T]) Front() (T, error) {
	if len(d.items) == 0 {
		var zero T
		return zero, errors.New("deque is empty")
	}
	return d.items[0], nil
}

// Back 返回队列后面的元素但不移除它
func (d *Deque[T]) Back() (T, error) {
	if len(d.items) == 0 {
		var zero T
		return zero, errors.New("deque is empty")
	}
	return d.items[len(d.items)-1], nil
}

// Size 返回队列的大小
func (d *Deque[T]) Size() int {
	return len(d.items)
}

// Empty 检查队列是否为空
func (d *Deque[T]) Empty() bool {
	return len(d.items) == 0
}

// Clear 清空队列
func (d *Deque[T]) Clear() {
	d.items = make([]T, 0)
}
