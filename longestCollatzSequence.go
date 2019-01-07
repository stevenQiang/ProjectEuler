// 这一题是求满足两个条件下面，最大的递归的数量
package main

import "fmt"

func main(){
	var n = 1000000
	fmt.Println(getLongest(n))
}

func getLongest(n int) (int,int) {
	sum := 0
	maxNumber := 0
	for i:=1;i<=n;i++{
		number := i
		a := 0
		for number!=1 {
			if number%2 == 0 {
				number = number / 2
			}else{
				number = number*3 + 1
			}
			a+=1
		}

		if a > sum {
			sum = a
			maxNumber = i
		}
	}
	return sum, maxNumber
}