package container

// Set 结构体
type Set[T comparable] struct {
	data map[T]bool
}

// NewSet 创建一个新的 Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]bool)}
}

// NewSetFromSlice 从切片创建一个新的 Set
func NewSetFromSlice[T comparable](items []T) *Set[T] {
	set := NewSet[T]()
	for _, item := range items {
		set.Add(item)
	}
	return set
}

// Add 添加元素
func (s *Set[T]) Add(item T) *Set[T] {
	s.data[item] = true
	return s
}

// Remove 移除元素
func (s *Set[T]) Remove(item T) *Set[T] {
	delete(s.data, item)
	return s
}

// Contains 检查元素是否存在
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.data[item]
	return exists
}

// Len 返回长度
func (s *Set[T]) Len() int {
	return len(s.data)
}

// ToSlice 转换为切片
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.data))
	for item := range s.data {
		slice = append(slice, item)
	}
	return slice
}

// Union 并集
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.data {
		result.Add(item)
	}
	for item := range other.data {
		result.Add(item)
	}
	return result
}

// Intersection 交集
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.data {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference 差集
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.data {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// ForEach 对每个元素应用函数
func (s *Set[T]) ForEach(fn func(T)) *Set[T] {
	for item := range s.data {
		fn(item)
	}
	return s
}

// Map 对每个元素应用函数并返回新的 Set
func (s *Set[T]) Map(fn func(T) T) *Set[T] {
	result := NewSet[T]()
	for item := range s.data {
		result.Add(fn(item))
	}
	return result
}

// Filter 过滤元素并返回新的 Set
func (s *Set[T]) Filter(fn func(T) bool) *Set[T] {
	result := NewSet[T]()
	for item := range s.data {
		if fn(item) {
			result.Add(item)
		}
	}
	return result
}

// Reduce 对元素进行归约
func (s *Set[T]) Reduce(fn func(T, T) T, initial T) T {
	result := initial
	for item := range s.data {
		result = fn(result, item)
	}
	return result
}
