/*
This will eventually be a Rock, Paper, Scissors program in Go.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	choices := [3]string{"r", "p", "s"}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	choice1 := random.Intn(3)
	choice2 := random.Intn(3)
	fmt.Println("Eventually this will be a Rock, Paper, Scissors in Go")
	fmt.Println("My guess is", choices[choice1])
	fmt.Println("Your guess is", choices[choice2])

}

// BUG(jbfink): This is actually not playable (yet).
