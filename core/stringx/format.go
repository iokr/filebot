package stringx

import (
	"fmt"
	"strings"
)

// FormatHandler 格式化方法.
type FormatHandler func(value string) string

// FormatEmpty 不做任何处理.
func FormatEmpty(value string) string {
	return value
}

// FormatToBackQuote 格式化成反引号.
//  eg: (11 => `11`).
func FormatToBackQuote(value string) string {
	return strings.ReplaceAll(fmt.Sprintf("\"%s\"", value), "\"", "`")
}

// FormatToBackQuoteBeforeComma 格式化成反引号在逗号之前.
//  eg: (11 => `11`,).
func FormatToBackQuoteBeforeComma(value string) string {
	return strings.ReplaceAll(fmt.Sprintf("\"%s\",", value), "\"", "`")
}

// FormatToBackQuoteAfterComma 格式化成反引号在逗号之后.
//  eg: (11 => ,`11`).
func FormatToBackQuoteAfterComma(value string) string {
	return strings.ReplaceAll(fmt.Sprintf(",\"%s\"", value), "\"", "`")
}

// FormatToSingleQuote 格式化成单引号.
//  eg: (11 => '11').
func FormatToSingleQuote(value string) string {
	return fmt.Sprintf(`'%s'`, value)
}

// FormatToSingleQuoteBeforeComma 格式化成单引号在逗号之前.
//  eg: (11 => '11',).
func FormatToSingleQuoteBeforeComma(value string) string {
	return fmt.Sprintf(`'%s',`, value)
}

// FormatToSingleQuoteAfterComma 格式化成单引号在逗号之后.
//  eg: (11 => ,'11').
func FormatToSingleQuoteAfterComma(value string) string {
	return fmt.Sprintf(`,'%s'`, value)
}

// FormatToDoubleQuote 格式化成双引号.
//  eg: (11 => "11").
func FormatToDoubleQuote(value string) string {
	return fmt.Sprintf("\"%s\"", value)
}

// FormatToDoubleQuoteBeforeComma 格式化成双引号在逗号之前.
//  eg: (11 => "11",).
func FormatToDoubleQuoteBeforeComma(value string) string {
	return fmt.Sprintf("\"%s\",", value)
}

// FormatToDoubleQuoteAfterComma 格式化成双引号在逗号之后.
//  eg: (11 => ,"11").
func FormatToDoubleQuoteAfterComma(value string) string {
	return fmt.Sprintf(",\"%s\"", value)
}
