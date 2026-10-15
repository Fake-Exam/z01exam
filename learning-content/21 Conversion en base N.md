# Conversion en base N

## Objectif
Implémenter `ItoaBase` avec un alphabet donné.

## À savoir (essentiel)
- Base = longueur de l’alphabet.
- N doit être positif (selon exo).
- Digit = `n % base`, puis `alphabet[digit]`.
- Même technique que `itoa`, mais base variable.

## Exemple
```go
package main

import "fmt"

func itoaBase(n int, alphabet string) string {
	base := len(alphabet)
	if base < 2 { return "" }
	if n == 0 { return string(alphabet[0]) }

	sign := false
	if n < 0 { sign = true; n = -n }

	var tmp []byte
	for n > 0 {
		tmp = append(tmp, alphabet[n%base])
		n /= base
	}
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	if sign { return "-" + string(tmp) }
	return string(tmp)
}

func main() { fmt.Println(itoaBase(31, "0123456789abcdef")) } // "1f"
```

## Pièges fréquents
- Alphabet invalide (doublons) si l’exo le demande.
- Base < 2.
- Oublier `0`.

## Check rapide
- Je sais calculer `base := len(alphabet)`.
- Je sais extraire digit via `% base`.
- Je sais construire la sortie en inversant.
