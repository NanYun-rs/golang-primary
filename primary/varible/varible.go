package main

import (
	"fmt"
	"math"
)

// fmt 格式化操作包

func main() {
	const a, b int = 1, 2
	fmt.Println(a, b)
	const c bool = false
	fmt.Println(c)
	var d int
	fmt.Println(d)
	e := 12
	fmt.Println(e)

	// 按默认宽度和精度输出
	fmt.Printf("%f\n", math.Pi)
	// 按默认宽度 小数点后2位精度输出
	fmt.Printf("%.2f\n", math.Pi)

	// f := `namename1name2`
	// fmt.Printf(f)

	g := 'a'
	h := '你'

	fmt.Printf("%d %T\n", g, g)
	fmt.Printf("%d %T\n", h, h)

	i := make([]int, 3)
	i[0] = 1
	i[1] = 2
	i[2] = 3

	fmt.Println(i)

	j := 12
	k := "banner"

	fmt.Printf("%p %p", &j, &k)
}

/*
整形
- 按长度分 int8 int16 int32 int64
- 无符号整形 unit8 unit16 unit32 unit64

浮点型
- float32 最大范围 3.4e38 math.MaxFloat32
- float64 最大范围 1.8e308 math.MacFloat64

布尔型
- 只有 true false 不允许进行类型转换

字符串
- 字符串实现基于 UTF-8 编码
- 多行字符串 使用 ``
-	字符
	- unit8 byte 类型
	- rune 类型 代表一个UTF-8字符 用于处理中文、日文或者其他复合字符 rune 类型实际上是一个 int32

切片
- 拥有相同类型的可变长度的序列

指针
- 类型指针 允许对这个指针类型的数据进行修改，传递数据使用指针、而无需拷贝数据、类型指针不能进行偏移和运算
- 切片 由指向起始元素的起始指针、元素数量、容量构成

结构体
map
通道
*/
