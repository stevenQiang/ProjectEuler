// 从底部开始算，可以减少时间复杂度

package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	f, err := os.OpenFile("18/data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	var result [][]string
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
			return
		}
		line = strings.TrimSpace(line)
		lineList := strings.Split(line, " ")
		result = append(result, lineList)
	}
	fmt.Println(analysisData(result))
}

func analysisData(data [][]string) string {
	for index := len(data) -2; index >= 0; index-- {
		var newList []string
		for i := 0; i < len(data[index]); i++ {
			oldData, _ := strconv.Atoi(data[index][i])
			a, _ := strconv.Atoi(data[index+1][i])
			a1, _ := strconv.Atoi(data[index+1][i+1])
			maxInt := getMaxInt(oldData+a, oldData+a1)
			newList = append(newList, strconv.Itoa(maxInt))
		}
		data[index] = newList
	}
	return data[0][0]
}

func getMaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//func analysisData(index int, data [][]string) int{
//	if index == len(data) - 2 {
//		var newList []int
//		for i := 0; i < len(data[index]); i++ {
//			var maxValue = 0
//			a, _ := strconv.Atoi(data[index+1][i])
//			a1, _ := strconv.Atoi(data[index+1][i+1])
//			if a+a1 > maxValue {
//				maxValue = a + a1
//			}
//			newList = append(newList, maxValue)
//		}
//		return 1
//	} else{
//		return analysisData(index+1, data)
//	}
//}
