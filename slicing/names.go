package slicing

import "fmt"

func Slot() {
	Names := []string{"okwiir", "okwii", "okiit"}

	for _, name := range Names {
		fmt.Println(name, "is a member")
	}
}
