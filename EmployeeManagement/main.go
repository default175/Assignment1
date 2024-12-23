package main

import "fmt"

func main() {

	company := Company{Employees: make(map[string]Employee)}


	fte1 := FullTimeEmployee{ID: 1, Name: "Alice", Salary: 5000}
	fte2 := FullTimeEmployee{ID: 2, Name: "Bob", Salary: 4000}

	pte1 := PartTimeEmployee{ID: 3, Name: "Charlie", HourlyRate: 50, HoursWorked: 20.5}
	pte2 := PartTimeEmployee{ID: 4, Name: "Diana", HourlyRate: 60, HoursWorked: 15.0}


	company.AddEmployee(fmt.Sprint(fte1.ID), fte1)
	company.AddEmployee(fmt.Sprint(fte2.ID), fte2)
	company.AddEmployee(fmt.Sprint(pte1.ID), pte1)
	company.AddEmployee(fmt.Sprint(pte2.ID), pte2)


	fmt.Println("\nEmployee List:")
	company.ListEmployees()
}
