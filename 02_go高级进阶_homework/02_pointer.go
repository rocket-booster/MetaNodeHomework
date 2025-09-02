package main

import (
	"fmt"
)

func solution(p *[]int) {
	for i := 0; i < len(*p); i++ {
		(*p)[i] = (*p)[i]*2
	}
}

func main() {	
	s := []int{1,2,3,4,5}
	fmt.Println("s:", s)

	solution(&s)
	fmt.Println("结果:", s)
}
	