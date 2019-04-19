package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	// 伪随机数 rand.Seed 设置种子
	// 真随机数 crypto/rand 
	rand.Seed(int64(time.Now().UnixNano()))
	fmt.Println(rand.Intn(100))
}