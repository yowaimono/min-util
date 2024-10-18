package minutil

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestJWTTokenManager(t *testing.T) {
	secretKey := "my_secret_key"

	// 使用 MapTokenStore 进行测试
	mapStore := NewMapTokenStore()
	tokenManager := GetTokenManager(secretKey, mapStore)

	Warn("Start Map Store Test!")
	testTokenManager(t, tokenManager)

	// 使用 RedisTokenStore 进行测试
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redisStore := NewRedisTokenStore(redisClient)
	tokenManager = GetTokenManager(secretKey, redisStore)

	Warn("Start Redis Store Test!")
	testTokenManager(t, tokenManager)
}

func testTokenManager(t *testing.T, tokenManager *JWTTokenManager) {
	userID := "123456"
	username := "testuser"
	expiration := 1 * time.Hour

	// 生成 Token
	token, err := tokenManager.GenerateToken(userID, username, expiration)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// 验证 Token
	claims, err := tokenManager.ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected user ID %s, got %s", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}

	// 使 Token 失效
	err = tokenManager.InvalidateToken(token)
	if err != nil {
		t.Fatalf("Failed to invalidate token: %v", err)
	}

	// 验证失效的 Token
	_, err = tokenManager.ValidateToken(token)
	if err == nil {
		t.Errorf("Expected token to be invalidated, but it was valid")
	} else if err.Error() != "token is invalidated" {
		t.Errorf("Expected error 'token is invalidated', got %v", err)
	}

	// 验证过期的 Token
	expiredToken, err := tokenManager.GenerateToken(userID, username, -1*time.Hour)
	if err != nil {
		t.Fatalf("Failed to generate expired token: %v", err)
	}

	_, err = tokenManager.ValidateToken(expiredToken)
	if err == nil {
		t.Errorf("Expected token to be expired, but it was valid")
	}
}
