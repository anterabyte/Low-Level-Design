package main

import "fmt"

type Student struct{
	name string
	id int
}

func (s *Student) Details() {
	fmt.Printf("Student Name is %s and Its ID is %d",s.name,s.id)
}

type Teacher struct {
	name string
	students []*Student
}

func (t *Teacher) AddStudent(student *Student) {
	t.students = append(t.students, student)
}

func (t *Teacher) ShowStudents() {
	fmt.Println("-----------------------------------")
	fmt.Println("All the students under",t.name)
	fmt.Println("-----------------------------------")
	for _,v := range t.students {
		fmt.Println(v.name)
	}
}


func main() {
	var stu1 = &Student{
		name: "Akash",
		id: 112233,
	}

	var stu2 = &Student{
		name: "Agni",
		id: 223344,
	}

	var teacher1 = Teacher{name: "Canakya"}

	var teacher2 = Teacher{name: "Adi"}

	teacher1.AddStudent(stu1)

	teacher1.AddStudent(stu2)
	
	teacher2.AddStudent(stu2)

	teacher1.ShowStudents()

	teacher2.ShowStudents()
}
