package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

type Player struct {
	Guess int
}

type Game struct {
	Random	int
}

func createPlayer(guess int) Player {
	player := Player{Guess: guess}

	return player
}

func createGame() *Game {
	game := Game{Random: rand.Intn(10)}

	return &game
}

func main() {
	for {
		var guess1str string
		var guess2str string
		var again string
		currentGame := createGame()
		fmt.Println("Player 1, your guess")
		fmt.Scanf("%s", &guess1str)
		guess1, _ := strconv.Atoi(guess1str)
		player1 := createPlayer(guess1)
		fmt.Println("Player 2, your guess")
		fmt.Scanf("%s", &guess2str)
		guess2, _ := strconv.Atoi(guess2str)
		player2 := createPlayer(guess2)
		delta1 := math.Abs(float64(currentGame.Random - player1.Guess))
		delta2 := math.Abs(float64(currentGame.Random - player2.Guess))
		if delta1 < delta2 {
			fmt.Println("Player 1 wins !")
		} else if delta2 < delta1 {
			fmt.Println("Player 2 wins !")
		} else {
			fmt.Println("It's a tie")
		}
		fmt.Println("Result : " + strconv.Itoa(currentGame.Random))
		fmt.Println("=====")
		fmt.Println("Do you want to play again ? (y/n)")
		fmt.Scanf("%s", &again)
		if again != "y" {
			break
		}
	}
}
