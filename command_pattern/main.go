package main

import (
	"github.com/RicardoKim/Study/DesignPattern/CommandPattern/command"
	"github.com/RicardoKim/Study/DesignPattern/CommandPattern/process"
)

func main() {
	normalProcess := process.NewProcess("1234")
	startCommand := command.NewStartProcessCommand(normalProcess)
	stopCommand := command.NewStopProcessCommand(normalProcess)

	ioProcess := process.NewIOProcess("5678")
	startIOCommand := command.NewStartIOProcessCommand(ioProcess)
	stopIOCommand := command.NewStopIOProcessCommand(ioProcess)

	invoker := &Invoker{}
	invoker.AddStartCommand(startCommand)
	invoker.AddStopCommand(stopCommand)

	invoker.AddStartCommand(startIOCommand)
	invoker.AddStopCommand(stopIOCommand)

	invoker.StartProcess()
	invoker.StopProcess()
}
