# Instructions du dépôt

## Git et livraison

- `main` est la branche d'intégration : les pushes directs y sont autorisés,
  sans branche dédiée ni pull/merge request.
- Ne pas demander une confirmation redondante avant un commit ou un push vers
  `main`, sauf si une règle obligatoire de la plateforme l'impose.
- Le déploiement de production est manuel. Ne pas le déclencher sans demande
  explicite de l'utilisateur.

## Revue

- Ne jamais lancer de revue Claude ni `ai-review-loop` dans ce dépôt.
- Lorsqu'une revue indépendante est nécessaire, elle est réalisée par Sol.

## Qualité

- Garder les changements ciblés et vérifier le comportement concerné avant la
  livraison.
- Ne pas modifier ou supprimer des tests uniquement pour obtenir un passage au
  vert.
