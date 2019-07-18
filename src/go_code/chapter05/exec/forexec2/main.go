package main

import (
	"fmt"
)

func main() {
	var total float64
	var passcount int32
	for j := 1; j <= 3; j++ {
		var sum float64
		var passnum int32
		for i := 1; i <= 5; i++ {
			var score float64
			fmt.Printf("输入%d班第%d名学生的分数:\n", j, i)
			fmt.Scanln(&score)
			if score >= 60 {
				passnum++
				passcount++
			}
			sum += score

		}
		total += sum
		fmt.Printf("%v班的平均分为:%v，及格人数为:%d\n", j, sum/5, passnum)
	}
	fmt.Printf("班级总分：%v\n及格人数：%d", total, passcount)

}
