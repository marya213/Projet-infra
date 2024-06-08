## Ce projet se lance sur visual ou  intellij IDEA en utilisant Makefile 1 make build 2 make run 3 make clean pour supprimer les fichiers générés.

 
# Task Management Application
Task Management Application :
 est une application de gestion de tâches basée sur Java, qui permet à l'utilisateur de créer, afficher, éditer et supprimer des tâches à l'aide d'une interface graphique simple. L'application utilise Gson pour la sérialisation et la désérialisation des données au format JSON.
// Ce projet se lance sur intellij IDEA

# Dépendances

Ce projet utilise les librairies suivantes :
Download gson JAR 2.8.2 with all dependencies: https://jar-download.com/artifacts/com.google.code.gson/gson/2.8.2/source-code.

## Gestions des Dépendances
Créez un dossier lib (ou un autre nom de votre choix) à la racine de votre projet.

Déplacez le fichier JAR téléchargé (gson-2.8.2-with-dependencies.jar) dans le dossier lib.

Dans votre IDE (par exemple, IntelliJ IDEA, Eclipse, etc.), ouvrez les paramètres du projet ou les propriétés du module, et trouvez la section pour ajouter des bibliothèques ou des dépendances externes.

Ajoutez le fichier JAR Gson que vous avez déplacé dans le dossier lib comme une bibliothèque externe dans votre projet.

 ## Fonctionnalités:
L'application TextFieldTest offre les fonctionnalités suivantes:
1-Charger et enregistrer des tâches à partir de fichiers JSON.
2-Ajouter une nouvelle tâche avec un nom, une description, une priorité et une catégorie.
3-Afficher les détails d'une tâche existante.
4-Éditer les détails d'une tâche existante.
5-Supprimer une tâche existante.

## Utilisation de l'Application :
1-Lorsque vous exécutez l'application TextFieldTest dans IntelliJ IDEA, vous serez accueilli par une interface graphique intuitive comprenant les éléments suivants 

2-Icône d'ajout de tâche : Cette icône représente un signe plus (+) et permet à l'utilisateur d'ajouter une nouvelle tâche à la liste

3-En cliquant sur cette icône, une boîte de dialogue s'ouvre où l'utilisateur peut saisir les détails de la nouvelle tâche, tels que le nom, la description, la priorité et la catégorie et il peut ajouter une nouvel catgory en faisant add new catgory et en l'enregistrant.

4-Icône d'ouverture de tâche : Cette icône représente un dossier ouvert et permet à l'utilisateur d'ouvrir une tâche existante pour afficher ses détails. En cliquant sur cette icône, une boîte de dialogue s'ouvre où l'utilisateur peut saisir la description de la tâche à ouvrir. Si une tâche correspondante est trouvée, ses détails sont affichés dans une boîte de dialogue séparée.

5-Icône d'édition de tâche : Cette icône représente un crayon et permet à l'utilisateur d'éditer les détails d'une tâche existante. En cliquant sur cette icône, une boîte de dialogue s'ouvre où l'utilisateur peut saisir la description de la tâche à éditer. Si une tâche correspondante est trouvée, ses détails sont affichés dans une boîte de dialogue pré-remplie, où l'utilisateur peut effectuer les modifications nécessaires.

6-Icône de suppression de tâche : Cette icône représente une corbeille et permet à l'utilisateur de supprimer une tâche existante. En cliquant sur cette icône, une boîte de dialogue s'ouvre où l'utilisateur peut saisir la description de la tâche à supprimer. Si une tâche correspondante est trouvée, une confirmation est demandée avant de supprimer définitivement la tâche.

7-Liste déroulante de fichiers de tâches : Cette liste déroulante permet à l'utilisateur de sélectionner un fichier de tâches parmi une liste prédéfinie. Chaque fichier représente un ensemble de tâches et est associé à un nom spécifique. Lorsque l'utilisateur sélectionne un fichier dans la liste, les tâches correspondantes sont chargées et affichées dans l'interface principale.

8-Zone de texte pour entrer le nom de la tâche : Cette zone de texte permet à l'utilisateur de saisir le nom de la tâche lorsqu'il souhaite ajouter une nouvelle tâche à la liste. Le nom de la tâche est un champ obligatoire et doit être fourni avant de pouvoir ajouter la tâche.

Merci d'avoir utilisé l'application Task Manageme nt Application. Nous espérons que vous apprécierez la simplicité et l'utilité de cette application pour gérer vos tâches quotidiennes.

