package main

import "fmt"

func main() {
	m1 := map[string]int{"小强":20}
	fmt.Println(m1)
	m1["小强"]=28
	fmt.Println(m1)
	val1,flag := m1["小红"]
	fmt.Println("小红的年龄",val1,flag)

	var m2 map[string]int
	m2 = make(map[string]int)
	m2["小强"] = 1
	fmt.Println(m2)

	var m3 = make(map[string]int)
	m3["小强"] = 1
	fmt.Println(m3)
}
