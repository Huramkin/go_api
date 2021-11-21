package domain

import "fmt"

type ServerInstance struct {
	Id     int
	Name   string
	Ram    float32
	isDedi string
}

func (si *ServerInstance) print(s string) {
	fmt.Println(si.Id, si.Name, si.Ram, si.isDedi)
}
