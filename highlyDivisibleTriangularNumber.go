package main

import "fmt"

func main() {
	var n = 500
	getTriangular(n)
}

func getTriangular(n int){
	for i:=1;true;i++{
		sum := 0
		for j:=1;j<=i;j++{
			sum += j
		}
		if getDivisors(sum) >= n {
			fmt.Println(sum)
			break
		}
	}
}

func getDivisors(n int) int {
	var divisorList []int
	for i:=1;i<=n;i++{
		if n%i == 0 {
			divisorList = append(divisorList, i)
		}
	}
	return len(divisorList)
}