# Convertisseur de température en ligne de commande.

```
$ go run . 100 C
100.0°C = 212.0°F
$ go run . 32 F
32.0°F = 0.0°C
$ go run .
usage: tempconv <valeur> <C|F>
```

### Contraintes

1. Sans argument : message d'usage sur **`os.Stderr`** (pas `os.Stdout`) et sortie avec le **code 1** (`os.Exit(1)`). Chercher pourquoi cette distinction compte en shell — c'est une exigence de « prêt pour la production », pas une coquetterie.
2. L'unité est acceptée en majuscule comme en minuscule.
3. Une valeur non numérique produit un message d'erreur clair, pas un plantage.
4. La conversion vit dans **une fonction séparée** de `main`, qui **n'affiche rien**.
5. Exactement une décimale à l'affichage.

*Indices : `strconv.ParseFloat`, `strings.ToUpper`, `fmt.Fprintf` (le `F` = *File* : écrit vers une destination au choix), le verbe `%.1f`.*

**Le point caché de l'exercice :** séparer le calcul des entrées/sorties. Une fonction qui calcule *et* qui affiche est presque intestable.
