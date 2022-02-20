package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	s "strings"
	"time"
)

func main() {
	fruPtr := flag.Bool("fruOnly", false, "Use FRU moves only")

	var moveArg string
	flag.StringVar(&moveArg, "moveCount", "20", "Number of moves: integer")

	flag.Parse()

	i, _ := strconv.Atoi(moveArg)
	moveCount := i

	scramble(moveCount, *fruPtr)
//    cubeState := initCube()
//    fmt.Println(cubeState)
//	cubePrint(cubeState)	
}

func scramble(moveCount int, fruOnly bool) {
	moves := []string{"F", "R", "U"}

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	if fruOnly != true {
		moves = append(moves, "B", "D", "L")
	}

	var algo []string
	prevMove := ""

	for i := 0; i < moveCount; i++ {
		var nextMove string
		for {
			nextMove = moves[seed.Intn(len(moves))]
			if nextMove != prevMove {
				break
			}
		}

		prevMove = nextMove

		var next []string
		next = append([]string{nextMove})

		wide := seed.Intn(2)
		if wide == 1 {
			next = append(next, "w")
		}

		ccw := seed.Intn(2)
		if ccw == 1 {
			next = append(next, "'")
		}

		if ccw == 0 {
			turnCount := seed.Intn(2)
			if turnCount == 1 {
				next = append(next, "2")
			}
		}

        algo = append(algo, s.Join(next, ""))
	}

	fmt.Println(algo)
}
