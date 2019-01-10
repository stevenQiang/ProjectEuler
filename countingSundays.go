//
// 最简单的方法： 1200/7 = 171
package main

import (
	"time"
	"fmt"
)

func main() {
	startDate,_ := time.Parse("2006-1-2", "1901-1-1")
	endDate,_ := time.Parse("2006-1-2", "2000-12-31")
	d := startDate
	count := 0
	for {
		if int(d.Weekday()) == 0 {
			count++
		}
		if int(d.Month()) == int(endDate.Month()) && int(d.Year()) == int(endDate.Year()) {
			break
		}
		d=d.AddDate(0, 1, 0)
	}
	fmt.Println(count)
}
