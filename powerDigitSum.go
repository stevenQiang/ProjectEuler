package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(getPowerDigit(2, 1000))
}

func getPowerDigit(x float64 ,y float64) int64 {
	digitSum := math.Pow(x,y)
	sum := int64(0)
	// float64转string, f代表没有指数，1是精度
	for _,v := range strconv.FormatFloat(digitSum,'f', 1, 64) {
		if int64(v - '0') > 0{
			sum += int64(v - '0')
		}
	}
	return sum
}