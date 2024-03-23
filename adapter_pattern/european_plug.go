package main

import "fmt"

type EuropeanPlug struct{}

func (e *EuropeanPlug) PlugIn() {
	fmt.Println("European plug connected")
}
