package main

import (
	"github.com/RicardoKim/Study/DesignPattern/CommandPattern/command"
	"github.com/RicardoKim/Study/DesignPattern/CommandPattern/process"
)

func main() {
	normalProcess := process.NewProcess("1234")
	startCommand := command.NewStartProcessCommand(normalProcess)
	stopCommand := command.NewStopProcessCommand(normalProcess)

	invoker := &Invoker{}
	invoker.SetStartCommand(startCommand)
	invoker.SetStopCommand(stopCommand)

	invoker.StartProcess()

	invoker.StopProcess()

	ioProcess := process.NewIOProcess("5678")
	startIOCommand := command.NewStartIOProcessCommand(ioProcess)
	stopIOCommand := command.NewStopIOProcessCommand(ioProcess)

	ioInvoker := &Invoker{}
	ioInvoker.SetStartCommand(startIOCommand)
	ioInvoker.SetStopCommand(stopIOCommand)

	ioInvoker.StartProcess()

	ioInvoker.StopProcess()
}
