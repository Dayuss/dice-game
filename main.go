package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// get value from input
	activePlayer := Input("Jumlah player")
	diceCount := Input("Jumlah dadu")

	// play the game
	play(activePlayer, diceCount)

}

func play(activePlayer int, diceCount int) {
	players := make(map[int]int)
	points := make(map[int]int)
	round := 0

	// mapping active player
	for p := 0; p < activePlayer; p++ {
		players[p+1] = diceCount
		points[p+1] = 0
	}

	// game begin
	for len(players) > 1 {

		// increase round
		round++

		fmt.Printf("Giliran %d lempar dadu:\n", round)

		evalResult := make(map[int][]int)
		for player, dice := range players {
			turnResult := []int{}
			diceTurn := dice
			// run dice
			for d := 0; d < diceTurn; d++ {
				turnResult = append(turnResult, getDice())
			}

			fmt.Printf("Pemain %d (%d):%+v\n", player, points[player], turnResult)

			// eval dice result
			for _, dice := range turnResult {
				if dice == 6 {

					// remove 6 and increase point
					turnResult = removeInt(turnResult, 6)

					//increase point
					points[player] = points[player] + 1

				} else if dice == 1 {

					// remove 1
					turnResult = removeInt(turnResult, 1)

					// get index of next player
					nextPlayer := (player + 1) % len(players)
					if nextPlayer == 0 {
						nextPlayer = 1
					}

					// set 1 for next player
					evalResult[nextPlayer] = append(evalResult[nextPlayer], 1)
				}
			}

			// merge eval result & current result
			evalResult[player] = append(turnResult, evalResult[player]...)
		}

		fmt.Println("Setelah evaluasi")
		for player := range players {

			// set dice count per player
			players[player] = len(evalResult[player])

			// show evaluated result
			fmt.Printf("Pemain %d (%d):%+v\n", player, points[player], evalResult[player])

			// check dice count player
			if players[player] == 0 {
				delete(players, player)
			}
		}

		fmt.Println("===========================================")

		// check end game
		if len(players) == 1 {
			// game over, calculate point for the winner
			winner := checkWinner(points)
			lastPlayer := checkLastPlayer(players)

			fmt.Printf("%s\n", lastPlayer)

			fmt.Printf("%s\n", winner)
		}

	}
}

func Input(title string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", title)
	in, _ := reader.ReadString('\n')
	in = strings.TrimRight(in, "\r\n")

	intOfInput, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println("Please input an number")
		os.Exit(0)
	}
	return intOfInput
}

func getDice() int {
	return rand.Intn(6) + 1
}

func removeInt(slices []int, val int) []int {
	var result []int

	for _, v := range slices {
		if v != val {
			result = append(result, v)
		}
	}

	return result
}

func checkWinner(data map[int]int) string {
	maxVal := 0
	maxIndex := 0
	for index, value := range data {
		if value > maxVal {
			maxVal = value
			maxIndex = index
		}
	}
	return fmt.Sprintf("Game dimenangkan oleh pemain %d karena memiliki poin lebih banyak dari pemain lainnya.", maxIndex)
}

func checkLastPlayer(data map[int]int) string {
	getIndex := 0
	for index, _ := range data {
		getIndex = index
	}
	return fmt.Sprintf("Game berakhir karena hanya pemain %d yang memiliki dadu.", getIndex)
}
