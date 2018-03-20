package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	fmt.Println("Eventually this will be a Rock, Paper, Scissors in Go")
	fmt.Println("My favourite number is", rand.Intn(99))

}
