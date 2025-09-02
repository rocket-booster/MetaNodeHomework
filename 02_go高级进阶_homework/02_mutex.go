package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {	
	var counter int64
	const (
		goroutines = 10  // 协程数量
		operations = 1000 // 每个协程的操作次数
	)

	var wg sync.WaitGroup
	wg.Add(goroutines)
	
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()			
			for j := 0; j < operations; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	
	wg.Wait()
	
	fmt.Printf("最终计数器值: %d\n", atomic.LoadInt64(&counter))
	fmt.Printf("预期值: %d\n", goroutines*operations)
}
    