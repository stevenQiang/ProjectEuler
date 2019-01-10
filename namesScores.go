package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"sort"
	"strings"
	"fmt"
)

func main() {
	f, err := os.OpenFile("22/p022_names.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()
	var result []string
	var sum int
	rd := bufio.NewReader(f)
	for {
		a,err := rd.ReadString(',')
		if err != nil {
			if err == io.EOF {
				break
			}
			loxg.Fatalf("read file line error: %v", err)
			return
		}
		result = append(result, a)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	for k, v := range result {
		k = k + 1
		value := strings.Replace(strings.Replace(v, "\"", "", -1), ",", "", -1)
		word := wordValue(value)
		fmt.Println(k, value, word)
		sum += word * k
	}
	fmt.Println(sum)
}

func wordValue(word string) int{
	num := 0
	for i := 0; i< len(word); i++ {
		num += int(rune(word[i] - 'A' + 1))
	}
	return num
}
