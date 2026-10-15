# Itérer une string

## Objectif
Choisir la bonne manière d’itérer selon ASCII vs Unicode.

## À savoir (essentiel)
- ASCII: boucle index `for i := 0; i < len(s); i++`.
- Unicode: `for _, r := range s` (runes).
- `range` donne l’index byte + rune décodée.

## Exemple
```go
package main

import "fmt"

func main() {
	s := "AZ"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // byte
	}
	for _, r := range s {
		fmt.Println(r) // rune
	}
}
```

## Pièges fréquents
- Utiliser `range` puis faire `s[i]` en pensant que `i` avance de 1.
- Mélanger logique ASCII avec des runes.

## Check rapide
- Je sais quand éviter `range` pour les exos ASCII.
- Je sais ce que retourne `range` sur string.
- Je sais expliquer index byte vs rune.
