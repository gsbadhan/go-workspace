package main

import (
	"fmt"
	"structsops/business"
)

func main() {
	fmt.Println("main..")
	address := business.Address{Id: 2001, Country: "india"}
	emp := business.Employee{Id: 1001, Name: "gs", Age: 45.2, Address: address}
	emp.Save(emp)
	fmt.Printf("emp.Get(1001): %v\n", emp.Get(1001))
	business.Total()
}
