package main

import (
	"fmt"
)

// longestCommonPrefix 查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 处理空数组情况
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准
	for i := 0; i < len(strs[0]); i++ {
		// 取出当前字符
		c := strs[0][i]

		// 与其他字符串的对应位置比较
		for j := 1; j < len(strs); j++ {
			// 如果超出了某个字符串的长度，或者字符不匹配，返回当前前缀
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	// 如果第一个字符串是所有字符串的前缀
	return strs[0]
}

func main() {
	// 测试用例
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"interspecies", "interstellar", "interstate"},
		{"", "b"},
		{"a"},
		{"ab", "a"},
	}

	for _, testCase := range testCases {
		fmt.Printf("字符串数组: %v\n", testCase)
		fmt.Printf("最长公共前缀: \"%s\"\n\n", longestCommonPrefix(testCase))
	}
}