# Brainfuck: mémoire + pointeur

## Objectif
Simuler la “tape” Brainfuck: mémoire byte + pointeur.

## À savoir (essentiel)
- Mémoire typique: `[]byte` (souvent 2048 ou 30000 selon exo).
- Pointeur `p` se déplace avec `>` `<`.
- `+` `-` modifie `mem[p]` (wrap byte).
- `.` imprime `mem[p]` comme char (souvent ASCII).

## Exemple
```go
package main

import "fmt"

func main() {
	mem := make([]byte, 10)
	p := 0

	mem[p]++
	mem[p]++
	fmt.Printf("%c\n", mem[p]) // char de valeur 2 (non imprimable)
}
```

## Pièges fréquents
- Sortir des bornes du pointeur (à gérer selon énoncé).
- Confondre `byte` et `int`.
- Mauvaise gestion du wrap (byte wrap naturel, mais attention si tu cast).

## Check rapide
- Je sais représenter la tape avec `[]byte`.
- Je sais appliquer `+ - > < .`.
- Je sais éviter out-of-range du pointeur.
