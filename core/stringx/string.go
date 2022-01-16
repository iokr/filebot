package stringx

// RemoveBlankLine 删除空白行.
func RemoveBlankLine(values []string) []string {
	s := make([]string, 0, len(values))
	for idx := range values {
		if len(values[idx]) == 0 {
			continue
		}
		s = append(s, values[idx])
	}
	return s
}

// RemoveDuplicate 删除重复的.
//  去重后,原顺序不变.
func RemoveDuplicate(values []string) []string {
	valueMap := make(map[string]struct{})
	s := make([]string, 0, len(values))
	for idx := range values {
		if _, ok := valueMap[values[idx]]; ok {
			continue
		}
		s = append(s, values[idx])
		valueMap[values[idx]] = struct{}{}
	}
	return s
}

// RemoveBlankLineDuplicate 内容删除空行后去重.
//  去重后,原顺序不变.
func RemoveBlankLineDuplicate(values []string) []string {
	valueMap := make(map[string]struct{})
	s := make([]string, 0, len(values))
	for idx := range values {
		if values[idx] == "" {
			continue
		}

		if _, ok := valueMap[values[idx]]; ok {
			continue
		}
		s = append(s, values[idx])
		valueMap[values[idx]] = struct{}{}
	}
	return s
}
