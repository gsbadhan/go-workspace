package business

import "fmt"

type Employee struct {
	Id      int
	Name    string
	Age     float32
	Address Address
}

type Address struct {
	Id      int
	Country string
}

var employeeDB = make(map[int]Employee, 10)

func (e *Employee) Save(emp Employee) bool {
	fmt.Println("save emp", emp)
	employeeDB[emp.Id] = emp
	return true
}

func (e *Employee) Get(empId int) Employee {
	return employeeDB[empId]
}

func Total() int {
	fmt.Printf("employeeDB: %v\n", len(employeeDB))
	return len(employeeDB)
}
