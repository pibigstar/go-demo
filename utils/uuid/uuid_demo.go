package main

import (
	"github.com/pborman/uuid"
	"log"
)

func main() {

	random := uuid.NewRandom()
	log.Println(random)
}
