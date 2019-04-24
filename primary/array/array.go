package main

import (
	"fmt"
)


func main()  {
	var a [4] int
	fmt.Println('a',a)

	a[3] =1
	fmt.Println("set number",a)
	fmt.Println("get number",a[3])

	fmt.Println("length",len(a))

	b:= [5]int {1,2,3,4,5}

	fmt.Println("b",b)

	// 二维数组
	twoD :=[2][3]int{{1,2,3},{4,5,6}}
	fmt.Println("two d",twoD)

	var twoDD [2][3] int

	// 遍历赋值
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoDD[i][j] = i+j
		}
	}

	fmt.Println("for towdd",twoDD)
}