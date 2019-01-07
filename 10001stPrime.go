package main

import "fmt"

func main() {
	var i = 1
	var num = 3
	for true {
		if isPrime(num) {
			i++
			if i == 10001 {
				fmt.Println(num)
				break
			}
		}
		num+=2
	}
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	for i:=2;i<n;i++{
		if n%i == 0 {
			return false
		}
	}
	return true
}
