package main

import (
	"fmt"
)

/*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func is_single(arr []int64) (int64) {
	m := make(map[int64]int, len(arr)+1)
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		}else {
			m[v]++
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}	
	return 0
}

func main(){
	var n int
	fmt.Print("输入数组长度n: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("无效长度")
		return
	}

	arr := make([]int64, n)
	fmt.Printf("输入%d个整数(空格分隔): ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("数组:", arr)

	println(is_single(arr))	
}
