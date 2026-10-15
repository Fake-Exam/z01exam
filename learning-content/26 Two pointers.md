# Two pointers

## Objectif
Scanner 2 strings/slices avec 2 index (ex: `wdmatch`, matching).

## À savoir (essentiel)
- Deux index `i` et `j` avancent indépendamment.
- Pattern “subsequence”:
    - si match: `i++`, `j++`
    - sinon: `j++`
- O(n) sans backtracking.

## Exemple
```go
package main

import "fmt"

func isSubsequence(a, b string) bool {
	i := 0
	for j := 0; j < len(b) && i < len(a); j++ {
		if a[i] == b[j] { i++ }
	}
	return i == len(a)
}

func main() { fmt.Println(isSubsequence("abc", "aXbYc")) } // true
```

## Pièges fréquents
- Accéder `a[i]` quand `i == len(a)`.
- Confondre “substring” vs “subsequence”.
- Mauvaise gestion des strings vides.

## Check rapide
- Je sais écrire un scan à deux index.
- Je sais éviter les out-of-range.
- Je sais reconnaître “subsequence”.
