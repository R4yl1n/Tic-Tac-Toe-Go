package main

import (
	"fmt"
)

var winnerknown bool = false
var fields = [9]string{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]", "[8]", "[9]"}

func printBoard() {
	fmt.Printf("\n %v | %v | %v \n", fields[0], fields[1], fields[2])
	fmt.Printf("________________\n")
	fmt.Printf(" %v | %v | %v \n", fields[3], fields[4], fields[5])
	fmt.Printf("________________\n")
	fmt.Printf(" %v | %v | %v \n", fields[6], fields[7], fields[8])
	fmt.Printf("\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n")

}

func checkwinner() string {
	if fields[0] == fields[1] && fields[0] == fields[2] {
		winnerknown = true
		return fields[0]

	} else if fields[3] == fields[4] && fields[3] == fields[5] {
		winnerknown = true
		return fields[3]

	} else if fields[6] == fields[7] && fields[6] == fields[8] {
		winnerknown = true
		return fields[6]

	} else if fields[0] == fields[4] && fields[0] == fields[8] {
		winnerknown = true
		return fields[0]

	} else if fields[2] == fields[4] && fields[2] == fields[6] {
		winnerknown = true
		return fields[2]

	} else if fields[0] == fields[3] && fields[0] == fields[6] {
		winnerknown = true
		return fields[0]

	} else if fields[1] == fields[4] && fields[1] == fields[7] {
		winnerknown = true
		return fields[1]

	} else if fields[2] == fields[5] && fields[2] == fields[8] {
		winnerknown = true
		return fields[2]

	} else {
		winnerknown = false
		return "noody"
	}

}

func Play(x player) {
	var validInput bool
	for validInput = false; validInput != true; {
		fmt.Printf("Player %v where do u wanna place ur symbole %v \n", x.Player, x.symbole)
		var field int
		fmt.Scan(&field)
		if field > 9 || fields[field-1] == "O" || fields[field-1] == "X" {
			validInput = false

		} else {
			validInput = true
			fields[field-1] = x.symbole

		}
	}
}

type player struct {
	Player  string
	symbole string
}

func clearBoard() {
	countdownClearBoard := 0
	for i := 0; i < 9; i++ {
		if fields[i] == "O" || fields[i] == "X" {
			countdownClearBoard += 1
		}

		if countdownClearBoard == 9 {
			fields = [9]string{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]", "[8]", "[9]"}
			fmt.Print("Restarting Board \n XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		}
	}
}

func main() {

	var personO player
	personO.symbole = "O"
	var personX player
	personX.symbole = "X"

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
				Play(personO)
				printBoard()
				winnerknownstring = checkwinner()
				clearBoard()
			case 1:
				Play(personX)
				printBoard()
				winnerknownstring = checkwinner()
				clearBoard()
			}
		}
	}
	fmt.Printf("Congratulation %v you won this game", winnerknownstring)
}

// what if nobody winns
