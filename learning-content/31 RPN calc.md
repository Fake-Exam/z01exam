# RPN calc

## Objectif
Évaluer une expression en notation polonaise inversée avec une stack.

## À savoir (essentiel)
- Scanner tokens.
- Si nombre: push.
- Si opérateur: pop 2, appliquer, push résultat.
- À la fin: pile doit contenir 1 valeur.

## Exemple
```go
package main

import "fmt"

func evalRPN(tokens []string) (int, bool) {
	st := []int{}
	// pseudo: parsing nombres + opérateurs (détails selon énoncé)
	_ = st
	return 0, false
}

func main() { fmt.Println("RPN = stack + parsing") }
```

## Pièges fréquents
- Ordre des pops: `b := pop(); a := pop()` puis `a op b`.
- Division par zéro.
- Tokens invalides ou pile pas au bon état.

## Check rapide
- Je sais le workflow nombre vs opérateur.
- Je sais gérer “pile insuffisante”.
- Je sais valider “pile finale = 1”.
