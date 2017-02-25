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
	"os"
	"strconv"
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

var Spring []Vertex

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
	generate()
	outputSpring()
}

func Sine(x float64) float64 {
	return Amplitude*math.Sin(Period*x-(3.0*math.Pi)/2.0) + Slope*x
}

// This function increases the accuracy of the amplitude by stepping.
func refineAmp() {
	count := 0

	for count < 5 {
		//fmt.Println(Amplitude)

		x1 := (math.Acos(-Slope/(Amplitude*Period)) + (3.0*math.Pi)/2.0) / Period
		x2 := (math.Acos(-Slope/(Amplitude*Period)) + (3.0*math.Pi)/2.0 + Circumference) / Period
		y1 := Sine(x1)
		y2 := Sine(x2)

		//fmt.Println(x1, x2)
		//fmt.Println(y1, y2)

		deltaY := math.Abs(y2 - y1)
		offset := (Thickness + Gap) - deltaY
		Amplitude -= offset / 2.0

		//fmt.Println(deltaY)
		//fmt.Println(offset)
		//fmt.Println(Amplitude)
		//fmt.Println()

		count += 1
	}
}

func generate() {
	numOfPoints := (360 / Degrees) * Revolutions
	radiansXY := (Degrees * math.Pi) / 180.0
	radiansZ := (Degrees / 360.0) * Circumference
	Spring = make([]Vertex, 2*int(numOfPoints))

	for p := 0; p < 2*int(numOfPoints); p += 2 {
		Spring[p].x = (InnerDiam / 2.0) * math.Cos(radiansXY*float64(p/2))
		Spring[p].y = (InnerDiam / 2.0) * math.Sin(radiansXY*float64(p/2))
		Spring[p].z = Sine(radiansZ * float64(p/2))

		Spring[p+1].x = (OuterDiam / 2.0) * math.Cos(radiansXY*float64(p/2))
		Spring[p+1].y = (OuterDiam / 2.0) * math.Sin(radiansXY*float64(p/2))
		Spring[p+1].z = Sine(radiansZ * float64(p/2))
	}

	//fmt.Println(Spring)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func outputSpring() {
	file, err := os.Create("spring.obj")
	check(err)
	defer file.Close()

	_, err = file.WriteString("# List of vertexes in (x,y,z) form\n")
	check(err)

	for l := 0; l < len(Spring); l++ {
		x := strconv.FormatFloat(Spring[l].x, 'f', 6, 64)
		y := strconv.FormatFloat(Spring[l].y, 'f', 6, 64)
		z := strconv.FormatFloat(Spring[l].z, 'f', 6, 64)
		_, err = file.WriteString("v " + x + " " + y + " " + z + "\n")
		check(err)
	}

	file.Sync()
}
