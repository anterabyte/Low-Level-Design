package main

import "fmt"

type Professor struct{
	name, department string
}

func (p *Professor) Teach() {
	fmt.Println(p.name+" Teaches in "+p.department+" department")
}

type College struct{
	name string
	professor   []*Professor
}

func (c *College) AddProfessor(professor *Professor) {
	c.professor = append(c.professor, professor)
}

func (c *College) ShowAllProfessors() {
	fmt.Println("Professors at",c.name)
	for _, v := range c.professor {
		v.Teach()
	}
}

func main() {

	p1 := &Professor{
		"Einstein",
		"Relativity",
	}

	p2 := &Professor{
		"Edison",
		"Electrical",
	}

	
	var college  = College{ name: "Berkley"}

	college.AddProfessor(p1)

	college.AddProfessor(p2)

	college.ShowAllProfessors()

	p1.Teach()

	p2.Teach()
}
