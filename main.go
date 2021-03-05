package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/caioeverest/scarlet-gopher/hex"
)

type Person struct {
	Name      string    `hex:"subject-name" json:"name"`
	Birthdate time.Time `hex:"birthdate" json:"birthdate"`
	Height    float32   `json:"height"`
}

type Car struct {
	Brand      string   `hex:"car-brand" json:"brand"`
	Year       int      `hex:"year" json:"year"`
	Model      string   `hex:"car-model" json:"model"`
	Passengers []Person `json:"passengers"`
	Weight     float32  `json:"weight"`
}

func main() {
	h := hex.Make(1950, 90)
	caio := Person{
		Name:      "Caio Everest",
		Birthdate: time.Date(1995, time.November, 29, 0, 0, 0, 0, time.UTC),
		Height:    1.72,
	}

	Corsa := &Car{
		Brand:      "Chevrolet",
		Year:       2012,
		Model:      "Corsa",
		Passengers: []Person{caio},
		Weight:     1.020,
	}

	fmt.Printf("Old Corsa: %s\n", JsonFormat(Corsa))

	h.Enter(Corsa)

	fmt.Printf("New Corsa: %s\n", JsonFormat(Corsa))
}

func JsonFormat(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
