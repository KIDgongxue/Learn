package main

import (
	"fmt"
)

/*
BMI = 体重(公斤) / (身高米 * 身高米)
体脂率 ： 1.2 * BMI + 0.23 * 年龄 - 5.4 - 10.8 * 性别 (男 = 1, 女 = 0)
作业要求：连续输入三人的姓名、性别、身高、体重、年龄信息
输出：
每个人的姓名、BMI、体脂率、建议
总人数、平均体脂率
*/
func main() {
	var names = [3]string{}
	var weights = [3]float64{}
	var heights = [3]float64{}
	var sexs = [3]string{}
	var ages = [3]int{}
	var i int = 0
	fmt.Println("请输入第3位测算人的信息")
	for i < 3 {
		fmt.Print("请输入测算人姓名：")
		fmt.Scanln(&names[i])
		fmt.Print("请输入体重(公斤)：")
		fmt.Scanln(&weights[i])
		fmt.Print("请输入身高(米)：")
		fmt.Scanln(&heights[i])
		for heights[i] >= 3 {
			fmt.Print("身高输入不符合常规判断(<3米)，请重新输入身高(米)：")
			fmt.Scanln(&heights[i])
		}
		fmt.Print("请输入性别(男/女)：")
		fmt.Scanln(&sexs[i])
		for sexs[i] != "男" && sexs[i] != "女" {
			fmt.Print("身高输入不符合规定，请重新输入性别(男/女)：")
			fmt.Scanln(&sexs[i])
		}
		fmt.Print("请输入年龄：")
		fmt.Scanln(&ages[i])
		for ages[i] < 18 {
			fmt.Print("测算人年龄不满足最低测试标准，请重新输入年龄(>=18)：")
			fmt.Scanln(&ages[i])
		}
		i++
	}
	var sumFatRate float64
	for j := 0; j < i; j++ {
		var sex = sexs[j]
		var weight = weights[j]
		var height = heights[j]
		var age = ages[j]
		var name = names[j]
		var sexWeight int
		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 2
		}
		var bmi float64 = weight / (height * height)
		fmt.Print(name, ",BMI为：", bmi, ",")
		var fatRate float64 = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)
		fmt.Print("体脂率为：", fatRate, ",")
		sumFatRate = sumFatRate + fatRate

		if age >= 18 && age <= 39 {
			if sex == "男" {
				if fatRate <= 10 {
					fmt.Println("目前是：偏瘦，需多吃多锻炼，增强体质")
				} else if fatRate <= 16 {
					fmt.Println("目前是：标准，继续保持")
				} else if fatRate <= 21 {
					fmt.Println("目前是：偏胖，吃完饭多散步，消化消化")
				} else if fatRate <= 26 {
					fmt.Println("目前是：肥胖，少吃多运动")
				} else {
					fmt.Println("目前是：严重肥胖，增加锻炼刻不容缓")
				}
			} else {
				if fatRate <= 20 {
					fmt.Println("目前是：偏瘦，需多吃多锻炼，增强体质")
				} else if fatRate <= 27 {
					fmt.Println("目前是：标准，继续保持")
				} else if fatRate <= 34 {
					fmt.Println("目前是：偏胖，吃完饭多散步，消化消化")
				} else if fatRate <= 39 {
					fmt.Println("目前是：肥胖，少吃多运动")
				} else {
					fmt.Println("目前是：严重肥胖，增加锻炼刻不容缓")
				}
			}
		} else if age >= 40 && age <= 59 {
			if sex == "男" {
				if fatRate <= 11 {
					fmt.Println("目前是：偏瘦，需多吃多锻炼，增强体质")
				} else if fatRate <= 17 {
					fmt.Println("目前是：标准，继续保持")
				} else if fatRate <= 22 {
					fmt.Println("目前是：偏胖，吃完饭多散步，消化消化")
				} else if fatRate <= 27 {
					fmt.Println("目前是：肥胖，少吃多运动")
				} else {
					fmt.Println("目前是：严重肥胖，增加锻炼刻不容缓")
				}
			} else {
				if fatRate <= 21 {
					fmt.Println("目前是：偏瘦，需多吃多锻炼，增强体质")
				} else if fatRate <= 28 {
					fmt.Println("目前是：标准，继续保持")
				} else if fatRate <= 35 {
					fmt.Println("目前是：偏胖，吃完饭多散步，消化消化")
				} else if fatRate <= 40 {
					fmt.Println("目前是：肥胖，少吃多运动")
				} else {
					fmt.Println("目前是：严重肥胖，增加锻炼刻不容缓")
				}
			}
		} else {
			if sex == "男" {
				if fatRate <= 13 {
					fmt.Println("目前是：偏瘦，需多吃多锻炼，增强体质")
				} else if fatRate <= 19 {
					fmt.Println("目前是：标准，继续保持")
				} else if fatRate <= 24 {
					fmt.Println("目前是：偏胖，吃完饭多散步，消化消化")
				} else if fatRate <= 29 {
					fmt.Println("目前是：肥胖，少吃多运动")
				} else {
					fmt.Println("目前是：严重肥胖，增加锻炼刻不容缓")
				}
			} else {
				if fatRate <= 22 {
					fmt.Println("目前是：偏瘦，需多吃多锻炼，增强体质")
				} else if fatRate <= 29 {
					fmt.Println("目前是：标准，继续保持")
				} else if fatRate <= 36 {
					fmt.Println("目前是：偏胖，吃完饭多散步，消化消化")
				} else if fatRate <= 41 {
					fmt.Println("目前是：肥胖，少吃多运动")
				} else {
					fmt.Println("目前是：严重肥胖，增加锻炼刻不容缓")
				}
			}
		}
	}
	fmt.Print("总测算人数：", i, "人,")
	fmt.Printf("平均体脂率：%.2f", sumFatRate/3/100)
}
