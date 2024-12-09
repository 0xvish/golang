package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	fuel uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.fuel
}

type electricEngine struct {
	mpkwh uint8
	charge   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.charge
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if e.milesLeft() >= miles {
		fmt.Println("You can make it to your destination: ", miles)
	} else {
		fmt.Println("You cannot make it to your destination: ", miles)
	}
}

func main() {

	var myEngine gasEngine // struct with fields mpg and gallons
	myEngine.mpg = 7
	myEngine.fuel = 10
	fmt.Println("Miles left: ", myEngine.milesLeft())
	canMakeIt(myEngine, 100)

	myElectricEngine := electricEngine{mpkwh: 30, charge: 6} // struct with fields mpkwh and kwh
	fmt.Println("Miles left: ", myElectricEngine.milesLeft())
	canMakeIt(myElectricEngine, 100)

}