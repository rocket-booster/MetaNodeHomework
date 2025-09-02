package main

import (
	"fmt"
	"sync"
)

func  main(){
	var counter int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	const (
		goroutines = 10  // 协程数量
		operations = 1000 // 每个协程的操作次数
	)
	wg.Add(goroutines)
	for i:=0; i<goroutines; i++{
		go func(counter *int64){
			defer wg.Done()
			for j:=0; j<operations; j++{				
				mu.Lock()
				(*counter)++
				mu.Unlock()								
			}
		}(&counter)
	}
	wg.Wait()	
	fmt.Printf("最终计数器值: %d\n", counter)
	fmt.Printf("预期值: %d\n", goroutines*operations)
}
