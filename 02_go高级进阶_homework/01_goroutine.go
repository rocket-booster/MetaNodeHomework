package main

import (
	"fmt"
)

func P1to10(c chan<- int){
	for i:=1; i<=10; i++{
		if i%2 ==0 {
			continue
		}
		fmt.Println(i)
	}
	c<-1
}

func P2to10(c chan<- int){
	for i:=2; i<=10; i++{
		if i%2 ==1 {
			continue
		}
		fmt.Println(i)
	}
	c<-1
}

func main(){
	end := make(chan int)
	end_count := 0
	go P1to10(end)
	go P2to10(end)

	end_count += <-end
	end_count += <- end
	if (end_count == 2){
		fmt.Println("结束")
	}
}