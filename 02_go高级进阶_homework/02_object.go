package main

import (
	"fmt"
)

type Person struct{
	Name string
	Age int32
}

type Employee struct{
	Person
	EmployeeID int32
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name:%s, Age:%d, EmployeeID:%d\n", e.Name, e.Age, e.EmployeeID)
}

func main(){
	var e Employee
	e.Name = "张三"
	e.Age = 30
	e.EmployeeID = 1001
	e.PrintInfo()
}