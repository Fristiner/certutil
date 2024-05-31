package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器

	// 定义字符集，包含大小写字母和数字
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 使用随机数生成器从字符集中随机选择字符，组成字符串
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
