package embed

import "fmt"

type Device struct {
	DeviceName string
	MadeIn     string
	MadeBy     string
}

type Dist struct {
	Device
	DistName string
	Located  string
	Plot     int
}

func (d Dist) Display() {
	fmt.Printf("%v is made in %v by %v ... \ndistributed by %v located in %v on plot %v", d.DeviceName, d.MadeIn, d.MadeBy, d.DistName, d.Located, d.Plot)
}

func Details() {
	dist := Dist{
		Device: Device{
			DeviceName: "BlackBerry",
			MadeIn:     "Waterloo",
			MadeBy:     "BlackBerry Inc",
		},
		DistName: "AT&T",
		Located:  "America",
		Plot:     3305,
	}

	dist.Display()
}
