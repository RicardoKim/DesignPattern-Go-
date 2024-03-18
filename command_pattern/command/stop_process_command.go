package command

import "github.com/RicardoKim/Study/DesignPattern/CommandPattern/process"

type StopProcessCommand struct {
	process *process.Process
}

func (c *StopProcessCommand) Execute() {
	c.process.Stop()
}

func NewStopProcessCommand(process *process.Process) *StopProcessCommand {
	return &StopProcessCommand{process: process}
}
