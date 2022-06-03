package main

import (
	"fmt"

	"github.com/R4yl1n/Tic-Tac-Toe-Go/player"
)

var winnerknown bool = false
var Fields = [9]string{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]", "[8]", "[9]"}

func printBoard() {
	fmt.Printf("\n %v | %v | %v \n", Fields[0], Fields[1], Fields[2])
	fmt.Printf("________________\n")
	fmt.Printf(" %v | %v | %v \n", Fields[3], Fields[4], Fields[5])
	fmt.Printf("________________\n")
	fmt.Printf(" %v | %v | %v \n", Fields[6], Fields[7], Fields[8])
	fmt.Printf("\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n")

}

func checkwinner() string {
	if Fields[0] == Fields[1] && Fields[0] == Fields[2] {
		winnerknown = true
		return Fields[0]

	} else if Fields[3] == Fields[4] && Fields[3] == Fields[5] {
		winnerknown = true
		return Fields[3]

	} else if Fields[6] == Fields[7] && Fields[6] == Fields[8] {
		winnerknown = true
		return Fields[6]

	} else if Fields[0] == Fields[4] && Fields[0] == Fields[8] {
		winnerknown = true
		return Fields[0]

	} else if Fields[2] == Fields[4] && Fields[2] == Fields[6] {
		winnerknown = true
		return Fields[2]

	} else if Fields[0] == Fields[3] && Fields[0] == Fields[6] {
		winnerknown = true
		return Fields[0]

	} else if Fields[1] == Fields[4] && Fields[1] == Fields[7] {
		winnerknown = true
		return Fields[1]

	} else if Fields[2] == Fields[5] && Fields[2] == Fields[8] {
		winnerknown = true
		return Fields[2]

	} else {
		winnerknown = false
		return "noody"
	}

}

func clearBoard() {
	countdownClearBoard := 0
	for i := 0; i < 9; i++ {
		if Fields[i] == "O" || Fields[i] == "X" {
			countdownClearBoard += 1
		}

		if countdownClearBoard == 9 {
			Fields = [9]string{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]", "[8]", "[9]"}
			fmt.Println("Restarting Board \n XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			printBoard()
		}
	}
}

func main() {

	personO := player.CreateSymbole("O")
	personX := player.CreateSymbole("X")
	fmt.Printf("Player one please insert your name, you are the symbol O \n")
	fmt.Scan(&personO.Player)

	fmt.Printf("Player two please insert your name, you are the symbol X \n")
	fmt.Scan(&personX.Player)
	printBoard()
	fmt.Printf("welcome %v and %v to my TicTacToe game ist complete written in go so please enjoy \n", personO.Player, personX.Player)
	var winnerknownstring string
	checkwinner()
	for winnerknown = false; winnerknown != true; {
		for i := 0; i < 3; i++ {
			if winnerknown == true {

				break

			}
			switch i {

			case 0:
				player.Play(personO, (&Fields))
				printBoard()
				winnerknownstring = checkwinner()
				clearBoard()
			case 1:
				player.Play(personX, (&Fields))
				printBoard()
				winnerknownstring = checkwinner()
				clearBoard()
			}
		}
	}
	fmt.Printf("Congratulation %v you won this game", winnerknownstring)
}
