package filex

// FilterHandler 过滤方法.
type FilterHandler func(values []string) ([]string, error)

// FilterEmptyHandler 不做任何处理.
func FilterEmptyHandler(values []string) ([]string, error) {
	return values, nil
}
