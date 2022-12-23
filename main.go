package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {
	var profile Profile
	if name == "Sasuke" {
		profile = Profile{
			Name:   name,
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	} else if name == "Goku" {
		profile = Profile{
			Name:   name,
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	} else if name == "Naruto" {
		profile = Profile{
			Name:   name,
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	}

	return profile
}

//nilai = nilai + (nilai * multiplier)
func (profile *Profile) PowerUp(multiplier int) {
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)
}

func main() {
	profile := MakeProfile("Naruto")
	fmt.Println("Name:", profile.Name)
	fmt.Println("Health:", profile.Health)
	fmt.Println("Power:", profile.Power)
	fmt.Println("Exp:", profile.Exp)
	fmt.Println("======Heroes Power Up======")
	profile.PowerUp(4)
	fmt.Println("Name:", profile.Name)
	fmt.Println("Health:", profile.Health)
	fmt.Println("Power:", profile.Power)
	fmt.Println("Exp:", profile.Exp)
}
