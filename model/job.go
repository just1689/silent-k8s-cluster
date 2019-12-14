package model

type Job struct {
	Name     string    `json:"name"`
	Machines []Machine `json:"machines"`
}

func (i *Job) ToString() string {
	result := "Name: " + i.Name + " Machines: "
	for _, m := range i.Machines {
		result += m.ToString() + ", "
	}
	return result
}
