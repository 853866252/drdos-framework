package utils

import (
	"bufio"
	"fmt"
	"os"
)

// FileLoads 加载文件
func FileLoads(path string) ([]string, error) {
	result := make([]string, 0)

	// 判断文件是否存在
	flag, err := PathExists(path)
	if flag == false {
		return result, err
	}

	// 读取文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("[!] Open \"%s\" failed, err: [%v]\n", path, err)
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && IPCheck(line) {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[!] Cannot scanner file: %s, err: [%v]\n", path, err)
		return result, err
	}

	return result, nil
}

func FileWrites(path string, content string) error {
	fd, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 追加写入
	defer fd.Close()
	if err != nil {
		return err
	}
	_, err = fd.WriteString(content + "\n")
	if err != nil {
		fmt.Printf("[!] FileWrites err : [%v] + \n", err)
		return err
	}
	return nil
}
