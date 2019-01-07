package main

import "fmt"

func main(){
	var n = 2000000
	fmt.Println(sumPrimes(n))
}

func sumPrimes(n int) int {
	sum := 2
	for i:=3;i<=n;i+=2{
		if isPrimeNumber(i) {
			sum += i
		}
	}
	return sum
}

func isPrimeNumber(n int) bool {
	i := 3
	for ; i < n; i+=2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

