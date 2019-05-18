package  main

import (
	"fmt"
)


// type IntSlice []int

// func (p IntSlice)  Swap(i,j int){
// 	p[i],p[j] = p[j],p[i]
// }


func GetData()(int,int)  {
	return 100,200
}


func main()  {
	attack := 40
	defence :=20
	var damageRate float32 = 0.17
	damage:= float32(attack - defence) * damageRate
	fmt.Println(damage)

	a,b := 100,200
	a,b = b,a
	fmt.Println("a",a,"b",b)




	c,_ := GetData()
	_,d := GetData()

	fmt.Println(c,d)
}