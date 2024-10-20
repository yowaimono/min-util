package minutil

import (
	"container/list"
)

// LinkedMap 是一个带有双索引的泛型map，使用list维护元素的顺序
type LinkedMap[K comparable, V any] struct {
	dict map[K]*list.Element
	list *list.List
}

// entry 是list中存储的元素，包含key和value
type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewLinkedMap 创建一个新的LinkedMap
func NewLinkedMap[K comparable, V any]() *LinkedMap[K, V] {
	return &LinkedMap[K, V]{
		dict: make(map[K]*list.Element),
		list: list.New(),
	}
}

// Set 设置键值对，如果键已存在则更新值，并返回当前LinkedMap
func (lm *LinkedMap[K, V]) Set(key K, value V) *LinkedMap[K, V] {
	if elem, exists := lm.dict[key]; exists {
		// 更新值
		elem.Value.(*entry[K, V]).value = value
	} else {
		// 插入新值
		elem := lm.list.PushBack(&entry[K, V]{key: key, value: value})
		lm.dict[key] = elem
	}
	return lm
}

// Get 获取键对应的值，如果不存在返回默认值
func (lm *LinkedMap[K, V]) Get(key K, defaultValue ...V) V {
	if elem, exists := lm.dict[key]; exists {
		return elem.Value.(*entry[K, V]).value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	var zero V
	return zero
}

// Delete 删除键值对，并返回当前LinkedMap
func (lm *LinkedMap[K, V]) Delete(key K) *LinkedMap[K, V] {
	if elem, exists := lm.dict[key]; exists {
		lm.list.Remove(elem)
		delete(lm.dict, key)
	}
	return lm
}

// Keys 返回所有键的列表
func (lm *LinkedMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(lm.dict))
	for elem := lm.list.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(*entry[K, V]).key)
	}
	return keys
}

// Values 返回所有值的列表
func (lm *LinkedMap[K, V]) Values() []V {
	values := make([]V, 0, lm.list.Len())
	for elem := lm.list.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(*entry[K, V]).value)
	}
	return values
}

// Filter 过滤元素，返回满足条件的元素列表
func (lm *LinkedMap[K, V]) Filter(predicate func(V) bool) []V {
	filtered := make([]V, 0)
	for elem := lm.list.Front(); elem != nil; elem = elem.Next() {
		value := elem.Value.(*entry[K, V]).value
		if predicate(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// Foreach 对每个元素执行操作
func (lm *LinkedMap[K, V]) Foreach(action func(K, V)) {
	for elem := lm.list.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(*entry[K, V])
		action(e.key, e.value)
	}
}

// ReverseForeach 从后向前对每个元素执行操作
func (lm *LinkedMap[K, V]) ReverseForeach(action func(K, V)) {
	for elem := lm.list.Back(); elem != nil; elem = elem.Prev() {
		e := elem.Value.(*entry[K, V])
		action(e.key, e.value)
	}
}

// Map 对每个元素执行操作，并返回一个新的LinkedMap
func (lm *LinkedMap[K, V]) Map(transform func(K, V) V) *LinkedMap[K, V] {
	newLm := NewLinkedMap[K, V]()
	for elem := lm.list.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(*entry[K, V])
		newValue := transform(e.key, e.value)
		newLm.Set(e.key, newValue)
	}
	return newLm
}

// Reduce 对所有元素执行归约操作，返回最终结果
func (lm *LinkedMap[K, V]) Reduce(initial V, reducer func(V, K, V) V) V {
	result := initial
	for elem := lm.list.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(*entry[K, V])
		result = reducer(result, e.key, e.value)
	}
	return result
}

// BatchSet 批量设置键值对
func (lm *LinkedMap[K, V]) BatchSet(items map[K]V) *LinkedMap[K, V] {
	for key, value := range items {
		lm.Set(key, value)
	}
	return lm
}

// BatchDelete 批量删除键值对
func (lm *LinkedMap[K, V]) BatchDelete(keys []K) *LinkedMap[K, V] {
	for _, key := range keys {
		lm.Delete(key)
	}
	return lm
}
