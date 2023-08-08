package main

import (
	"fmt"
	"packer/addition"
	"packer/embed"
	"packer/personality"
	"packer/routines"
	"packer/slicecap"
	"packer/slicing"
)

func main() {
	fmt.Println("You must become good")

	fmt.Println("\nSum: ", addition.Add(3, 5))

	fmt.Println()
	slicing.Slot()

	fmt.Println()
	slicecap.Cities()

	fmt.Println()
	personality.Details()

	fmt.Println()
	routines.Hunt()

	fmt.Println()
	embed.Details()
}
