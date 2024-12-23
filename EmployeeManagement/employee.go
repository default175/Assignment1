package main

import (
	"fmt"
	"strconv"
)


type Employee interface {
	GetDetails() string
}


type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}


type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}


func (fte FullTimeEmployee) GetDetails() string {
	return "Full-Time Employee [ID: " + strconv.FormatUint(fte.ID, 10) +
		", Name: " + fte.Name +
		", Salary: " + strconv.Itoa(int(fte.Salary)) + " Tenge]"
}


func (pte PartTimeEmployee) GetDetails() string {
	return "Part-Time Employee [ID: " + strconv.FormatUint(pte.ID, 10) +
		", Name: " + pte.Name +
		", Hourly Rate: " + strconv.FormatUint(pte.HourlyRate, 10) +
		" Tenge, Hours Worked: " + fmt.Sprintf("%.1f", pte.HoursWorked) + "]"
}


type Company struct {
	Employees map[string]Employee
}


func (c *Company) AddEmployee(id string, emp Employee) {
	if c.Employees == nil {
		c.Employees = make(map[string]Employee)
	}
	c.Employees[id] = emp
	fmt.Println("Employee added successfully.")
}


func (c *Company) ListEmployees() {
	if len(c.Employees) == 0 {
		fmt.Println("No employees in the company.")
		return
	}
	for id, emp := range c.Employees {
		fmt.Println("ID:", id, "|", emp.GetDetails())
	}
}
