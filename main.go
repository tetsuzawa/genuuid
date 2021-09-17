package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(u.String())
}
