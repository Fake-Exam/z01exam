package main

import (
	"fmt"
	"os"

	backend "Z02INTRA"
)

func main() {
	if err := backend.Run("web"); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur web:", err)
		os.Exit(1)
	}
}
