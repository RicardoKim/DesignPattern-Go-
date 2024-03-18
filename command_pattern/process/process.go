package process

import "fmt"

type Process struct {
	id string
}

func (p *Process) Start() {
	fmt.Printf("Process %s is started.\n", p.id)
}

func (p *Process) Stop() {
	fmt.Printf("Process %s is stopped.\n", p.id)
}

func NewProcess(id string) *Process {
	return &Process{id: id}
}
