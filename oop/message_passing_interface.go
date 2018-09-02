package main

import (
	"fmt"
)

// Cow represents a cow
type Cow struct {
}

// Goat represents a goat
type Goat struct {
}

type vocal interface {
	MakeSound()
}

// MakeSound makes a asound
func (c *Cow) MakeSound() {
	fmt.Println("moo")
}

// MakeSound makes a asound
func (g *Goat) MakeSound() {
	fmt.Println("mehehehehe")
}

// MakeWhateverSound makes whatever sound given an animal
func MakeWhateverSound(item vocal) {
	item.MakeSound()
}

func main() {
	cow := Cow{}
	goat := Goat{}
	MakeWhateverSound(&cow)
	MakeWhateverSound(&goat)
}
