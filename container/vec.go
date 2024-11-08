package container

import (
	"reflect"
)

// Vec 结构体
type Vec[T any] struct {
	data []T
}

// NewVec 创建一个新的 Vec
func NewVec[T any](data []T) *Vec[T] {
	return &Vec[T]{data: data}
}

// Map 对每个元素应用函数
func (v *Vec[T]) Map(fn func(T) T) *Vec[T] {
	result := make([]T, len(v.data))
	for i, val := range v.data {
		result[i] = fn(val)
	}
	return &Vec[T]{data: result}
}

// Filter 过滤元素
func (v *Vec[T]) Filter(fn func(T) bool) *Vec[T] {
	result := make([]T, 0)
	for _, val := range v.data {
		if fn(val) {
			result = append(result, val)
		}
	}
	return &Vec[T]{data: result}
}

// ForEach 对每个元素应用函数
func (v *Vec[T]) ForEach(fn func(T)) *Vec[T] {
	for _, val := range v.data {
		fn(val)
	}
	return v
}

// Contains 检查元素是否存在
func (v *Vec[T]) Contains(val T) bool {
	for _, item := range v.data {
		if reflect.DeepEqual(item, val) {
			return true
		}
	}
	return false
}

// Len 返回长度
func (v *Vec[T]) Len() int {
	return len(v.data)
}

// Container 返回内部数据
func (v *Vec[T]) Container() []T {
	return v.data
}

// Reduce 对元素进行归约
func (v *Vec[T]) Reduce(fn func(T, T) T, initial T) T {
	result := initial
	for _, val := range v.data {
		result = fn(result, val)
	}
	return result
}

// FlatMap 对每个元素应用函数并展平结果
func (v *Vec[T]) FlatMap(fn func(T) []T) *Vec[T] {
	result := make([]T, 0)
	for _, val := range v.data {
		result = append(result, fn(val)...)
	}
	return &Vec[T]{data: result}
}

// GroupBy 根据函数分组
func (v *Vec[T]) GroupBy(fn func(T) string) map[string][]T {
	result := make(map[string][]T)
	for _, val := range v.data {
		key := fn(val)
		result[key] = append(result[key], val)
	}
	return result
}

// Distinct 去重
func (v *Vec[T]) Distinct() *Vec[T] {
	result := make([]T, 0)
	seen := make(map[interface{}]bool)
	for _, val := range v.data {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return &Vec[T]{data: result}
}

// Take 获取前 n 个元素
func (v *Vec[T]) Take(n int) *Vec[T] {
	if n > len(v.data) {
		n = len(v.data)
	}
	return &Vec[T]{data: v.data[:n]}
}

// Skip 跳过前 n 个元素
func (v *Vec[T]) Skip(n int) *Vec[T] {
	if n > len(v.data) {
		n = len(v.data)
	}
	return &Vec[T]{data: v.data[n:]}
}

// Reverse 反转 Vec
func (v *Vec[T]) Reverse() *Vec[T] {
	result := make([]T, len(v.data))
	for i, j := 0, len(v.data)-1; i < len(v.data); i, j = i+1, j-1 {
		result[i] = v.data[j]
	}
	return &Vec[T]{data: result}
}
