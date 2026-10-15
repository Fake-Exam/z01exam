# Factorisation

## Objectif
Factoriser un entier (ex: `fprime`) en divisant par petits facteurs.

## À savoir (essentiel)
- Diviser par 2 puis par impairs.
- Tant que `n % f == 0`, sortir `f` et diviser `n`.
- À la fin, si `n > 1`, c’est un facteur.

## Exemple
```go
package main

import "fmt"

func factors(n int) []int {
	var out []int
	if n < 0 { n = -n }
	for n%2 == 0 {
		out = append(out, 2)
		n /= 2
	}
	for f := 3; f*f <= n; f += 2 {
		for n%f == 0 {
			out = append(out, f)
			n /= f
		}
	}
	if n > 1 { out = append(out, n) }
	return out
}

func main() { fmt.Println(factors(60)) } // [2 2 3 5]
```

## Pièges fréquents
- Oublier la boucle “tant que divisible”.
- Ne pas gérer `n` négatif si demandé.
- Format de sortie (espaces, `*`) selon l’énoncé.

## Check rapide
- Je sais factoriser par divisions successives.
- Je sais m’arrêter à `f*f <= n`.
- Je sais gérer le dernier facteur `n > 1`.
