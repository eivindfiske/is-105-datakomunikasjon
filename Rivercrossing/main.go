package main

import (
	"bufio"
	"fmt"
	"os"

	"main.go/event"
	"main.go/state"
)

func main() {

	fmt.Println("Skriv Commands for å se kommandoer")
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Print("\033[H\033[2J")
		//Info Command
		if "Fullstate" == input.Text() {
			state.Fullstate()
		}
		if "Weststate" == input.Text() {
			state.Weststate()
		}
		if "Boatstate" == input.Text() {
			state.Boatstate()
		}
		if "Eaststate" == input.Text() {
			state.Eaststate()
		}
		if "Commands" == input.Text() {
			event.Commands()
		}
		if "Restart" == input.Text() {
			event.Restart()
		}
		if "KyllingtoWest" == input.Text() {
			event.KyllingtoWest()
		}
		if "KyllingtoBoat" == input.Text() {
			event.KyllingtoBoat()
		}
		if "KyllingtoEast" == input.Text() {
			event.KyllingtoEast()
		}
		if "RevtoWest" == input.Text() {
			event.RevtoWest()
		}
		if "RevtoBoat" == input.Text() {
			event.RevtoBoat()
		}
		if "RevtoEast" == input.Text() {
			event.RevtoEast()
		}
		if "KorntoWest" == input.Text() {
			event.KorntoWest()
		}
		if "KorntoBoat" == input.Text() {
			event.KorntoBoat()
		}
		if "KorntoEast" == input.Text() {
			event.KorntoEast()
		}
		if "ManntoWest" == input.Text() {
			event.ManntoWest()
		}
		if "ManntoBoat" == input.Text() {
			event.ManntoBoat()
		}
		if "ManntoEast" == input.Text() {
			event.ManntoEast()
		}
	}
}
