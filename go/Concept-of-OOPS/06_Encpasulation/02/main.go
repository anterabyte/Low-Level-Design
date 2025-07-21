package main

import "fmt"

type Employee struct {
	name string
	age int 
}

// Getter
func (e *Employee) GetEmployeeName() string {
	return fmt.Sprintln("Employee Name -",e.name)
}

// Setter
func (e *Employee) SetEmployeeName(name string) {
	e.name = name
}

// Getter for the age
func (e *Employee) GetEmployeeAge() string {
	return fmt.Sprintln("Employee Age -",e.age)
}

// Setter for the age
func (e *Employee) SetEmployeeAge(age int) {
	if age > 18 {
		e.age = 18
	}else{
		fmt.Println("Go first become 18 of age")
	}
}

func main(){

	emp1 := Employee{}

	emp1.SetEmployeeName("Hulu")

	emp1.SetEmployeeAge(21)

	fmt.Println (emp1.GetEmployeeName())

	fmt.Println(emp1.GetEmployeeAge())
	
}
