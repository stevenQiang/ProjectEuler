package main

import (
	"math/big"
	"fmt"
	"strconv"
)

func main() {
	var f big.Int
	sum := 0.0
	// 阶乘的实现
	f.MulRange(1,100)
	value := f.String()
	for i := 0; i < len(value); i++ {
		data, _ := strconv.ParseFloat(string(value[i]), 10)
		sum += data
	}
	fmt.Println(sum)

}