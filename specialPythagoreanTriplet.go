package main

import "fmt"

func main() {
	var a = 1000
	fmt.Println(getSpecialPythagorean(a))
}

func getSpecialPythagorean(n int) (int, int, int){
	a := 0
	for true {
		b := a + 1
		for; b<n; b++{
			c := n - a - b
			if a*a+b*b==c*c && c!=a && c!=b && a!=b {
				return a,b,c
			}
		}
		a++
	}
	return 0,0,0
}
