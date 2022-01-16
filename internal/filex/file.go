package filex

import (
	"github.com/iokr/filebot/core/filex"
)

// CopyInputToOutputFile 拷贝input文件内容到output文件.
func CopyInputToOutputFile(inputFilePath, outputFilePath string) error {
	return InputFilterToOutputFile(inputFilePath, outputFilePath, FilterEmptyHandler)
}

// InputFilterToOutputFile input文件过滤内容到output文件.
func InputFilterToOutputFile(inputFilePath, outputFilePath string, filterHandler FilterHandler) error {
	values, err := filex.ReadLine(inputFilePath)
	if err != nil {
		return err
	}

	values, err = filterHandler(values)
	if err != nil {
		return err
	}

	return filex.WriteLineByTruncate(outputFilePath, values)
}
