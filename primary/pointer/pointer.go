package main

import (
	"fmt"
)

func main() {

	i, j := 100, 1000

	a := &i
	b := &j

	fmt.Println(*a)
	fmt.Println(*b)

	*a = 1
	*b = 11

	fmt.Println(i)
	fmt.Println(j)

}


// Go 具有指针。 指针保存了变量的内存地址。
// 类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。
// & 符号会生成一个指向其作用对象的指针。