package main

import "fmt"

func main()  {
	var money float64
	fmt.Println("我有", money , "元，中午只能")
	if money <= 0 {
		fmt.Println("在家做饭")
	} else if money <= 20 {
		fmt.Println("点个外卖")
	} else if money <= 200 {
		fmt.Println("下个馆子")
	} else if money <= 2000 {
		fmt.Println("去米其林")
	} else if money <= 20000 {
		fmt.Println("去新马泰")
	} else {
		fmt.Println("容我想想")
	}

}
