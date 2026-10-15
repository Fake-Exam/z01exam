package main

import (
	"fmt"
	"os"

	backend "Z02INTRA"
)

func main() {
	if err := backend.Run("runner"); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur runner:", err)
		os.Exit(1)
	}
}
