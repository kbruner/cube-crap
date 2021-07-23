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
    flag.StringVar(&moveArg, "moveRange", "20", "Number of moves: integer or range")

    flag.Parse()

    var moveRange []int
    if s.Contains(moveArg, "-") {
        for _, val := range s.Split(moveArg, "-") {
            i, _ := strconv.Atoi(val)
            moveRange = append(moveRange, i)
        }
    } else {
        i, _ := strconv.Atoi(moveArg)
        moveRange = append(moveRange, i, i)
    }

    scramble(moveRange, *fruPtr)
}

func scramble(moveRange []int, fruOnly bool) {
    moves := []string{"F", "R", "U"}
    modifiers := [3]string{"", "'", "2"}

    seed := rand.New(rand.NewSource(time.Now().UnixNano()))

    if fruOnly != true {
        moves = append(moves, "B", "D", "L")
    }

    var moveCount int
    if moveRange[0] == moveRange[1] {
        moveCount = moveRange[0]
    } else {
        moveCount = seed.Intn(moveRange[1] - moveRange[0] + 1) + moveRange[0]
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
        
        algo = append(algo, s.Join([]string{nextMove, modifiers[seed.Intn(len(modifiers))]}, ""))
        prevMove = nextMove
    }

    fmt.Println(algo)
}
