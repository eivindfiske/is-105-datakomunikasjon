package event

import (
	"fmt"

	"main.go/state"
)

var Westverdi int = 0
var Boatverdi int = 0
var Eastverdi int = 0
var Kyllingverdi int = 1
var Revverdi int = 2
var Kornverdi int = 9
var Mannverdi int = 4

//restart når du taper
func Restart() {
	state.KyllingWest = true
	state.RevWest = true
	state.KornWest = true
	state.MannWest = true
	state.KyllingBoat = false
	state.RevBoat = false
	state.KornBoat = false
	state.MannBoat = false
	state.KyllingEast = false
	state.RevEast = false
	state.KornEast = false
	state.MannEast = false
}
func Commands() {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("For å se States Skriv Fullstate/Weststate/Boatstate/Eaststate")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("For å flytte på Kyllingen, bruK: KyllingtoWest/KyllingtoBoat/KyllingtoEast")
	fmt.Println("For å flytte på Reven, bruk: RevtoWest(RevtoBoat/RevtoEast")
	fmt.Println("For å flytte på Kornet, bruk: KorntoWest/KorntoBoat/KorntoEast")
	fmt.Println("For å flytte på mannen, bruk: ManntoWest/ManntoBoast/ManntoEast")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("For å restarte, bruk: Restart")
}

//flytte kyllingen
func KyllingtoWest() {
	state.KyllingWest = true
	state.KyllingBoat = false
	state.KyllingEast = false

}
func KyllingtoBoat() {
	state.KyllingWest = false
	state.KyllingBoat = true
	state.KyllingEast = false
}
func KyllingtoEast() {
	state.KyllingWest = false
	state.KyllingBoat = false
	state.KyllingEast = true
}

//Flytte på reven
func RevtoWest() {
	state.RevWest = true
	state.RevBoat = false
	state.RevEast = false
}
func RevtoBoat() {
	state.RevWest = false
	state.RevBoat = true
	state.RevEast = false
}
func RevtoEast() {
	state.RevWest = false
	state.RevBoat = false
	state.RevEast = true
}

//flytte kornet
func KorntoWest() {
	state.KornWest = true
	state.KornBoat = false
	state.KornEast = false
}
func KorntoBoat() {
	state.KornWest = false
	state.KornBoat = true
	state.KornEast = false
}
func KorntoEast() {
	state.KornWest = false
	state.KornBoat = false
	state.KornEast = true
}

//flytte mannen
func ManntoWest() {
	state.MannWest = true
	state.MannBoat = false
	state.MannEast = false
}
func ManntoBoat() {
	state.MannWest = false
	state.MannBoat = true
	state.MannEast = false
}
func ManntoEast() {
	state.MannWest = false
	state.MannBoat = false
	state.MannEast = true
}
