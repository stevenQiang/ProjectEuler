package main

import "fmt"

func main() {
	var n = 100
	fmt.Println(squareOfSum(n) - sumOfSquare(n))
}

func sumOfSquare(n int) int {
	sum := 0
	for i:=1;i<=n;i++ {
		sum += i*i
	}
	return sum
}

func squareOfSum(n int) int {
	sum := 0
	for i:=1;i<=n;i++ {
		sum += i
	}
	return sum*sum
}
