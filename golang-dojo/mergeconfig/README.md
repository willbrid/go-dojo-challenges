# L'énoncé

```go
type Config struct {
    Name     string
    Tags    []string
    Limits map[string]int
}

func MergeConfig(base, surcharge Config) Config
```

### Contraintes

1. Elle produit une configuration où les valeurs de `surcharge` écrasent celles de `base` **quand elles sont renseignées**.
2. Contrainte impérative : ni `base` ni `surcharge` ne doivent être modifiés, et le résultat ne doit partager **aucune** structure mutable avec eux.
3. Écris aussi, en commentaire au-dessus, le contrat que ta fonction garantit.
