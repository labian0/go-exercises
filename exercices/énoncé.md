# GOSTERCLASS - exercices 

## 1. JEU DEVIN
c'est un jeu devin classique, vous devez trouver le bon nombre pour gagner

démarche : 
générer un nombre aléatoire -> envoyer une requête api vers l'api -> traiter la réponse et agir en fonction
le serveur renvoie du json sous cette forme si tout va bien:
{"message": 1}
la valeur est 1 si votre nombre est supérieur au nombre à deviner, -1 si inférieur, 0 si vous avez trouvé!
