# Bitwise bases

## Objectif
Comprendre `& | ^ << >>` pour manipuler des flags.

## À savoir (essentiel)
- Mettre un bit: `x |= (1 << k)`
- Tester un bit: `(x&(1<<k)) != 0`
- Enlever un bit: `x &^= (1 << k)` (AND NOT)

## Exemple
```go
package main

import "fmt"

func main() {
	var x uint32 = 0
	x |= 1 << 2
	fmt.Println((x&(1<<2)) != 0) // true
	x &^= 1 << 2
	fmt.Println(x) // 0
}
```

## Pièges fréquents
- Décaler avec un `k` négatif (impossible).
- Utiliser `int` et overflow mental.
- Confondre `^` (XOR) et “power”.

## Check rapide
- Je sais set/test/clear un bit.
- Je sais ce que fait `&^`.
- Je sais lire une expression bitwise.
