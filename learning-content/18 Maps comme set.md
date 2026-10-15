# Maps comme set

## Objectif
Gérer unicité et appartenance rapidement (ex: `union`, `inter`, `weareunique`).

## À savoir (essentiel)
- “Set” = `map[byte]bool` ou `map[rune]bool`.
- Test présence: `_, ok := m[k]`.
- `make(map[T]U)` pour initialiser.

## Exemple
```go
package main

import "fmt"

func uniqueASCII(s string) string {
	seen := make(map[byte]bool)
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !seen[c] {
			seen[c] = true
			out = append(out, c)
		}
	}
	return string(out)
}

func main() { fmt.Println(uniqueASCII("abac")) } // "abc"
```

## Pièges fréquents
- Écrire dans une map `nil` (panic).
- Utiliser `range` et croire à un ordre stable (non garanti).
- Choisir `rune` alors que tout est ASCII (ou l’inverse).

## Check rapide
- Je sais faire un set avec `map`.
- Je sais tester présence avec `ok`.
- Je sais conserver l’ordre en scannant la string.
