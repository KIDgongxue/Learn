package main

import "fmt"

/**
计算圆的面积并保留3位小数
 */
func main()  {
	const pai= 3.1415926
	var radius float64 = 3.2
	var area float64 = radius * radius * pai
	fmt.Println("未格式化", area)
	fmt.Printf("保留3为小数:%.3f", area)
}