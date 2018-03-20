package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	fmt.Println("Eventually this will be a Rock, Paper, Scissors in Go")
	fmt.Println("My favourite number is", random.Intn(99))

}
