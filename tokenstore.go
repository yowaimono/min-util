package minutil

// TokenStore 是 Token 存储的接口
type TokenStore interface {
	Get(token string) (bool, error)
	Set(token string, invalid bool) error
}