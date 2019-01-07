package main

import "fmt"

func main(){
	var (
		i = 0
		sum = 0
	)
	for true {
		var value = fib(i)
		if value > 4000000 {
			break
		}
		if value %2 == 0 {
			sum += value
		}
		i += 1
	}
	fmt.Println(sum)
}

func fib(n int) int {
	if n <= 1{
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}