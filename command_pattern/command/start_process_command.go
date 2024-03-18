package command

import "github.com/RicardoKim/Study/DesignPattern/CommandPattern/process"

type StartProcessCommand struct {
	process *process.Process
}

func (c *StartProcessCommand) Execute() {
	c.process.Start()
}

func NewStartProcessCommand(process *process.Process) *StartProcessCommand {
	return &StartProcessCommand{process: process}
}
