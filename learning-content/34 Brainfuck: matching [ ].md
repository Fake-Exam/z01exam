# Brainfuck: matching [ ]

## Objectif
Gérer les boucles `[` `]` en reliant les positions correspondantes.

## À savoir (essentiel)
- `[` : si `mem[p] == 0`, jump après le `]` correspondant.
- `]` : si `mem[p] != 0`, jump après le `[` correspondant.
- Technique fiable: pré-calcul des paires avec une stack + map index->index.

## Exemple
```go
package main

import "fmt"

func buildJumps(code string) map[int]int {
	st := []int{}
	j := make(map[int]int)
	for i := 0; i < len(code); i++ {
		if code[i] == '[' {
			st = append(st, i)
		} else if code[i] == ']' {
			if len(st) == 0 { return nil }
			open := st[len(st)-1]
			st = st[:len(st)-1]
			j[open] = i
			j[i] = open
		}
	}
	if len(st) != 0 { return nil }
	return j
}

func main() { fmt.Println("jumps via stack+map") }
```

## Pièges fréquents
- Pile vide sur `]` ou pile non vide à la fin (brackets invalides).
- Jump au mauvais index (off-by-one).
- Scanner à chaque `[` `]` sans pré-calcul (trop lent et bug-prone).

## Check rapide
- Je sais pré-calculer les couples `[` `]`.
- Je sais utiliser `jumps[ip]` pour sauter.
- Je sais gérer code invalide (crochets non appariés).
