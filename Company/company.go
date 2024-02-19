package main

import "fmt"

type programming struct {
	title     string
	salary    float64
	dephead   string
	employees []string
}

type design struct {
	title     string
	salary    float64
	dephead   string
	employees []string
}

type Department interface {
	EmployeesList() []string
	Title() string
	TotalEmployessCount() int
	HeadOfDepartment() string
	NetSalary() float64
}

func (d design) EmployeesList() []string {
	return d.employees
}
func (p programming) EmployeesList() []string {
	return p.employees
}

func (d design) Title() string {
	return d.title
}
func (p programming) Title() string {
	return p.title
}

func (d design) TotalEmployessCount() int {
	return len(d.employees)
}
func (p programming) TotalEmployessCount() int {
	return len(p.employees)
}

func (d design) HeadOfDepartment() string {
	return d.dephead
}
func (p programming) HeadOfDepartment() string {
	return p.dephead
}

func (d design) NetSalary() float64 {
	return d.salary * 0.75
}
func (p programming) NetSalary() float64 {
	return p.salary * 0.5
}

func ShowDepartmentDetails(de Department) {
	fmt.Println(de.EmployeesList())
	fmt.Println(de.Title())
	fmt.Println(de.TotalEmployessCount())
	fmt.Println(de.HeadOfDepartment())
	fmt.Println(de.Title())
	fmt.Println(de.NetSalary())
}

func main() {
	d := design{"SomeOne", 123.45, "abc", []string{"employee", "employ"}}
	ShowDepartmentDetails(d)
	p := programming{"Person", 123.45, "abc", []string{"employee1", "employsee2"}}
	ShowDepartmentDetails(p)
}
