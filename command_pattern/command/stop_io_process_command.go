package command

import "github.com/RicardoKim/Study/DesignPattern/CommandPattern/process"

type StopIOProcessCommand struct {
	ioProcess *process.IOProcess
}

func (ic *StopIOProcessCommand) Execute() {
	ic.ioProcess.TerminateIOProcess()
	ic.ioProcess.StopProcess()
}

func NewStopIOProcessCommand(ioProcess *process.IOProcess) *StopIOProcessCommand {
	return &StopIOProcessCommand{ioProcess: ioProcess}
}
