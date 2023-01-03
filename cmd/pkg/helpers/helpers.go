package helpers

import (
	"log"
	"math/rand"
	"time"
)

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {
	log.Println("Random Pool", n)
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	log.Println("Random Number", value)

	return value
}
