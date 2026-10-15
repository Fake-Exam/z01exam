# Print au caractère près

## Objectif
Imprimer exactement ce que la moulinette attend.

## À savoir (essentiel)
- `fmt.Print` n’ajoute rien.
- `fmt.Println` ajoute un `\n`.
- `fmt.Printf` pour formater (`%d`, `%s`, `%02x`).
- Beaucoup d’exos échouent à cause d’un espace ou d’un `\n` en trop.

## Exemple
```go
package main

import "fmt"

func main() {
	fmt.Print("A")        // "A"
	fmt.Println("B")      // "B\n"
	fmt.Printf("%02x\n", 15) // "0f\n"
}
```

## Pièges fréquents
- `Println` au lieu de `Print` (newline en trop).
- Ajouter des logs/debug.
- Mauvais format hex (`%x` vs `%02x`).

## Check rapide
- Je sais contrôler les espaces et les `\n`.
- Je sais quand utiliser `Printf`.
- Je sais produire un hex “paddé” sur 2 chars.
