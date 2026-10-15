# Renverser / rotation / zip

## Objectif
Maîtriser les transformations classiques de string.

## À savoir (essentiel)
- Reverse ASCII: lire de fin à début.
- Rotation (ex: `rostring`): extraire mots, déplacer le 1er à la fin.
- Zip (ex: `zipstring`): regrouper des répétitions, ou intercaler selon l’énoncé.

## Exemple
```go
package main

import "fmt"

func reverseASCII(s string) string {
	out := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return string(out)
}

func main() { fmt.Println(reverseASCII("abcd")) } // "dcba"
```

## Pièges fréquents
- Reverse avec `range` sur Unicode (résultat inattendu).
- Rotation: espaces multiples, mot vide.
- Zip: oublier de gérer le dernier groupe.

## Check rapide
- Je sais reverse une string ASCII sans panic.
- Je sais faire une rotation de mots en une passe de parsing.
- Je sais “zipper” en gérant le dernier segment.
