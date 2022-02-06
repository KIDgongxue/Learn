package main

import "fmt"

func main()  {
	var personInfos = [...][3]string{
		[3]string{"小强","男","28"},
		[3]string{"小红","女","28"},
		[3]string{"小明","男","28"},
		[3]string{"小将","男","28"},
	}
	for _,val1 := range personInfos{
		for _,val2 := range val1{
		fmt.Println("数组：",val1,"的信息", val2)
		}
	}
}
