// 如果d(a) = b且d(b) = a，且a ≠ b，那么a和b构成一个亲和数对，a和b被称为亲和数。

package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i<=10000; i++ {
		a := d(i)
		b := d(a)
		if i == b && a!=b {
			sum += a
		}
	}
	fmt.Println(sum)
}

func d(a int) int{
	result := 0
	for i := 1; i<= a/2; i++ {
		if a%i == 0 {
			result += i
		}
	}
	return result
}
