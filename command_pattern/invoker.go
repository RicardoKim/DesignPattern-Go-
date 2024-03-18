package main

import "github.com/RicardoKim/Study/DesignPattern/CommandPattern/command"

type Invoker struct {
	startCommands []command.Command
	stopCommands  []command.Command
}

func (i *Invoker) AddStartCommand(command command.Command) {
	i.startCommands = append(i.startCommands, command)
}

func (i *Invoker) AddStopCommand(command command.Command) {
	i.stopCommands = append(i.stopCommands, command)
}

func (i *Invoker) StartProcess() {
	for _, cmd := range i.startCommands {
		cmd.Execute()
	}
}

func (i *Invoker) StopProcess() {
	for _, cmd := range i.stopCommands {
		cmd.Execute()
	}
}
