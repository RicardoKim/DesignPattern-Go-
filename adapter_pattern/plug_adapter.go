package main

import "fmt"

type EuropeanPlugAdapter struct {
	europeanPlug *EuropeanPlug
}

func (adapter *EuropeanPlugAdapter) ChargePhone() {
	adapter.europeanPlug.PlugIn()
	fmt.Println("Phone is charging using European plug")
}
