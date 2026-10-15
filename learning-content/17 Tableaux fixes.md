# Tableaux fixes

## Objectif
Comprendre `[N]T` et son usage (ex: mémoire 10 bytes).

## À savoir (essentiel)
- Un array a une taille dans le type: `[10]byte`.
- Assignation copie tout l’array (valeur).
- On accède par index comme un slice.

## Exemple
```go
package main

import "fmt"

func main() {
	var mem [10]byte
	mem[0] = 'A'
	fmt.Println(mem[0], len(mem)) // 65 10
}
```

## Pièges fréquents
- Confondre `[10]byte` et `[]byte`.
- Passer un array à une fonction en pensant l’éditer sans pointeur (copie).

## Check rapide
- Je sais la différence array vs slice.
- Je sais déclarer `var a [N]T`.
- Je sais pourquoi l’assignation copie.
