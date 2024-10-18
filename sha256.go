package minutil

import (
	"crypto/sha256"
	"encoding/hex"
)

// Encrypt 使用 SHA-256 对输入字符串进行哈希，并返回十六进制编码的字符串
func Encrypt(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// Verify 验证输入字符串的 SHA-256 哈希是否与给定的哈希值匹配
func Verify(input, hashed string) bool {
	return Encrypt(input) == hashed
}