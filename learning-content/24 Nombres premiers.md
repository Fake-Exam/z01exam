# Nombres premiers

## Objectif
Tester si un nombre est premier et itérer proprement.

## À savoir (essentiel)
- `n <= 1` non premier.
- Tester jusqu’à `i*i <= n`.
- Éviter float.

## Exemple
```go
package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 { return false }
	if n == 2 { return true }
	if n%2 == 0 { return false }
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 { return false }
	}
	return true
}

func main() { fmt.Println(isPrime(29)) } // true
```

## Pièges fréquents
- Boucler jusqu’à `n-1` (trop lent).
- Oublier `n <= 1`.
- Overflow potentiel de `i*i` si n très grand (rare en exos).

## Check rapide
- Je sais tester jusqu’à `i*i <= n`.
- Je sais optimiser en sautant les pairs.
- Je sais gérer `n <= 1`.
