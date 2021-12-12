package main

import "fmt"

func main()  {
	const circ = 3.1415926
	var banjing float64 = 1.23
	var result = circ*banjing*banjing
	fmt.Println("%.2f",result)
	var resultF = fmt.Sprintf("%.3f",result)
	fmt.Println(resultF)
}