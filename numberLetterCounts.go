package main

import (
	"github.com/divan/num2words"
	"strings"
	"fmt"
)


func main() {
	var sum = 0
	for i:= 1;i<=1000;i++ {
		sum += getWord(i)
	}
	fmt.Println(sum)
}

func getWord(n int) int {
	word := num2words.ConvertAnd(n)
	// 去掉空格及特殊字符
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)
	return len(word)
}
