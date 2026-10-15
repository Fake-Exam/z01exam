# Atoi “maison”

## Objectif
Convertir une string ASCII en int, avec validation simple.

## À savoir (essentiel)
- Scanner char par char.
- `n = n*10 + int(c-'0')`.
- Gérer optionnellement `+`/`-` selon l’énoncé.
- Détecter invalides: char hors `'0'..'9'`.

## Exemple
```go
package main

import "fmt"

func atoi(s string) (int, bool) {
	if len(s) == 0 { return 0, false }
	sign := 1
	i := 0
	if s[0] == '-' { sign = -1; i = 1 }
	if i == len(s) { return 0, false }

	n := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' { return 0, false }
		n = n*10 + int(s[i]-'0')
	}
	return sign * n, true
}

func main() { fmt.Println(atoi("-42")) }
```

## Pièges fréquents
- Accepter des strings vides ou juste `"-"`.
- Oublier la validation des chars.
- Overflow int (souvent ignoré en exos, mais à connaître).

## Check rapide
- Je sais parser des digits avec `n*10 + ...`.
- Je sais refuser une entrée invalide.
- Je sais gérer un signe si demandé.
    