package main

func main() {
	europeanPlug := &EuropeanPlug{}
	charger := &EuropeanPlugAdapter{
		europeanPlug: europeanPlug,
	}
	charger.ChargePhone()
}
