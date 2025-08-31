package main
/*
题目：判断一个整数是否是回文数
*/
import (
	"fmt"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var s string 
	println("请输入一个字符串:")
	_, err := fmt.Scan(&s)
	if err != nil {
		println("输入错误:", err.Error())
		return
	}
	if isPalindrome(s) {	
		println("是回文")
	} else {
		println("不是回文")
	}
}	