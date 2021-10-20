# TP 1 Info002 - Cryptologie
## Compromis temps / mémoire, les tables arc-en-ciel

> Hugo Hersemeule & Tom Kubasik

> Sujet du TP : https://www.lama.univ-savoie.fr/pagesmembres/hyvernat/Enseignement/2122/info910/tp1.html


Pour le moment, nous sommes rendus à la question 3, avec size_min == size_max.

## Question 2 : 
go run . CALCUL_N ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz 4 7

## Question 3 : Indices et textes clairs
go run . INDEX ABCDEFGHIJKLMNOPQRSTUVWXYZ 12345 3 3
-> RFV

## Question 5 : programmez la fonction h2i
go run . Q5 

## Question 6 : programmez i2i pour nouvelle_chaine(index, largeur)
go run . Q6 MD5 ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz 4 5   

## Question 7 : 

> En quoi est-ce que l'ajout du paramètre t dans la fonction h2i permet d'augmenter la couverture de la table ?

Si on trouve que l'empreinte de deux textes clairs sont identique pour un t donné, alors on peut recommencer avec une autre valeur de t pour trouver un résultat différent.

## Question 8 :

go run . Q8 MD5 abcdefghijklmnopqrstuvwxyz 5 5 200 100

Args après l'alphabet : size_min size_max width height

## Question 9 :

go run . Q9 MD5 abcdefghijklmnopqrstuvwxyz 4 5 200 50
(on remet les arguments pour éviter les erreurs par rapport à notre fonction main)