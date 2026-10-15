# Types de base + conversions

## Objectif
Comprendre les types Go et caster proprement (surtout ASCII).

## À savoir (essentiel)
- Types clés: `int`, `bool`, `string`, `byte` (`uint8`), `rune` (`int32`).
- Pas de conversion implicite entre types numériques.
- `byte` utile pour ASCII (`s[i]`).
- `rune` utile si tu fais du “caractère” Unicode.
- Conversion: `T(x)`.

## Exemple
```go
package main

import "fmt"

func main() {
	var x int = 65
	b := byte(x)      // 65
	r := rune(b)      // 65
	c := string([]byte{b}) // "A"
	fmt.Println(b, r, c)
}
```

## Pièges fréquents
- Mélanger `byte` et `rune` sans comprendre ce que tu traites.
- Comparer `byte` à `'A'` ok, mais oublier que `'A'` est une rune.
- Caster un `int` négatif vers `byte` (wrap).

## Check rapide
- Je sais expliquer `byte` vs `rune`.
- Je sais caster `int` -> `byte` -> `string`.
- Je sais pourquoi Go refuse les conversions implicites.
