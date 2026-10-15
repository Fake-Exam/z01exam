# Structure d’un fichier Go

## Objectif
Comprendre la structure minimale d’un fichier Go et ce que le compilateur attend.

## À savoir (essentiel)
- Un fichier Go commence par `package ...`.
- Un exo “fonction” utilise souvent `package piscine` et pas de `main()`.
- Un exo “programme” utilise `package main` et doit avoir `func main()`.
- Les imports se font avec `import "x"` ou un bloc `import ( ... )`.
- Une fonction exportée commence par une majuscule (`PrintStr`).
- Les fichiers d’un même dossier doivent avoir le même package.
- `go.mod` définit le module. En exam, il peut être absent selon la plateforme.

## Exemple

```go
// Programme (exo qui demande d'afficher quelque chose)
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

```go
// Librairie (exo qui demande "Write a function")
// Le correcteur importera ton package et appellera ta fonction.
package piscine

func Add(a, b int) int {
	return a + b
}
```

## Pièges fréquents
- Mettre `package main` alors que l’énoncé attend une fonction dans `package piscine`.
- Ajouter un `main()` dans un exo “fonction” (ça casse le grading selon la moulinette).
- Oublier un import nécessaire, ou importer sans l’utiliser (erreur de compilation).
- Mélanger plusieurs package dans le même dossier.
- Mauvaise casse du nom de fonction (maj/min) par rapport à l’énoncé.

## Check rapide
- Je sais choisir entre `package main` et `package piscine`.
- Je sais quand je dois écrire `func main()` et quand je ne dois pas.
- Je sais où placer `import` et pourquoi un import inutilisé fait échouer la compile.
