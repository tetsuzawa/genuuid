package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func main() {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return
	}
	fmt.Println(u.String())
}
