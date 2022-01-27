package state

import (
	"fmt"
	"strconv"
)

// starter på west siden
var KyllingWest = true
var RevWest = true
var KornWest = true
var MannWest = true

//båten
var KyllingBoat = false
var RevBoat = false
var KornBoat = false
var MannBoat = false

//avslutter på øst siden
var KyllingEast = false
var RevEast = false
var KornEast = false
var MannEast = false

func Fullstate() {
	fmt.Println("West siden av elven")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Kyllingen er på west sida = " + strconv.FormatBool(KyllingWest))
	fmt.Println("Reven er på west sida = " + strconv.FormatBool(RevWest))
	fmt.Println("Kornet er på west sida = " + strconv.FormatBool(KornWest))
	fmt.Println("Mannen er på west sidan = " + strconv.FormatBool(MannWest))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Kyllingen er i båten = " + strconv.FormatBool(KyllingBoat))
	fmt.Println("Reven er i båten = " + strconv.FormatBool(RevBoat))
	fmt.Println("Kornet er i båten = " + strconv.FormatBool(KornBoat))
	fmt.Println("Mannen er i båten = " + strconv.FormatBool(MannBoat))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Kyllingen er på west sida = " + strconv.FormatBool(KyllingEast))
	fmt.Println("Reven er på west sida = " + strconv.FormatBool(RevEast))
	fmt.Println("Kornet er på west sida = " + strconv.FormatBool(KornEast))
	fmt.Println("Mannen er på west sidan = " + strconv.FormatBool(MannEast))
	fmt.Println("----------------------------------------------------------------")
}

func Weststate() {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Kyllingen er på west sida = " + strconv.FormatBool(KyllingWest))
	fmt.Println("Reven er på west sida = " + strconv.FormatBool(RevWest))
	fmt.Println("Kornet er på west sida = " + strconv.FormatBool(KornWest))
	fmt.Println("Mannen er på west sidan = " + strconv.FormatBool(MannWest))
	fmt.Println("----------------------------------------------------------------")
}

func Boatstate() {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Kyllingen er i båten = " + strconv.FormatBool(KyllingBoat))
	fmt.Println("Reven er i båten = " + strconv.FormatBool(RevBoat))
	fmt.Println("Kornet er i båten = " + strconv.FormatBool(KornBoat))
	fmt.Println("Mannen er i båten = " + strconv.FormatBool(MannBoat))
	fmt.Println("----------------------------------------------------------------")
}

func Eaststate() {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Kyllingen er på east sida = " + strconv.FormatBool(KyllingEast))
	fmt.Println("Reven er på east sida = " + strconv.FormatBool(RevEast))
	fmt.Println("Kornet er på east sida = " + strconv.FormatBool(KornEast))
	fmt.Println("Mannen er på east sida = " + strconv.FormatBool(MannEast))
	fmt.Println("----------------------------------------------------------------")
}
