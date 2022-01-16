package main

import (
	"log"

	"github.com/iokr/filebot/core/stringx"
	"github.com/iokr/filebot/internal/filex"
)

func main() {
	inputFilePath := "./files/input.txt"
	outputFilePath := "./files/output.txt"

	log.Println(InputAddQuoteToOutputFile(inputFilePath, outputFilePath))
}

// InputFilterToOutputFile input文件过滤内容到output文件.
func InputFilterToOutputFile(inputFilePath, outputFilePath string) error {
	return filex.InputFilterToOutputFile(inputFilePath, outputFilePath, func(values []string) ([]string, error) {
		// todo 自定义过滤方式处理.
		return values, nil
	})
}

// InputAddQuoteToOutputFile input文件内容添加双引号后写入output文件.
//  eg: 111  => "111".
func InputAddQuoteToOutputFile(inputFilePath, outputFilePath string) error {
	return filex.InputFilterToOutputFile(inputFilePath, outputFilePath, func(values []string) ([]string, error) {
		newValues := make([]string, 0, len(values))
		for idx := range values {
			if len(values[idx]) == 0 {
				continue
			}
			newValues = append(newValues, stringx.FormatToDoubleQuote(values[idx]))
		}
		return newValues, nil
	})
}
