package model

type Job struct {
	Name     string    `json:"name"`
	Machines []Machine `json:"machines"`
}
