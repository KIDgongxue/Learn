package main

import "fmt"

func main()  {
	var slice = []int{}
	slice = append(slice,1)
	fmt.Println(slice)
	slice = append(slice,2,3,4,5)
	fmt.Println(slice)
	slice = append(slice[:1],slice[2:4]...)
	fmt.Println(slice)

	var slice2 = make([]int,2)
	fmt.Println("slice长度：",len(slice2),"值：", slice2)
	slice2 = append(slice2,1)
	fmt.Println("slice长度：",len(slice2),"值：", slice2)
}
