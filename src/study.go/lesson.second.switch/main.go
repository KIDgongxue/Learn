package main

import "fmt"

func main()  {
	var money float64
	fmt.Println("我有", money , "元，中午只能")

	switch money{
	case 20:
		fmt.Println("点个外卖1")
	case 200:
		fmt.Println("去下馆子1")
	default:
		fmt.Println("在家做饭1")
	}

	switch {
	case money <= 20:
		fmt.Println("点个外卖")
		fallthrough
	case money <= 200:
		fmt.Println("去下馆子")
	default:
		fmt.Println("在家做饭")
	}
}
