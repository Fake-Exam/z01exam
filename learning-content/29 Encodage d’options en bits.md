# Encodage d’options en bits

## Objectif
Encoder des lettres (`a..z`) en flags binaires et afficher par groupes.

## À savoir (essentiel)
- Représentation courante: 26 bits (un par lettre).
- Bit `k` correspond à lettre `'a'+k`.
- Affichage possible en 4 groupes de 8 bits (selon exo `options`).

## Exemple
```go
package main

import "fmt"

func setLetter(mask uint32, c byte) uint32 {
	if c >= 'a' && c <= 'z' {
		mask |= 1 << (c - 'a')
	}
	return mask
}

func main() {
	var m uint32
	m = setLetter(m, 'a')
	m = setLetter(m, 'c')
	fmt.Println(m) // 0b101 (en valeur)
}
```

## Pièges fréquents
- Oublier de filtrer `a..z`.
- Mauvais mapping (off-by-one).
- Format binaire demandé strict (espaces, padding, groupes).

## Check rapide
- Je sais mapper `c-'a'` vers un bit.
- Je sais accumuler un mask.
- Je sais tester la présence d’une option.
