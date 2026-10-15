package main

import (
	"fmt"
	"os"

	backend "Z02INTRA"
)

func main() {
	if err := backend.Run("monolith"); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur serveur:", err)
		os.Exit(1)
	}
}
