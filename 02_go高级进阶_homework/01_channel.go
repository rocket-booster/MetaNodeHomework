package main

import (
	"fmt"	
	"sync"
)

func  main(){
	dataChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=1; i<=10; i++{
			fmt.Println("生成数据:", i)
			dataChan <- i
		}
		close(dataChan)
	}()
	wg.Add(1)
	go func(){
		defer wg.Done()
		for num := range dataChan{
			fmt.Println("接收数据:", num)
		}
	}()
	wg.Wait()
}