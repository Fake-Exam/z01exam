package main

import (
	"fmt"
	"os"

	backend "Z02INTRA"
)

func main() {
	if err := backend.Run("catalog"); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur catalogue:", err)
		os.Exit(1)
	}
}
