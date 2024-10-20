package minutil

import (
	"math/rand"
	"time"
)

// SkipListNode 是跳表中的节点
type SkipListNode[K comparable, V any] struct {
	key   K
	value V
	next  []*SkipListNode[K, V]
}

// SkipList 是跳表结构
type SkipList[K comparable, V any] struct {
	head     *SkipListNode[K, V]
	maxLevel int
	level    int
	compare  func(K, K) int
	randSrc  rand.Source
	rand     *rand.Rand
}

// NewSkipList 创建一个新的跳表
func NewSkipList[K comparable, V any](maxLevel int, compare func(K, K) int) *SkipList[K, V] {
	head := &SkipListNode[K, V]{
		next: make([]*SkipListNode[K, V], maxLevel),
	}
	return &SkipList[K, V]{
		head:     head,
		maxLevel: maxLevel,
		compare:  compare,
		randSrc:  rand.NewSource(time.Now().UnixNano()),
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// randomLevel 生成一个随机的层级
func (sl *SkipList[K, V]) randomLevel() int {
	level := 1
	for sl.rand.Float64() < 0.5 && level < sl.maxLevel {
		level++
	}
	return level
}

// Set 设置键值对，如果键已存在则更新值
func (sl *SkipList[K, V]) Set(key K, value V) {
	update := make([]*SkipListNode[K, V], sl.maxLevel)
	x := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && sl.compare(x.next[i].key, key) < 0 {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]
	if x != nil && sl.compare(x.key, key) == 0 {
		x.value = value
		return
	}

	level := sl.randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	x = &SkipListNode[K, V]{
		key:   key,
		value: value,
		next:  make([]*SkipListNode[K, V], level),
	}
	for i := 0; i < level; i++ {
		x.next[i] = update[i].next[i]
		update[i].next[i] = x
	}
}

// Get 获取键对应的值，如果不存在返回nil
func (sl *SkipList[K, V]) Get(key K) (V, bool) {
	x := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && sl.compare(x.next[i].key, key) < 0 {
			x = x.next[i]
		}
	}
	x = x.next[0]
	if x != nil && sl.compare(x.key, key) == 0 {
		return x.value, true
	}
	var zero V
	return zero, false
}

// Delete 删除键值对
func (sl *SkipList[K, V]) Delete(key K) {
	update := make([]*SkipListNode[K, V], sl.maxLevel)
	x := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && sl.compare(x.next[i].key, key) < 0 {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]
	if x == nil || sl.compare(x.key, key) != 0 {
		return
	}

	for i := 0; i < sl.level; i++ {
		if update[i].next[i] != x {
			break
		}
		update[i].next[i] = x.next[i]
	}

	for sl.level > 1 && sl.head.next[sl.level-1] == nil {
		sl.level--
	}
}

// Keys 返回所有键的列表
func (sl *SkipList[K, V]) Keys() []K {
	keys := make([]K, 0)
	x := sl.head.next[0]
	for x != nil {
		keys = append(keys, x.key)
		x = x.next[0]
	}
	return keys
}

// Values 返回所有值的列表
func (sl *SkipList[K, V]) Values() []V {
	values := make([]V, 0)
	x := sl.head.next[0]
	for x != nil {
		values = append(values, x.value)
		x = x.next[0]
	}
	return values
}

// Foreach 对每个元素执行操作
func (sl *SkipList[K, V]) Foreach(action func(K, V)) {
	x := sl.head.next[0]
	for x != nil {
		action(x.key, x.value)
		x = x.next[0]
	}
}
