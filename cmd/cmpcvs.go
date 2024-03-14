package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func compare(leftFile, rightFile string) {
	var leftSum, rightSum float64

	// 打开 CSV 文件
	left, _ := readFile(leftFile)
	right, _ := readFile(rightFile)

	// 打印读取的数据
	leftMap := make(map[string]float64)
	for _, record := range left {
		num, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			fmt.Println(err)
		}
		leftMap[record[0]] = num
		leftSum += num
	}

	rightMap := make(map[string]float64)
	for _, record := range right {
		num, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			fmt.Println(err)
		}
		rightMap[record[0]] = num
		rightSum += num
	}
	for k, v := range leftMap {
		if v2, ok := rightMap[k]; ok {
			if v != v2 {
				fmt.Printf("%s not equals, %s vs %s\n", k, v, v2)
			}
		} else {
			fmt.Printf("%s not found\n", k)
		}
	}

	fmt.Println(leftSum, rightSum)
}

func readFile(fileName string) ([][]string, bool) {
	homeDir, err := os.UserHomeDir()
	file, err := os.Open(filepath.Join(homeDir, "Downloads", fileName))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, true
	}
	defer file.Close()

	// 创建一个 CSV Reader
	reader := csv.NewReader(file)

	// 读取 CSV 文件中的数据
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil, true
	}
	return records, false
}
