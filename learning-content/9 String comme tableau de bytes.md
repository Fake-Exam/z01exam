# String comme tableau de bytes

## Objectif
Manipuler une string en mode ASCII avec `s[i]`.

## À savoir (essentiel)
- Une `string` est immuable.
- `s[i]` retourne un `byte`.
- `len(s)` retourne le nombre de bytes.
- Parfait pour ASCII (la majorité de tes exos).

## Exemple
```go
package main

import "fmt"

func main() {
	s := "Hello"
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
```

## Pièges fréquents
- Penser que `len(s)` = nombre de “lettres” (faux en Unicode).
- Essayer de modifier `s[i]` (impossible).
- Confondre `byte` et `rune`.

## Check rapide
- Je sais lire un caractère ASCII avec `s[i]`.
- Je sais que `string` est immuable.
- Je sais quand `len` parle en bytes.
