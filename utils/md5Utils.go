package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// internal/md5 相关工具类

// []bytes获取md5的值

//

func GetMD5Value(data []byte) string {
	hash := md5.New()
	_, _ = hash.Write(data)
	// fmt.Println(n)
	sum := hash.Sum(nil)
	toString := hex.EncodeToString(sum)
	return toString
}

func GetMD5ValueFile(filepath string) (string, error) {
	bytedata, err := ReadFileToBytesWithMultiThread(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return "", err
	}
	value := GetMD5Value(bytedata)
	return value, nil
}

func IsValidMD5(data string) bool {
	if len(data) != 32 {
		return false
	}
	for _, char := range data {
		if (char < '0' || char > '9') && (char < 'a' || char > 'f') {
			return false
		}
	}
	return true
}
