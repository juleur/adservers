# Adservers
Adservers est constitué de 2 serveurs web permettant de gérer:
* - les demandes de publicités de la part d'un éditeur pour un emplacement publicité donné
* - l'affichage des publicités réussis avec possibilité de recevoir la somme des prix des pubs affichées 

## Exemples
POST requête avec l'ID placement 
```
http://localhost:2323/ad?placement=f64551fcd6f07823cb87971cfb914464
```
accompagnée du body
```
{
 "country":"FRA",
 "device":"DESKTOP"
}
```
GET requête permettant de récupérer la somme totale des pubs affichées
```
http://localhost:3030/sum
```
GET requête permettant de récupérer la somme totale par campagnes des pubs affichées via l'ID placement
```
http://localhost:3030/sum_placement?placement=f64551fcd6f07823cb87971cfb914464
```
