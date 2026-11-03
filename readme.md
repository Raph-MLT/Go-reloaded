# go-reloaded

Programme Go qui corrige automatiquement un texte.

## Lancer le programme

```
go run . sample.txt result.txt
```

Le programme lit `sample.txt`, corrige le texte et écrit le résultat dans `result.txt`.

## Fonctionnement

`main.go` lit le fichier d'entrée et passe son contenu à la fonction `process`, dans `texte.go`. Le texte y est coupé en lignes, puis chaque ligne en mots. Les mots passent ensuite par quatre étapes, dans cet ordre :

1. `markers.go` applique `(hex)`, `(bin)`, `(up)`, `(low)` et `(cap)` au mot précédent, ou aux n mots précédents avec `(up, n)`, puis retire le marqueur.
2. `punctuation.go` colle `. , ! ? : ;` au mot d'avant.
3. `quotes.go` colle les apostrophes aux mots qu'elles entourent.
4. `article.go` remplace `a` par `an` devant une voyelle ou un h.

Les mots sont recollés avec un espace entre chacun, puis `main.go` écrit le résultat dans le fichier de sortie.

Les marqueurs passent avant la ponctuation. Dans `go (up) !`, si le `!` était collé en premier, il finirait sur `(up)` et le programme ne reconnaîtrait plus le marqueur.

## Exemple

```
Entrée : Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
Sortie : Simply add 66 and 2 and you will see the result is 68.
```

## Tests

Les tests sont dans `main_test.go` :


go test -v