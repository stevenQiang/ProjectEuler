/*
可以简化为fibonacci
f(x,y) = {
	0    (x=0 && y=0)
	1 	 (x=0 || y=0)
	f(x-1,y)+f(x,y-1)
}
 */
package main

import "fmt"

func main() {
	var grid = 20
	fmt.Println(getRoutes(grid, grid))
}

func getRoutes(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	} else if x == 0 || y == 0 {
		return 1
	} else {
		return getRoutes(x-1, y)+getRoutes(x,y-1)
	}
}