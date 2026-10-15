# Stack via slice

## Objectif
Implémenter une pile (push/pop) avec `[]int` ou `[]rune`.

## À savoir (essentiel)
- Push: `st = append(st, v)`
- Pop: prendre `st[len-1]` puis `st = st[:len-1]`
- Toujours vérifier pile non vide avant pop.

## Exemple
```go
package main

import "fmt"

func main() {
	st := []int{}
	st = append(st, 10)
	st = append(st, 20)

	top := st[len(st)-1]
	st = st[:len(st)-1]
	fmt.Println(top, st) // 20 [10]
}
```

## Pièges fréquents
- Pop sur pile vide (panic).
- Oublier de raccourcir le slice.
- Réutiliser `top` sans copie sur types référence (rare ici).

## Check rapide
- Je sais push/pop sans panic.
- Je sais tester `len(st)`.
- Je sais utiliser une pile pour parsing.
