package process

import "fmt"

type IOProcess struct {
	id         string
	doesItWait bool
}

func (ip *IOProcess) StartProcess() {
	fmt.Printf("Process %s is started.\n", ip.id)
	ip.doesItWait = true
	fmt.Printf("Request IO\n")
}

func (ip *IOProcess) TerminateIOProcess() {
	fmt.Printf("Terminate IO Process\n")
	ip.doesItWait = false
}

func (ip *IOProcess) StopProcess() {
	if ip.doesItWait {
		fmt.Printf("Fail to Terminate\n")
	} else {
		fmt.Printf("Terminate Process\n")
	}
}

func NewIOProcess(id string) *IOProcess {
	return &IOProcess{id: id}
}
