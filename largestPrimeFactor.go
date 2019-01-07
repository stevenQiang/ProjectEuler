// this solution is very slow, must have so quick
package main

import "fmt"

func main(){
	var primeList []int
	var number = 600851475143
	primeList = primeNumberList(number)
	fmt.Println(primeList)
}

func primeNumberList(n int) []int {
	var primeList []int
	for i:=3; i<= n/2; i+=2 {
		if isPrimeNumber(i) && n % i == 0 {
			primeList = append(primeList, i)
		}
	}
	return primeList
}

func isPrimeNumber(n int) bool {
	i := 3
	for ; i < n; i+=2 {
		if n%i == 0 {
			break
		}
	}
	if i < n {
		return false
	}
	return true
}
