# Arguments CLI

## Objectif
Lire `os.Args` et gérer les cas invalides proprement.

## À savoir (essentiel)
- `os.Args` inclut le nom du programme en index 0.
- Le 1er vrai argument est `os.Args[1]`.
- Toujours valider `len(os.Args)` avant accès.
- Sortie d’erreur souvent exigée au format exact.

## Exemple
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Error\n")
		return
	}
	fmt.Print(os.Args[1] + "\n")
}
```

## Pièges fréquents
- Oublier que `Args[0]` existe.
- Ne pas gérer `len != attendu`.
- Panic index out of range.

## Check rapide
- Je sais valider `len(os.Args)`.
- Je sais récupérer `os.Args[1]`.
- Je sais imprimer l’erreur demandée sans extras.
