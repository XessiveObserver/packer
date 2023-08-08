package slicecap

import "fmt"

func Cities() {
	CityNames := make([]string, 3, 6)
	CityNames[0] = "Mbale"
	CityNames[1] = "Soroti"
	CityNames[2] = "Jinja"

	for _, city := range CityNames {
		fmt.Printf("%v is a city now\n", city)
	}
	fmt.Println("\nYou have", len(CityNames), "cities")
}
