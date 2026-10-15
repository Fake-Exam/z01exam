# Slices

## Objectif
Utiliser `[]T` pour stocker et manipuler des listes.

## À savoir (essentiel)
- Un slice = vue sur un tableau (ptr, len, cap).
- `append` peut réallouer.
- Slicing: `a[i:j]` avec `0 <= i <= j <= len(a)`.
- Copier: `copy(dst, src)`.

## Exemple
```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[:2]        // [1 2]
	b = append(b, 9)  // peut réallouer
	fmt.Println(a, b)
}
```

## Pièges fréquents
- Modifier un sous-slice et impacter l’original (même backing array).
- Oublier les bornes dans `a[i:j]`.
- Supposer que `append` modifie toujours en place.

## Check rapide
- Je sais expliquer len vs cap.
- Je sais faire `a[i:j]` sans crash.
- Je sais quand utiliser `copy`.
