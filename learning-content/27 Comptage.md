# Comptage

## Objectif
Compter occurrences et classer des caractères.

## À savoir (essentiel)
- Pour ASCII: tableau `[256]int` rapide.
- Pour lettres `a-z`: `[26]int`.
- Parcours unique.

## Exemple
```go
package main

import "fmt"

func countAlpha(s string) [26]int {
	var c [26]int
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= 'a' && b <= 'z' {
			c[b-'a']++
		}
	}
	return c
}

func main() { fmt.Println(countAlpha("abca")[0], countAlpha("abca")[1]) } // a=2, b=1
```

## Pièges fréquents
- Index négatif (si tu ne testes pas le range).
- Confondre `byte` et `rune`.
- Oublier que majuscules et minuscules sont différents.

## Check rapide
- Je sais compter avec un tableau fixe.
- Je sais indexer `b-'a'`.
- Je sais filtrer ASCII avant de compter.
