package model

import "fmt"

type Job struct {
	Name     string    `json:"name"`
	Machines []Machine `json:"machines"`
}

func (i *Job) Println() {
	fmt.Println("---")
	fmt.Println("Job Name: ", i.Name)
	for _, m := range i.Machines {
		fmt.Println("   " + m.ToString())
	}
}
