package main

import (
	"fmt"
)

func plusPointObjectOne(p *int) {
	*p = *p + 10
}

func main() {	
	var n int
	fmt.Println("请输入整数:")
	_, err := fmt.Scan(&n)
	if err != nil {
		println("输入错误:", err.Error())
		return
	}

	plusPointObjectOne(&n)
	fmt.Println("结果:", n)
}
	