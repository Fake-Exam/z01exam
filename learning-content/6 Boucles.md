# Boucles

## Objectif
Utiliser `for` dans ses 3 formes pour scanner strings et slices.

## À savoir (essentiel)
- Go n’a qu’une boucle: `for`.
- Formes:
    - `for i := 0; i < n; i++ {}`
    - `for condition {}`
    - `for i, v := range xs {}`
- `break` et `continue` existent.

## Exemple
```go
package main

import "fmt"

func main() {
	s := "abc"
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i]) // bytes
	}

	for i, r := range s {
		fmt.Println(i, r) // runes
	}
}
```

## Pièges fréquents
- Utiliser `range` sur string en pensant obtenir des bytes.
- Faire `i <= len(s)` (panic index).
- Oublier d’incrémenter dans `for condition`.

## Check rapide
- Je sais choisir `for i := 0; i < len(s); i++`.
- Je sais quand `range` est dangereux sur string.
- Je sais éviter les index out of range.
