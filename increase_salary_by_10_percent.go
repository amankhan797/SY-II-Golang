// WAP in Go language to take details of employees (EmpId, EmpName, Contact, Address, Salary, and Designation), take Designation from user for which you have to update the salary of the employee and decrease the salary by 10%.
package main

import "fmt"

func main() {
	var EmpId int
	var EmpName string
	var Contact int
	var Address string
	var Salary int
	var Designation string
	fmt.Print("Enter Employee ID : ")
	fmt.Scan(&EmpId)
	fmt.Print("Enter Employee Name : ")
	fmt.Scan(&EmpName)
	fmt.Print("Enter Employee Contact : ")
	fmt.Scan(&Contact)
	fmt.Print("Enter Employee Address : ")
	fmt.Scan(&Address)
	fmt.Print("Enter Employee Salary : ")
	fmt.Scan(&Salary)
	fmt.Println()
	fmt.Print("Enter Employee Designation : ")
	fmt.Scan(&Designation)
	if Designation == "manager" {
		Salary = Salary - (Salary/100)*10
		print("Decreased Salary is :", Salary)
	} else {
		fmt.Print("cannot decreased salary on this Designation")
	}
}
