package main

import "github.com/RicardoKim/Study/DesignPattern/CommandPattern/command"

type Invoker struct {
	startCommand command.Command
	stopCommand  command.Command
}

func (i *Invoker) SetStartCommand(command command.Command) {
	i.startCommand = command
}

func (i *Invoker) SetStopCommand(command command.Command) {
	i.stopCommand = command
}

func (i *Invoker) StartProcess() {
	i.startCommand.Execute()
}

func (i *Invoker) StopProcess() {
	i.stopCommand.Execute()
}
