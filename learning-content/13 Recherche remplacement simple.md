# Recherche/remplacement simple

## Objectif
Remplacer ou détecter un motif en scannant une string.

## À savoir (essentiel)
- Comparer `s[i:i+k]` à un motif (si `i+k <= len(s)`).
- Pour “search & replace”, construire une sortie.
- Pour “match” (ex: `wdmatch`), utiliser deux index.

## Exemple
```go
package main

import "fmt"

func replaceAll(s, old, neu string) string {
	if len(old) == 0 { return s }

	out := make([]byte, 0, len(s))
	newBytes := []byte(neu)

	for i := 0; i < len(s); {
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			out = append(out, newBytes...)
			i += len(old)
		} else {
			out = append(out, s[i])
			i++
		}
	}
	return string(out)
}

func main() { fmt.Println(replaceAll("abab", "ab", "X")) } // "XX"
```

## Pièges fréquents
- Oublier la condition `i+len(old) <= len(s)`.
- Boucle infinie si `old == ""`.
- Remplacement qui “re-scanne” la partie remplacée (si mal fait).

## Check rapide
- Je sais protéger les bornes.
- Je sais construire une sortie en une passe.
- Je sais gérer le cas motif vide.
