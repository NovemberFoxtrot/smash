package main

import (
	"fmt"
	"math"
)

type Gear struct {
	chainring float64
	cog       float64
	wheel     *wheel
}

type wheel struct {
	rim  float64
	tire float64
}

func NewGear(chainring, cog, rim, tire float64) Gear {
	return Gear{chainring: chainring, cog: cog, wheel: &wheel{rim, tire}}
}

func (gear Gear) Chainring() float64 {
	return gear.chainring
}

func (gear Gear) Cog() float64 {
	return gear.cog
}

func (wheel wheel) Rim() float64 {
	return wheel.rim
}

func (wheel wheel) Tire() float64 {
	return wheel.tire
}

func (wheel wheel) Diameter() float64 {
	return wheel.Rim() + (wheel.Tire() * 2)
}

func (wheel wheel) Circumference() float64 {
	return wheel.Diameter() * math.Pi
}

func (gear Gear) Ratio() float64 {
	return gear.Chainring() / gear.Cog()
}

func (gear Gear) Wheel() wheel {
	return *gear.wheel
}

func (gear Gear) GearInches() float64 {
	return gear.Ratio() * gear.Wheel().Diameter()
}

func main() {
	fmt.Println(NewGear(52, 11, 26, 1.5).GearInches())
	fmt.Println(NewGear(52, 11, 24, 1.25).GearInches())
	fmt.Println(NewGear(52, 11, 24, 1.25).Wheel().Circumference())
}
