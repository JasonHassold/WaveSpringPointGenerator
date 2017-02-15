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
)

type Vertex struct {
	x, y, z float64
}

var InnerDiam, OuterDiam, Height float64
var Degrees, Revolutions float64
var Thinkness, Gap float64
var Origin Vertex = Vertex{0, 0, 0}

func main() {
	fmt.Println("-------------------------------------")
	fmt.Println("| Wave Wire Spring Points Generator |")
	fmt.Println("-------------------------------------")
	fmt.Print("")
	fmt.Print("Inner Diameter: ")
	_, err := fmt.Scanf("%f", &InnerDiam)
	fmt.Print("Outer Diameter: ")
	_, err := fmt.Scanf("%f", &OuterDiam)
	fmt.Print("Height: ")
	_, err := fmt.Scanf("%f", &Height)
	fmt.Print("Degrees between points: ")
	_, err := fmt.Scanf("%f", &Degrees)
	fmt.Print("Number of revolutions: ")
	_, err := fmt.Scanf("%f", &Revolutions)
	fmt.Print("Thinkness: ")
	_, err := fmt.Scanf("%f", &Thinkness)
	fmt.Print("Gap: ")
	_, err := fmt.Scanf("%f", &Gap)
}
