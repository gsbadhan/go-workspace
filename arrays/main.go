package main

import (
	"fmt"
)

func main() {
	fmt.Println("main ..")

	// static array
	var stringsDim1 [2]string
	stringsDim1[0] = "hello"
	stringsDim1[1] = "world"
	fmt.Println("stringsDim1", stringsDim1)

	stringsSliceDim1 := []string{"hello", "world", "!!"}
	stringsSliceDim1 = append(stringsSliceDim1, "do it")
	stringsSliceDim1 = append(stringsSliceDim1, "will come back", "wait for me")
	for indx, val := range stringsSliceDim1 {
		fmt.Println(indx, val)
	}

	// 1D int array
	var intDim1 [2]int
	intDim1[0] = 1001
	intDim1[1] = 1002
	fmt.Println(intDim1)

	intDim1Slice := []int{1, 2, 3}
	intDim1Slice = append(intDim1Slice, 4)
	fmt.Println(intDim1Slice)

	var intDim2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(intDim2)

	intDim2Slice := [][]int{}
	intDim2Slice = append(intDim2Slice, []int{1, 2, 3})
	intDim2Slice = append(intDim2Slice, []int{3, 4, 5})
	fmt.Println(intDim2Slice)

}
