package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("main..")

	// init
	intDim1 := []int{}

	//append
	intDim1 = append(intDim1, 1, 2, 3, 4, 5, 6, 7, 100)
	fmt.Println(intDim1, len(intDim1), intDim1[len(intDim1)-1])

	// copy
	intDim1Copy := intDim1[1:6]
	fmt.Println(intDim1Copy, len(intDim1))
	intDim1Copy2 := intDim1[0 : len(intDim1)/2]
	fmt.Println(intDim1Copy2, len(intDim1Copy2))

	// replace at index
	intDim1[len(intDim1)-1] = 200
	fmt.Println(intDim1)

	// search
	for idx, val := range intDim1 {
		if val == 3 {
			fmt.Println("found", idx, val)
			break
		}
	}

	//sort
	sort.Ints(intDim1)
	fmt.Println("sort asc", intDim1)
	sort.Slice(intDim1, func(i, j int) bool {
		return intDim1[i] > intDim1[j]
	})
	fmt.Println("sort desc", intDim1)

}
