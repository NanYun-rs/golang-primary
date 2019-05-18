package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	age   int
	score float32
}

func main() {
	a := User{1, "i am a", 12, 12.33}

	//结构体字段可以通过结构体指针来访问。
	p := &a
	p.name = "i am p"

	// 结构体字段使用点号来访问
	// a.id = 2

	fmt.Println(a)
}
