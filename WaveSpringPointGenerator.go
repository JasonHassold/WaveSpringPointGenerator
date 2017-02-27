/*
   -------------------------------------
   | Wave Wire Spring Points Generator |
   -------------------------------------
   Written by Jason Hassold

   This program creates an .obj file containing the vertexes for a spring that
   is in the shape of a spiral sine wave. The user enters in a series of
   conditions to change things like the size, width, how many degrees between
   each point.
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

// Variables inputted by the user
var InnerDiam, OuterDiam float64
var Height, Thickness float64
var Gap float64 // Gap between overlapping peak and trough
var Degrees, Revolutions, Periods float64

var Circumference float64
var Slope float64
var Period float64
var RevHeight float64
var MaxAmplitude float64
var Center [2]float64
var Amplitude float64

var Spring []Vertex
var NumOfPoints int

// Parameter is the reference of a variable
func input(i *float64) {
	_, err := fmt.Scanf("%f\r\n", i)

	if err != nil {
		fmt.Println(err)
	}
}

/******************
   Main Function
******************/
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
	fmt.Print("Thickness: ")
	input(&Thickness)
	fmt.Print("Gap: ")
	input(&Gap)
	fmt.Print("Degrees between points: ")
	input(&Degrees)
	fmt.Print("Number of revolutions (needs to be odd): ")
	input(&Revolutions)
	fmt.Print("Number of periods (*.5): ")
	input(&Periods)
	fmt.Print("")

	// C = Pi * d
	Circumference = math.Pi * ((OuterDiam + InnerDiam) / 2.0)
	// m = h / (C * rev)
	Slope = Height / (Circumference * Revolutions)
	// T = (2*Pi) / (C / #ofT)
	Period = (2 * math.Pi) / (Circumference / Periods)
	// h per rev = h / rev
	RevHeight = Height / Revolutions
	// MaxA = h per rev / 2
	MaxAmplitude = RevHeight / 2.0
	//
	Center[0] = ((Circumference / Periods) * (.75)) * Slope
	//
	Center[1] = (Circumference + (Circumference/Periods)*(.75)) * Slope
	//
	Amplitude = (Center[1] - Center[0] - Thickness - Gap) / 2.0

	refineAmp()

	NumOfPoints = (360 / int(Degrees)) * int(Revolutions) * 2
	Spring = make([]Vertex, NumOfPoints)
	generate()

	outputSpring()
}

func SpringSine(x float64) float64 {
	return Amplitude*math.Sin(Period*x-(3.0*math.Pi)/2.0) + Slope*x
}

// This function increases the accuracy of the amplitude
func refineAmp() {
	//fmt.Println("Amp: " + Amplitude)

	x1 := (math.Acos(-Slope/(Amplitude*Period)) + (3.0*math.Pi)/2.0) / Period
	x2 := (math.Acos(Slope/(Amplitude*Period)) + (3.0*math.Pi)/2.0 + Circumference) / Period // Removing the - is because it is a min
	z1 := SpringSine(x1)
	z2 := SpringSine(x2)

	deltaY := math.Abs(z2 - z1)
	offset := (Thickness + Gap) - deltaY
	Amplitude -= offset / 2.0

	//fmt.Println("x: " + x1, x2)
	//fmt.Println("y: " + y1, y2)
	//fmt.Println("deltaY: " + deltaY)
	//fmt.Println("offset: " + offset)
	//fmt.Println("Amp: " + Amplitude)
	//fmt.Println()
}

func generate() {
	radiansXY := (Degrees * math.Pi) / 180.0      // Standard 2pi circle
	radiansZ := (Degrees / 360.0) * Circumference // Total sine wave length

	for p := 0; p < NumOfPoints; p += 2 {
		// Inner Diameter
		Spring[p].x = (InnerDiam / 2.0) * math.Cos(radiansXY*float64(p/2))
		Spring[p].y = (InnerDiam / 2.0) * math.Sin(radiansXY*float64(p/2))
		Spring[p].z = SpringSine(radiansZ * float64(p/2))
		// Outer Diameter
		Spring[p+1].x = (OuterDiam / 2.0) * math.Cos(radiansXY*float64(p/2))
		Spring[p+1].y = (OuterDiam / 2.0) * math.Sin(radiansXY*float64(p/2))
		Spring[p+1].z = SpringSine(radiansZ * float64(p/2))
	}

	//fmt.Println(Spring)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func outputSpring() {
	file1, err := os.Create("springInner.txt")
	check(err)
	defer file1.Close()

	file2, err := os.Create("springOuter.txt")
	check(err)
	defer file2.Close()

	//_, err = file.WriteString("# List of vertexes in (x,y,z) form\n")
	//check(err)

	for l := 0; l < len(Spring); l++ {
		x := strconv.FormatFloat(Spring[l].x, 'f', 6, 64)
		y := strconv.FormatFloat(Spring[l].y, 'f', 6, 64)
		z := strconv.FormatFloat(Spring[l].z, 'f', 6, 64)

		if l%2 == 0 {
			_, err = file1.WriteString(x + " " + y + " " + z + "\n")
		} else {
			_, err = file2.WriteString(x + " " + y + " " + z + "\n")
		}

		check(err)
	}

	file1.Sync()
	file2.Sync()
}
