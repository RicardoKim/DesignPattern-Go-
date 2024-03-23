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

	ioProcess2 := process.NewIOProcess("91011")
	startIOCommand2 := command.NewStartIOProcessCommand(ioProcess2)
	stopIOCommand2 := command.NewStopIOProcessCommand(ioProcess2)

	invoker := &Invoker{}
	invoker.AddStartCommand(startCommand)
	invoker.AddStopCommand(stopCommand)

	invoker.AddStartCommand(startIOCommand)
	invoker.AddStopCommand(stopIOCommand)

	invoker.AddStartCommand(startIOCommand2)
	invoker.AddStopCommand(startIOCommand2)

	invoker.StartProcess()
	invoker.StopProcess()
}
