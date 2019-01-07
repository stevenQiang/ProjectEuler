package main

import "fmt"

func main() {
	var n = 20
	var maxNumber = smallestPositive(n)
	fmt.Println(maxNumber)
}

func smallestPositive(n int) int {
	j := n
	for true {
		j += n
		isPositive := true
		for i:=n/2;i<=n;i++{
			if j%i != 0 {
				isPositive = false
				break
			}
		}
		if isPositive {
			break
		}
	}
	return j
}
