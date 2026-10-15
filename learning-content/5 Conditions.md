# Conditions

## Objectif
Maîtriser `if` et `switch` pour faire du parsing propre.

## À savoir (essentiel)
- `if` peut inclure une init: `if v := f(); v > 0 {}`.
- `switch` peut être sans expression: `switch { case ...: }`.
- `switch` en Go break automatiquement (pas comme C).

## Exemple
```go
package main

import "fmt"

func main() {
	x := 7
	if x%2 == 0 {
		fmt.Println("pair")
	} else {
		fmt.Println("impair")
	}

	ch := byte('a')
	switch {
	case ch >= 'a' && ch <= 'z':
		fmt.Println("minuscule")
	case ch >= 'A' && ch <= 'Z':
		fmt.Println("majuscule")
	default:
		fmt.Println("autre")
	}
}
```

## Pièges fréquents
- Faire un `switch` “C-like” en espérant un fallthrough (non).
- Confondre conditions ASCII avec Unicode.
- Oublier un `default` quand c’est nécessaire.

## Check rapide
- Je sais utiliser `switch {}` pour classer des caractères.
- Je sais que `switch` ne fallthrough pas par défaut.
- Je sais écrire un `if` avec init sans me tromper de portée.
