# Itoa “maison”

## Objectif
Convertir un `int` en `string` sans `strconv`.

## À savoir (essentiel)
- Cas `0` à traiter.
- Pour construire: extraire digits avec `% 10`, puis inverser.
- Gérer le signe si nécessaire.

## Exemple
```go
package main

import "fmt"

func itoa(n int) string {
	if n == 0 { return "0" }
	sign := false
	if n < 0 { sign = true; n = -n }

	var tmp []byte
	for n > 0 {
		tmp = append(tmp, byte('0'+n%10))
		n /= 10
	}
	// reverse
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	if sign { return "-" + string(tmp) }
	return string(tmp)
}

func main() { fmt.Println(itoa(-123)) }
```

## Pièges fréquents
- Oublier le cas `0`.
- Inverser mal le buffer.
- `-n` overflow sur min int (rarement testé).

## Check rapide
- Je sais construire un nombre en digits inversés.
- Je sais reverse un `[]byte`.
- Je sais gérer le signe.
