package main

import (
	"fmt"
	"strconv"
)

func main(){
	var i,j int
	var maxPalindrome = 0
	var palindromeList []int
	for i=100;i<1000;i++{
		for j=100;j<1000;j++{
			palindrome := i*j
			if isPalindrome(strconv.Itoa(palindrome)) {
				fmt.Println(palindrome)
				if palindrome > maxPalindrome {
					maxPalindrome = palindrome
				}
				palindromeList = append(palindromeList, palindrome)
			}
		}
	}
	fmt.Println(palindromeList)
	fmt.Println(maxPalindrome)
}


func isPalindrome(n string) bool {
	l := len(n)
	left := 0
	right := l-1
	for left < right {
		if n[left] != n[right] {
			return false
		}
		left++
		right--
	}
	return true
}
