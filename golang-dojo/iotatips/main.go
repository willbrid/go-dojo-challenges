package main

import "fmt"

type Planet int

const (
	Mercury Planet = iota + 1
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
)

// planetName returns the name of the given planet and a boolean indicating
// whether the planet is known.
func planetName(p Planet) (string, bool) {
	switch p {
	case Mercury:
		return "Mercure", true
	case Venus:
		return "Vénus", true
	case Earth:
		return "Terre", true
	case Mars:
		return "Mars", true
	case Jupiter:
		return "Jupiter", true
	case Saturn:
		return "Saturne", true
	case Uranus:
		return "Uranus", true
	case Neptune:
		return "Neptune", true
	default:
		return "", false
	}
}

func main() {
	for p := Mercury; p <= Neptune; p++ {
		planet, exists := planetName(p)
		if exists {
			fmt.Println(planet)
		}
	}

	_, exists := planetName(99)
	if !exists {
		fmt.Println("this planet does not exist")
	}
}
