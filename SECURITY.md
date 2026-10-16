# Politique de sécurité

Z01EXAM exécute du code envoyé par les utilisateurs. Une faille peut donc
toucher le serveur et tous les apprenants : merci de la signaler en privé.

## Signaler une faille

**N'ouvre pas d'issue publique et n'en parle pas sur le Discord.**

Utilise le signalement privé de GitHub :
[Security → Report a vulnerability](https://github.com/z01exam/z01exam/security/advisories/new).
Seul le mainteneur voit ton rapport.

Indique si possible :

- ce que la faille permet de faire ;
- les étapes pour la reproduire (requête, soumission, parcours concerné) ;
- la version affichée en bas à droite du site ou renvoyée par
  `GET /api/version`.

Le mainteneur accuse réception dès que possible, te tient informé de la
correction et te cite dans l'avis publié si tu le souhaites. Le projet est
bénévole : il n'y a pas de récompense.

## Ce qui est concerné

- Le code de ce dépôt : interface, catalogue, correcteur et serveur Go.
- L'isolation des soumissions : sortie du conteneur Docker, accès au réseau,
  contournement des limites de mémoire ou de processus.
- Le site [z01exam.online](https://z01exam.online).

Ne sont pas concernés : les images de test `ghcr.io/01-edu/*` et les contenus
de l'API Zone01, à signaler directement à Zone01.

## Tester sans nuire

- Teste en priorité sur ta copie locale (`make start`).
- Sur le site public, pas de déni de service, pas de charge massive, et pas
  d'accès aux données d'autres personnes.
- Arrête-toi dès que la faille est démontrée.
