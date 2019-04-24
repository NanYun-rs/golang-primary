package main

import (
	"fmt"
	"time"
)



func main()  {
	switch time.Now().Weekday(){
	case time.Saturday , time.Sunday:
		fmt.Println("today is weekday")
	default:
		fmt.Println("today is workday")
	}

// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式
	t :=time.Now()
	switch {
	case t.Hour()>12:
		fmt.Println("current is afternonne")
	default:
		fmt.Println("current is morning")
	}
}