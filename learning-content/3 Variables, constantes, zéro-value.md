# Variables, constantes, zéro-value

## Objectif
Savoir déclarer des variables et éviter les pièges de portée et de type.

## À savoir (essentiel)
- `:=` déclare + assigne (dans une fonction).
- `var x T` déclare avec zéro-value.
- `const` pour valeurs immuables (souvent nombres, strings).
- Zéro-values: `0`, `false`, `""`, `nil` (slice/map/pointer).
- Portée: bloc `{}`. Attention au shadowing avec `:=`.

## Exemple
```go
package main

import "fmt"

func main() {
	var a int        // 0
	var s string     // ""
	b := 42          // int
	const pi = 3.14  // non modifiable

	fmt.Println(a, s, b, pi)
}
```

## Pièges fréquents
- Utiliser `:=` alors qu’on voulait réassigner `=` (shadowing).
- Penser que `const` peut être calculée avec n’importe quoi (non).
- Oublier qu’une `map` non initialisée est `nil` (panic à l’écriture).

## Check rapide
- Je sais quand utiliser `:=` vs `var`.
- Je connais les zéro-values des types courants.
- Je repère un shadowing qui casse la logique.
