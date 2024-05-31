package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 文件工具

func saveMD5AsHexFile(filePath string, md5Value []byte) error {
	hexFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating hex file: %w", err)
	}
	defer func(hexFile *os.File) {
		err := hexFile.Close()
		if err != nil {
			fmt.Printf("error closing hex file: %w", err)
		}
	}(hexFile)

	writer := bufio.NewWriter(hexFile)
	_, err = writer.Write(md5Value)
	if err != nil {
		return fmt.Errorf("error writing to hex file: %w", err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing hex file: %w", err)
	}

	return nil
}

// HandleHexFlag 读取文件并计算MD5值，并将结果保存为HEX文件
func HandleHexFlag(sourceFilePath string) {
	// byteContent, err := readFileToBytesTree(sourceFilePath)
	byteContent, err := ReadFileToBytesWithMultiThread(sourceFilePath)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	md5Value := GetMD5Value(byteContent)
	fmt.Printf("MD5 value of file '%s' is: %s\n", sourceFilePath, md5Value)
	// md5Value := md5Util.GetMD5ValueFromFile(filePath)

	// hexFilePath := filepath.Base() + ".hex"
	hexfilepath := filepath.Join(filepath.Dir(sourceFilePath), GenerateRandomString(6)+".hex")
	err = saveMD5AsHexFile(hexfilepath, []byte(md5Value))
	if err != nil {
		fmt.Printf("Error saving MD5 as hex file: %v\n", err)
		return
	}

	fmt.Printf("MD5 value saved to '%s'\n", hexfilepath)
}

func ReadFileToBytesWithMultiThread(filePath string) ([]byte, error) {
	now := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %w", err)
		}
	}(file)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var buffer bytes.Buffer

	chunkSize := 10 * 1024 * 1024 // 10MB chunk size
	maxRead := 200 * 1024 * 1024  // 200MB max read limit

	reader := io.LimitReader(file, int64(maxRead))

	// Read and process chunks in a loop
	buf := make([]byte, chunkSize)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			wg.Add(1)
			go func(chunk []byte) {
				mutex.Lock()
				defer mutex.Unlock()
				defer wg.Done()
				buffer.Write(chunk)
			}(buf[:n])
		}

		if err != nil {
			if err == io.EOF {
				break // End of file reached
			} else {
				return nil, fmt.Errorf("error reading file: %w", err)
			}
		}
	}

	wg.Wait()
	fmt.Println("Time taken:", time.Since(now))
	return buffer.Bytes(), nil
}
