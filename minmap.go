package minutil

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"time"
)

// PersistenceStrategy 持久化策略
type PersistenceStrategy int

const (
	// ImmediateFlush 每次写操作后立即刷新
	ImmediateFlush PersistenceStrategy = iota
	// PeriodicFlush 定期刷新
	PeriodicFlush
	// ManualFlush 手动刷新
	ManualFlush
)

// MinMap 结构体
type MinMap struct {
	mu          sync.RWMutex
	data        map[string]ValueWithExpiry
	walFile     *os.File
	walWriter   *bufio.Writer
	strategy    PersistenceStrategy
	flushTicker *time.Ticker
}

// ValueWithExpiry 包含值和过期时间
type ValueWithExpiry struct {
	Value     interface{}
	ExpiresAt time.Time
}

// NewMinMap 创建一个新的 MinMap
func NewMinMap(walFilePath string, strategy PersistenceStrategy) (*MinMap, error) {
	walFile, err := os.OpenFile(walFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	mm := &MinMap{
		data:      make(map[string]ValueWithExpiry),
		walFile:   walFile,
		walWriter: bufio.NewWriter(walFile),
		strategy:  strategy,
	}

	// 从 WAL 文件中恢复数据
	if err := mm.recoverFromWAL(); err != nil {
		return nil, err
	}

	// 根据持久化策略启动定时刷新
	if strategy == PeriodicFlush {
		mm.flushTicker = time.NewTicker(time.Second * 5)
		go mm.periodicFlush()
	}

	return mm, nil
}

// recoverFromWAL 从 WAL 文件中恢复数据
func (mm *MinMap) recoverFromWAL() error {
	scanner := bufio.NewScanner(mm.walFile)
	for scanner.Scan() {
		var entry struct {
			Op    string
			Key   string
			Value interface{}
			Exp   time.Time
		}
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			return err
		}

		switch entry.Op {
		case "set":
			mm.data[entry.Key] = ValueWithExpiry{
				Value:     entry.Value,
				ExpiresAt: entry.Exp,
			}
		case "delete":
			delete(mm.data, entry.Key)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Set 设置键值对，支持过期时间
func (mm *MinMap) Set(key string, value interface{}, expires time.Duration) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	var expiresAt time.Time
	if expires > 0 {
		expiresAt = time.Now().Add(expires)
	}

	entry := struct {
		Op    string
		Key   string
		Value interface{}
		Exp   time.Time
	}{
		Op:    "set",
		Key:   key,
		Value: value,
		Exp:   expiresAt,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	if _, err := mm.walWriter.Write(data); err != nil {
		return err
	}
	if err := mm.walWriter.WriteByte('\n'); err != nil {
		return err
	}

	mm.data[key] = ValueWithExpiry{
		Value:     value,
		ExpiresAt: expiresAt,
	}

	// 根据持久化策略决定是否立即刷新
	if mm.strategy == ImmediateFlush {
		return mm.walWriter.Flush()
	}

	return nil
}

// Get 获取键值对，支持惰性清除
func (mm *MinMap) Get(key string) (interface{}, error) {
	mm.mu.RLock()
	valueWithExpiry, exists := mm.data[key]
	mm.mu.RUnlock()

	if !exists {
		return nil, errors.New("key not found")
	}

	if !valueWithExpiry.ExpiresAt.IsZero() && time.Now().After(valueWithExpiry.ExpiresAt) {
		mm.mu.Lock()
		delete(mm.data, key)
		mm.mu.Unlock()
		return nil, errors.New("key expired")
	}

	return valueWithExpiry.Value, nil
}

// Delete 删除键值对
func (mm *MinMap) Delete(key string) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	entry := struct {
		Op  string
		Key string
	}{
		Op:  "delete",
		Key: key,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	if _, err := mm.walWriter.Write(data); err != nil {
		return err
	}
	if err := mm.walWriter.WriteByte('\n'); err != nil {
		return err
	}

	delete(mm.data, key)

	// 根据持久化策略决定是否立即刷新
	if mm.strategy == ImmediateFlush {
		return mm.walWriter.Flush()
	}

	return nil
}

// Flush 手动刷新 WAL 缓冲区
func (mm *MinMap) Flush() error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	return mm.walWriter.Flush()
}

// periodicFlush 定期刷新 WAL 缓冲区
func (mm *MinMap) periodicFlush() {
	for range mm.flushTicker.C {
		if err := mm.Flush(); err != nil {
			log.Printf("Failed to flush WAL buffer: %v", err)
		}
	}
}

// Close 关闭 MinMap
func (mm *MinMap) Close() error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	if mm.flushTicker != nil {
		mm.flushTicker.Stop()
	}

	if err := mm.walWriter.Flush(); err != nil {
		return err
	}

	return mm.walFile.Close()
}
