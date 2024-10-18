package minutil

import (
	"errors"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenClaims 是 JWT 的声明部分
type TokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// TokenManager 是 Token 管理的接口
type TokenManager interface {
	GenerateToken(userID, username string, expiration time.Duration) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
	InvalidateToken(tokenString string) error
}

// JWTTokenManager 是基于 JWT 的 Token 管理实现
type JWTTokenManager struct {
	secretKey string
	store     TokenStore
}

var (
	instance *JWTTokenManager
	once     sync.Once
)

// GetTokenManager 返回单例模式的 TokenManager
func GetTokenManager(secretKey string, store TokenStore) *JWTTokenManager {
	once.Do(func() {
		instance = &JWTTokenManager{
			secretKey: secretKey,
			store:     store,
		}
	})
	return instance
}

// GenerateToken 生成一个新的 JWT Token
func (tm *JWTTokenManager) GenerateToken(userID, username string, expiration time.Duration) (string, error) {
	claims := TokenClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.secretKey))
}

// ValidateToken 验证 JWT Token
func (tm *JWTTokenManager) ValidateToken(tokenString string) (*TokenClaims, error) {
	invalid, err := tm.store.Get(tokenString)
	if err != nil {
		return nil, err
	}

	if invalid {
		return nil, errors.New("token is invalidated")
	}

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// InvalidateToken 使 Token 失效
func (tm *JWTTokenManager) InvalidateToken(tokenString string) error {
	return tm.store.Set(tokenString, true)
}