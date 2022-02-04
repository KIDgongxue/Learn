package main

import "fmt"

/**
判断两条线是否平行
 */
func main()  {
	var x1,y1 float64 = 4.3,5.6
	var x2,y2 float64 = 4.3,5.6
	var slope1 float64 = y1/x1
	var slope2 float64 = y2/x2
	if slope1 == slope2 {
		fmt.Println("直线(",x1,",",y1,")与直线(" ,x2,",",y2,")","平行")
	} else {
		fmt.Println("直线(",x1,",",y1,")与直线(" ,x2,",",y2,")","不平行")
	}
}
