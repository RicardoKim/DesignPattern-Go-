package command

import "github.com/RicardoKim/Study/DesignPattern/CommandPattern/process"

type StartIOProcessCommand struct {
	ioProcess *process.IOProcess
}

func (ic *StartIOProcessCommand) Execute() {
	ic.ioProcess.StartProcess()
}

func NewStartIOProcessCommand(ioProcess *process.IOProcess) *StartIOProcessCommand {
	return &StartIOProcessCommand{ioProcess: ioProcess}
}
