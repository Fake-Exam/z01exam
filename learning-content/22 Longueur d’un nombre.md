# Longueur d’un nombre

## Objectif
Calculer le nombre de chiffres d’un int (souvent en base 10).

## À savoir (essentiel)
- Cas `0` => longueur 1.
- Convertir `n` en positif si nécessaire.
- Compter en divisant par 10.

## Exemple
```go
package main

import "fmt"

func digitLen(n int) int {
	if n == 0 { return 1 }
	if n < 0 { n = -n }
	c := 0
	for n > 0 {
		c++
		n /= 10
	}
	return c
}

func main() { fmt.Println(digitLen(-1200)) } // 4
```

## Pièges fréquents
- Oublier `0`.
- Oublier le signe.
- Confondre “digits” et “valeur absolue”.

## Check rapide
- Je sais traiter `0`.
- Je sais compter avec division entière.
- Je sais gérer le signe.
