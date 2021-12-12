package main

import "fmt"

func main()  {
	fmt.Println("你好，golang")
	var age, length float64= 12,111
	fmt.Println("age=" , age, "length=",length)
	var ageInt, lengthFloat = 12,111.13
	fmt.Println("ageInt=" , ageInt, "lengthFloat=",lengthFloat)
	var sex,name = 1,"xiaoli"
	fmt.Println("name=",name,"sex=",sex)
	var a,b int = 30,10 // 如为int8，会造成溢出
	fmt.Println(a*b)

}