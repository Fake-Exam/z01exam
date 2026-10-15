# Parenthèses / brackets

## Objectif
Vérifier si `()[]{}` sont bien équilibrés avec une stack.

## À savoir (essentiel)
- Push quand tu vois un ouvrant.
- Pop et matcher quand tu vois un fermant.
- À la fin: stack vide.

## Exemple
```go
package main

import "fmt"

func ok(s string) bool {
	st := []byte{}
	match := func(open, close byte) bool {
		return (open == '(' && close == ')') ||
			(open == '[' && close == ']') ||
			(open == '{' && close == '}')
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			st = append(st, c)
		} else if c == ')' || c == ']' || c == '}' {
			if len(st) == 0 { return false }
			open := st[len(st)-1]
			st = st[:len(st)-1]
			if !match(open, c) { return false }
		}
	}
	return len(st) == 0
}

func main() { fmt.Println(ok("([])")) } // true
```

## Pièges fréquents
- Pop sur pile vide.
- Oublier de vérifier stack vide à la fin.
- Matcher les mauvais couples.

## Check rapide
- Je sais push sur ouvrant.
- Je sais pop+match sur fermant.
- Je sais valider stack vide à la fin.
