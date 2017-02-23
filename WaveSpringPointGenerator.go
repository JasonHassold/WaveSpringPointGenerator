/*
   -------------------------------------
   | Wave Wire Spring Points Generator |
   -------------------------------------
   Written by Jason Hassold

   This program creates an ASCII point file for a spring that is in the shape
   of a spiral sine wave. The user enters in a series of conditions to change
   things like the size, width, how many degrees between each point.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y, z float64
}

// User defined vars
var InnerDiam, OuterDiam, Height float64
var Degrees, Revolutions, Periods float64
var Thickness, Gap float64
var Origin Vertex = Vertex{0, 0, 0}

var Circumference float64
var Slope float64
var RevHeight float64
var Amplitude float64
var MaxAmplitude float64
var Period float64
var Center [2]float64

func input(i *float64) {
	_, err := fmt.Scanf("%f", i)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("-------------------------------------")
	fmt.Println("| Wave Wire Spring Points Generator |")
	fmt.Println("-------------------------------------")
	fmt.Print("")
	fmt.Print("Inner Diameter: ")
	input(&InnerDiam)
	fmt.Print("Outer Diameter: ")
	input(&OuterDiam)
	fmt.Print("Height: ")
	input(&Height)
	fmt.Print("Degrees between points: ")
	input(&Degrees)
	fmt.Print("Number of revolutions: ")
	input(&Revolutions)
	fmt.Print("Number of periods: ")
	input(&Periods)
	fmt.Print("Thickness: ")
	input(&Thickness)
	fmt.Print("Gap: ")
	input(&Gap)

	Circumference = math.Pi * ((OuterDiam + InnerDiam) / 2.0)
	Slope = Height / (Circumference * Revolutions)
	RevHeight = Height / Revolutions
	MaxAmplitude = RevHeight / 2.0
	Period = (2 * math.Pi) / (Circumference / Periods)
	Center[0] = ((Circumference / Periods) * (.75)) * Slope
	Center[1] = (Circumference + (Circumference/Periods)*(.75)) * Slope
	Amplitude = (Center[1] - Center[0] - Thickness - Gap) / 2.0

	refineAmp()
}

func Sine(x float64) float64 {
	return Amplitude*math.Sin(Period*x-math.Pi) + Slope*x
}

// This function increases the accuracy of the amplitude by stepping.
func refineAmp() {
	count := 0

	for count < 5 {
		fmt.Println(Amplitude)
		x1 := Period*math.Acos(-Slope/Amplitude) + math.Pi
		x2 := x1 + Circumference
		y1 := Sine(x1)
		y2 := Sine(x2)

		fmt.Println(x1, x2)
		fmt.Println(y1, y2)

		deltaY := math.Abs(y2 - y1)
		offset := (Thickness + Gap) - deltaY
		Amplitude -= offset / 2.0

		fmt.Println(deltaY)
		fmt.Println(offset)
		fmt.Println(Amplitude)
		fmt.Println()

		count += 1
	}
}

func generate() {

}
