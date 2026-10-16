# Contribuer à Z01EXAM

Bienvenue ! Ce guide suppose que tu débutes avec Git : tout se fait avec des
boutons, sans terminal.

## Ce qu'il te faut

- Un compte GitHub, ajouté au projet par le mainteneur, avec :
  - la **double authentification** activée
    ([Settings → Password and authentication](https://github.com/settings/security)),
    obligatoire pour rejoindre l'organisation ;
  - ton **e-mail privé** : coche *Keep my email addresses private* et *Block
    command line pushes that expose my email* dans
    [Settings → Emails](https://github.com/settings/emails). Sinon, ton
    adresse personnelle apparaît publiquement dans chaque commit.
- [GitHub Desktop](https://desktop.github.com/) pour récupérer et envoyer le code.
- Un éditeur : VS Code ou WebStorm.
- Pour lancer le site en local : Go 1.22 ou plus et Docker (voir le README).

## Le parcours en 6 étapes

### 1. Choisir un ticket

Ouvre l'onglet **Projects** du dépôt, puis le tableau Z01EXAM. Choisis une
carte de la colonne **À faire**. Pour commencer, prends un ticket étiqueté
`good first issue`.

### 2. Le prendre

Ouvre la carte, puis clique sur **Assign yourself** dans la colonne de droite.
Ton avatar apparaît sur la carte : tout le monde sait que tu t'en occupes.
Fais glisser la carte dans **En cours**.

Tu ne peux pas t'assigner ? Écris « Je prends » en commentaire, le mainteneur
t'assignera.

### 3. Créer ta branche depuis le ticket

1. Dans le ticket, colonne de droite, section **Development**, clique sur
   **Create a branch**, puis valide. La branche est créée et liée au ticket :
   tout le monde verra tes commits depuis le ticket.
2. Dans GitHub Desktop, **File → Clone repository** la première fois, puis
   choisis z01exam.
3. **Fetch origin**, puis **Current branch** : sélectionne la branche que tu
   viens de créer.

### 4. Coder et tester

Ouvre le dossier dans ton éditeur, fais ta modification, puis lance le site
avec `make start` et vérifie dans `http://localhost:8080`.

Garde la modification petite : un ticket, un sujet.

### 5. Montrer ton avancement dès le début

Dès ton premier commit, ouvre une PR en **brouillon** : tout le monde voit où
tu en es, et le mainteneur peut t'aider en cours de route.

Dans GitHub Desktop :

1. Écris un résumé en bas à gauche, puis clique sur **Commit**.
2. Clique sur **Push origin**.
3. Clique sur **Create Pull Request** : GitHub s'ouvre.
4. Remplis le formulaire et garde la ligne `Closes #42` avec le numéro de ton
   ticket : il se fermera tout seul quand la PR sera acceptée.
5. Clique sur la flèche à côté du bouton vert, choisis **Create draft pull
   request**, puis valide.

Ensuite, à chaque fois que tu avances : **Commit**, puis **Push origin**. La PR
se met à jour toute seule. Pousse au moins une fois par séance de travail.

Quand tu as fini, clique sur **Ready for review** en bas de la PR, et fais
glisser la carte dans **En relecture**.

### 6. La relecture

- La CI (les tests automatiques) se lance. Une croix rouge ? Ouvre le détail,
  corrige, puis refais Commit et Push : la PR se met à jour.
- Le mainteneur relit et peut laisser des commentaires. Corrige de la même
  façon, puis réponds au commentaire.
- Une fois approuvée, la PR est fusionnée par le mainteneur et la carte passe
  dans **Fait**.

## Les règles

- On ne pousse jamais directement sur `main` : tout passe par une PR.
- Une PR = un seul sujet, testé.
- Toute PR est relue et approuvée par le mainteneur avant d'être fusionnée.
- La mise en production est faite par le mainteneur, jamais automatiquement.
- Un ticket pris sans nouvelles pendant deux semaines peut être réattribué.
- Tout le monde respecte le [code de conduite](CODE_OF_CONDUCT.md).

## Où modifier quoi

| Je veux changer… | Je regarde… |
| --- | --- |
| L'interface | `index.html`, `app.js`, `style.css` |
| Un énoncé Go ou JS | `exam-instructions/`, `Js_exam_instructions/` |
| Une leçon du mode Learning | `learning-content/` |
| Les tests et la moulinette | `backend/runner*.go`, `backend/rust.go`, `backend/java.go` |

Les rôles de chaque brique sont décrits dans
[docs/architecture-services.md](docs/architecture-services.md).

## Une question ?

Passe sur le [Discord de la communauté](https://discord.gg/cP2Dn5DF6H).
