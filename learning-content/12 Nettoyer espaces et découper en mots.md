# Nettoyer espaces et découper en mots

## Objectif
Gérer les espaces comme dans `cleanstr`, `expandstr`, `rostring`.

## À savoir (essentiel)
- Pattern classique:
    - skipper espaces au début
    - lire un mot
    - skipper espaces
    - ajouter un séparateur unique si besoin
- Un “espace” d’exam = souvent `' '` uniquement.

## Exemple
```go
package main

import "fmt"

func cleanSpaces(s string) string {
	out := make([]byte, 0, len(s))
	i := 0

	// skip leading spaces
	for i < len(s) && s[i] == ' ' { i++ }

	for i < len(s) {
		// read word
		for i < len(s) && s[i] != ' ' {
			out = append(out, s[i])
			i++
		}
		// skip spaces
		for i < len(s) && s[i] == ' ' { i++ }
		if i < len(s) {
			out = append(out, ' ')
		}
	}
	return string(out)
}

func main() { fmt.Println(cleanSpaces("  a   b  c ")) }
```

## Pièges fréquents
- Produire un espace final.
- Ne pas gérer la string vide ou “que des espaces”.
- Traiter tab `\t` alors que l’énoncé ne le demande pas (ou l’inverse).

## Check rapide
- Je sais extraire des mots en scannant.
- Je sais normaliser les séparateurs.
- Je sais éviter espace au début et à la fin.
