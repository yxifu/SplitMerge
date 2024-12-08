package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 	sizeFlag := flag.String("s", "1", "size 10K 10M")
	// 	flag.Parse()
	// 	fmt.Println("sizeFlag:", *sizeFlag)
	// }

	// func main1() {

	fmt.Println(os.Args)
	// i, err := parseSize("1M")
	// fmt.Println("1M", i, err)
	filePath := ""
	fmt.Println("=============")
	if len(os.Args) == 1 {
		fmt.Println("The parameter cannot be empty!")
		//fmt.Println("-m:  s:分割  m:合并")
		fmt.Println("-d:  true/false 默认:false (If the target file exists, delete the original file.)")
		fmt.Println("-s:  Size of the split files（10M 10K）")
		fmt.Println("<<file name>>: ")
		fmt.Println("Example 1: Splitting a File: SplitMerge -s=50M <<file name>>")
		fmt.Println("Example 2： merge files:  SplitMerge 文件名.s001")
		return
	}

	for i, s := range os.Args {
		if i > 0 && s[0:1] != "-" {
			filePath = s
		}
	}
	if filePath == "" {
		fmt.Println("filePath cannot be empty!")
		return
	}
	fmt.Println("filePath:", filePath)

	// 定义一个命令行参数，类型为字符串
	//modFlag := flag.String("m", "split", "mod：s/split;m/merge")
	sizeFlag := flag.String("s", "1", "size 10K 10M")
	fmt.Println("sizeFlag:", *sizeFlag)
	deleteFlag := flag.Bool("d", false, "If the target file exists, delete the original file")

	// 解析命令行参数
	flag.Parse()
	if *sizeFlag == "1" {
		//合并
		fmt.Println("merge")
		mergeFile(filePath, *deleteFlag)

	} else {
		//分割
		size, err := parseSize(*sizeFlag)
		if err != nil {
			fmt.Println("err:" + err.Error())
			return
		}
		fmt.Println("Size:", size)
		fmt.Println("Delete:", *deleteFlag)
		if size < 1024 {
			fmt.Println("Size至少1K", size)
			return
		}
		fmt.Println("being Split")
		SplitFile(filePath, size, filePath+".s")
		fmt.Println("END Split")

	}
}

// SplitFile Split the input file into multiple files according to the specified size.
func SplitFile(inputFilePath string, chunkSize int64, outputPrefix string) error {
	//var chunkSize int64 = 1024
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	if fileSize <= chunkSize {
		// If the file size is less than or equal to the chunk size, no splitting is required.。
		return nil
	}

	var count int
	var ts int64 = 0
	for {
		// create output file
		outputFileName := fmt.Sprintf("%s%03d", outputPrefix, count+1)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		reader := io.LimitReader(inputFile, chunkSize)

		_, err = io.Copy(outputFile, reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		count++
		ts += chunkSize
		if ts >= fileSize {
			break
		}
	}

	return nil
}

func mergeFile(firstFilePath string, deleteFlag bool) error {

	outputFile := removeSuffix(firstFilePath, ".s001")
	if outputFile == firstFilePath {
		fmt.Println("The file must be the first file(*.s001)")
		return nil
	}
	fmt.Println("outputFile:", outputFile)
	if CheckFileExist(outputFile) {
		if deleteFlag {
			// 如果文件存在，则删除它。
			if err := os.Remove(outputFile); err != nil {
				fmt.Println("Failed to delete the file:" + outputFile)
				return err
			}
		} else {
			fmt.Println("文件已经存在！" + outputFile)
			return nil
		}
	}

	// 创建输出文件。
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var count int = 0
	for {
		// 创建输出文件。
		filename := fmt.Sprintf("%s%s%03d", outputFile, ".s", count+1)

		if !CheckFileExist(filename) {
			break
		}

		// 打开每个文件。
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		// 将文件内容复制到输出文件。
		_, err = io.Copy(outFile, file)
		if err != nil {
			return err
		}
		count++
	}

	return nil

}

// CheckFileExist 检查文件是否存在。
func CheckFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在
			return false
		}
		// 其他错误
		fmt.Println("Error:", err)
		return false
	}
	// 文件存在
	return true
}
func removeSuffix(s, suffix string) string {
	// 使用 strings.HasSuffix 检查字符串是否以指定的后缀结尾
	if strings.HasSuffix(s, suffix) {
		// 使用切片去掉后缀
		return s[:len(s)-len(suffix)]
	}
	// 如果字符串不以该后缀结尾，则返回原始字符串
	return s
}
func parseSize(sizeStr string) (int64, error) {
	var size int64
	var unit string
	var factor int64

	// Extract the number part and the unit suffix.
	for i := len(sizeStr) - 1; i >= 0; i-- {
		if '0' <= sizeStr[i] && sizeStr[i] <= '9' {
			continue
		} else if sizeStr[i] == 'B' { // Skip 'B'
			continue
		} else {
			unit = string(sizeStr[i])
			break
		}
	}

	// Parse the number part.
	numberPart := sizeStr[:len(sizeStr)-len(unit)]
	value, err := strconv.ParseInt(numberPart, 10, 64)
	if err != nil {
		return 0, err
	}
	size = value

	// Determine the factor based on the unit.
	switch unit {
	case "K", "k":
		factor = 1024
	case "M", "m":
		factor = 1024 * 1024
	case "G", "g":
		factor = 1024 * 1024 * 1024
	default:
		factor = 1
	}

	return size * factor, nil
}
