# Fonctions et signatures

## Objectif
Savoir écrire une fonction Go avec la **signature exacte** attendue par l’énoncé.

## À savoir (essentiel)
- Une signature = **nom + paramètres typés + retours typés**.
- En Go, le type est **après** le nom: `a int`, pas `int a`.
- Si plusieurs paramètres ont le même type: `func F(a, b int)`.
- Une fonction peut retourner plusieurs valeurs: `func DivMod(a, b int) (int, int)`.
- Les retours peuvent être nommés: `func F() (n int)` (souvent inutile en exam).
- La casse du nom compte:
    - **Majuscule** = exportée (appelable depuis un autre package).
    - **Minuscule** = locale au package.
- Respecte le package demandé (`main` vs `piscine`) et **ne change pas** les types.

## Exemple

```go
// 1) Retour simple
package piscine

func Add(a, b int) int {
	return a + b
}
```

```go
// 2) Retours multiples
package piscine

func DivMod(a, b int) (int, int) {
	return a / b, a % b
}
```

```go
// 3) Aucun retour (void)
package piscine

import "fmt"

func PrintHello() {
	fmt.Println("Hello")
}
```

## Pièges fréquents
- Se tromper de type (`int` vs `string` vs `[]int`) par “intuition”.
- Oublier un paramètre ou changer l’ordre des paramètres.
- Mettre un `return` manquant ou au mauvais type (compile error).
- Renommer la fonction (ou mauvaise casse) alors que la moulinette appelle un nom précis.
- Ajouter des prints dans un exo “fonction” alors que le correcteur compare juste le retour.

## Check rapide
- Je sais recopier une signature d’énoncé **sans la modifier**.
- Je sais écrire `func Nom(a, b int) int` et `func Nom(...) (int, int)`.
- Je sais quand le nom doit commencer par une majuscule.
