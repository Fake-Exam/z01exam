# GCD (PGCD)

## Objectif
Implémenter l’algorithme d’Euclide.

## À savoir (essentiel)
- `gcd(a,b)`:
    - tant que `b != 0`: `a, b = b, a%b`
- Utiliser valeurs positives.

## Exemple
```go
package main

import "fmt"

func gcd(a, b int) int {
	if a < 0 { a = -a }
	if b < 0 { b = -b }
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() { fmt.Println(gcd(54, 24)) } // 6
```

## Pièges fréquents
- Oublier `a%b` avec `b == 0`.
- Ne pas normaliser signe.
- Croire que gcd(0,0) est défini (à traiter selon exo).

## Check rapide
- Je sais écrire la boucle Euclide.
- Je sais utiliser `%` correctement.
- Je sais normaliser en positif.
