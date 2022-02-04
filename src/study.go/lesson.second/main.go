package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var hello string = "hello,string"
	fmt.Println(hello)

	var sex = "女"
	fmt.Println(sex)

	var hg,vgr float64 = 3,4.24
	var sc = hg * vgr
	fmt.Println(sc)

	var hg1,vgr1 = 3,4.24
	var sc1 = hg1 * int(vgr1)
	fmt.Println(sc1)
	fmt.Println(reflect.TypeOf(sc) , "," , reflect.TypeOf(sc1))

	fmt.Println("异或运算")
	a,b :=  100,31
	fmt.Println(a ^ b)
	fmt.Println(b ^ a)

	fmt.Printf("我的名字%s\n","小强")
	fmt.Printf("我的名字%q\n","小强")
	fmt.Printf("我的名字%x\n","小强")
}
