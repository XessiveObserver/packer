package routines

import (
	"fmt"
	"time"
)

func RunAlong() {
	Runners := []string{"Moses", "Amos", "Abram"}
	for _, Runner := range Runners {
		time.Sleep(time.Second * 1)
		fmt.Printf("%v is running along\n", Runner)
	}
}

func Hunt() {
	go RunAlong()

	fmt.Println("Hunters list")
	for hunters := 0; hunters < 1; hunters++ {
		Names := []string{"Oboi", "Omagul", "Okopit", "Omiina"}
		for _, Name := range Names {
			time.Sleep(time.Second)
			fmt.Println(Name)
		}
	}
}
