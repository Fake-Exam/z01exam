# Transformations de casse

## Objectif
Faire maj/min/capitalisation en ASCII.

## À savoir (essentiel)
- ASCII:
    - min: `'a'..'z'`
    - maj: `'A'..'Z'`
- Convertir:
    - min -> maj: `c - 'a' + 'A'`
    - maj -> min: `c - 'A' + 'a'`
- Capitaliser: 1ère lettre d’un mot en maj, le reste selon règle de l’exo.

## Exemple
```go
package main

import "fmt"

func toUpperASCII(b byte) byte {
	if b >= 'a' && b <= 'z' { return b - 'a' + 'A' }
	return b
}

func main() { fmt.Println(string([]byte{toUpperASCII('b')})) } // "B"
```

## Pièges fréquents
- Appliquer ASCII sur Unicode.
- Oublier que “mot” dépend de la définition (espace seulement, ou ponctuation aussi).
- Modifier des chiffres/symboles par erreur.

## Check rapide
- Je sais tester `a..z` et `A..Z`.
- Je sais convertir une lettre ASCII.
- Je sais détecter un début de mot en scan.
