package player

import (
	"fmt"
)

type Player struct {
	Player  string
	Symbole string
}

func Play(x Player, Fields *[9]string) {
	var validInput bool
	for validInput = false; validInput != true; {
		fmt.Printf("Player %v where do u wanna place ur symbole %v \n", x.Player, x.Symbole)
		var field int
		fmt.Scan(&field)
		if field > 9 || (*Fields)[field-1] == "O" || (*Fields)[field-1] == "X" {
			validInput = false

		} else {
			validInput = true
			(*Fields)[field-1] = x.Symbole

		}
	}

}
