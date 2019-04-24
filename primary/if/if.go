package main

import (
	"fmt"
)


func main()  {
	i:=3

	if i==3{
		fmt.Println("i == 3")
	}

	if i>3 {
		fmt.Println(" i more then 3")
	}else if i==3{
		fmt.Println("i eqal 3")
	}else{
		fmt.Println("i less then 3")
	}
}