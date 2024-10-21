package minutil

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Storage 接口定义了存储后端需要实现的方法
type Storage interface {
	Get(key string) (time.Time, bool)
	Set(key string, value time.Time)
}

// RateLimiter 是一个简单的限流器，使用自定义的存储后端
type RateLimiter struct {
	storage Storage
}

// NewRateLimiter 创建一个新的 RateLimiter 实例
func NewRateLimiter(storage Storage) *RateLimiter {
	return &RateLimiter{
		storage: storage,
	}
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow(key string, duration time.Duration) bool {
	now := time.Now()
	if val, ok := rl.storage.Get(key); ok {
		lastRequestTime := val
		if now.Sub(lastRequestTime) < duration {
			return false
		}
	}
	rl.storage.Set(key, now)
	return true
}

// generateKey 生成一个唯一的键，基于用户 IP 和 User-Agent
func generateKey(c *gin.Context) string {
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	hash := sha256.Sum256([]byte(ip + userAgent))
	return hex.EncodeToString(hash[:])
}

// RateLimitMiddleware 是一个 Gin 中间件，用于限制重复请求
func RateLimitMiddleware(limiter *RateLimiter, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := generateKey(c)
		if !limiter.Allow(key, duration) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests, please try again later.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// 示例：使用 sync.Map 作为存储后端
type SyncMapStorage struct {
	cache sync.Map
}

func (s *SyncMapStorage) Get(key string) (time.Time, bool) {
	if val, ok := s.cache.Load(key); ok {
		return val.(time.Time), true
	}
	return time.Time{}, false
}

func (s *SyncMapStorage) Set(key string, value time.Time) {
	s.cache.Store(key, value)
}

// 示例：使用 Redis 作为存储后端
// 你可以根据需要实现 RedisStorage