package main

import (
  "crypto/rand"
  "fmt"
  "math/big"
)

func main() {
  // 生成 20 个 [0, 100) 范围的真随机数。
  for i := 0; i < 20; i++ {
    result, _ := rand.Int(rand.Reader, big.NewInt(100))
    fmt.Println(result)
  }
}