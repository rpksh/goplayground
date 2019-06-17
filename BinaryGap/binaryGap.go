package main

import (
	"fmt"
	"strconv"
)

//Solution to get binary Gap
func Solution(N int) int {
	binaryNumber := strconv.FormatInt(int64(N), 2)
	fmt.Println(" binaryNumber ", binaryNumber)
	var leng []int
	length := len(binaryNumber)
	//fmt.Println(length)
	head := 0
	tail := 0
	diff := 0
	for tail < length {
		fmt.Println(" = loop start = ")
		fmt.Println("char ", string(binaryNumber[tail]))
		fmt.Println("head ", head, "tail ", tail, "diff ", diff)
		if string(binaryNumber[tail]) == "1" {
			diff = tail - head - 1
			leng = append(leng, diff)
			head = tail
		}
		fmt.Println("head ", head, "tail ", tail, "diff ", diff)
		fmt.Println("leng ", leng)
		//fmt.Println("digit : ", string(binaryNumber[tail]), "head : ", tail)
		tail = tail + 1
		fmt.Println(" = loop end = ")
	}
	fmt.Println(leng)
	/*
		for position, digit := range binaryNumber {
			fmt.Println("digit : ", string(digit), "position : ", position)
		}
	*/
	return 0
}

func main() {
	Solution(1041)
}
