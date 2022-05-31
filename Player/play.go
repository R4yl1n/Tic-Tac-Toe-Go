package player

import (
	"fmt"
)

type Player struct {
	Player  string
	symbole string
}

func Play(x Player) {
	var validInput bool
	for validInput = false; validInput != true; {
		fmt.Printf("Player %v where do u wanna place ur symbole %v \n", x.Player, x.symbole)
		var field int
		fmt.Scan(&field)
		if field > 9 || Fields[field-1] == "O" || Fields[field-1] == "X" {
			validInput = false

		} else {
			validInput = true
			Fields[field-1] = x.symbole

		}
	}

}

func CreatePlayer(returnment Player, symbol string) Player {
	returnment.symbole = string(symbol)
	return returnment
}
