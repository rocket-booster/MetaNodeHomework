package main

import (
	"fmt"
)

func validParentheses(s string) bool {
	bracketMap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
	
	stack := []rune{}
    for _, char := range s {
        if correspondingLeft, ok := bracketMap[char]; ok {            
            if len(stack) == 0 || stack[len(stack)-1] != correspondingLeft {
                return false
            }            
            stack = stack[:len(stack)-1]
        } else {            
            stack = append(stack, char)
        }
    }
        
    return len(stack) == 0
}

func main() {
	var s string 
	println("请输入一个字符串:")
	_, err := fmt.Scan(&s)
	if err != nil {
		println("输入错误:", err.Error())
		return
	}
	if validParentheses(s) {	
		println("有效")
	} else {
		println("无效")
	}
}	