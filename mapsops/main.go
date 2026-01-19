package main

import "fmt"

func main() {
	fmt.Println("main..")

	// init
	mapStringInt := make(map[string]int, 0)

	// put
	mapStringInt["1001"] = 1001
	mapStringInt["1002"] = 1002
	fmt.Println("all", mapStringInt)

	//get
	fmt.Println("get", mapStringInt["1001"])
	fmt.Println("get", mapStringInt["1110"])

	//delete
	delete(mapStringInt, "1001")
	fmt.Println("after delete", mapStringInt)
	delete(mapStringInt, "1110")
	fmt.Println("after delete", mapStringInt)

	// show keys
	keys := make([]string, 0, len(mapStringInt))
	for k := range mapStringInt {
		keys = append(keys, k)
	}
	fmt.Println("keys", keys)

	// delete all
	for k := range mapStringInt {
		delete(mapStringInt, k)
	}

}
