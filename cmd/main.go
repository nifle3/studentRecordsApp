package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	val, err := uuid.Parse("")
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
