# Projet de Gestion de Jeux Vidéo

## But du Projet

Ce projet a pour objectif de développer une application de gestion de jeux vidéo permettant d'ajouter, de modifier, de consulter et de supprimer des jeux vidéo.

## Cadre de Développement

- **École** : Ynov campus
- **Niveau d'études** : B1 informatique
- **Groupe** : Réalisé en groupe de 2 personnes


## Recherche et Tutoriels

Pour développer ce projet, j'ai effectué des recherches approfondies sur Internet et suivi plusieurs tutoriels sur YouTube pour comprendre les concepts de base et avancés de PHP, MySQL et CSS. Voici quelques-unes des ressources que j'ai utilisées :

- [Tutoriels PHP sur YouTube](https://www.youtube.com/results?search_query=php+tutorial)
- [Documentation officielle de PHP](https://www.php.net/docs.php)
- [Tutoriels MySQL sur YouTube](https://www.youtube.com/results?search_query=mysql+tutorial)
- [Documentation officielle de MySQL](https://dev.mysql.com/doc/)
- [Tutoriels CSS sur YouTube](https://www.youtube.com/results?search_query=css+tutorial)
- [Documentation officielle de CSS](https://developer.mozilla.org/en-US/docs/Web/CSS)


## Le système de gestion de jeux vidéo utilise PHP et MySQL pour gérer les données. 
L'interface utilisateur est stylisée avec CSS pour offrir une expérience de style gaming.

## Configuration

1. Clonez ce dépôt sur votre machine locale.

    git clone https://github.com/votre_nom_d_utilisateur/votre_repertoire.git
    
2. Importez le fichier `db.sql` dans votre base de données MySQL.
3. Configurez les informations de connexion à la base de données dans `config/db.php`.

    ```php
    <?php
    define('DB_HOST', 'votre_hôte');
    define('DB_NAME', 'votre_base_de_donnees');
    define('DB_USER', 'votre_nom_d_utilisateur');
    define('DB_PASS', 'votre_mot_de_passe');
    ?>
    ```

4. Ouvrez votre navigateur et accédez à `index.php` pour commencer à utiliser le système de gestion de jeux vidéo.

## Utilisation

### Ajouter un jeu

1. Cliquez sur "Add New Game" dans le menu de navigation.
2. Remplissez le formulaire avec les informations du jeu.
3. Cliquez sur "Submit" pour ajouter le jeu à la base de données.

### Modifier un jeu

1. Cliquez sur "Edit" à côté du jeu que vous souhaitez modifier.
2. Modifiez les informations dans le formulaire.
3. Cliquez sur "Save" pour enregistrer les modifications.

### Supprimer un jeu

1. Cliquez sur "Delete" à côté du jeu que vous souhaitez supprimer.
2. Confirmez la suppression.


## Fonctionnalités

- Ajouter un nouveau jeu
- Modifier les informations d'un jeu
- Supprimer un jeu

## Stack Technique

- **Langage** : PHP
- **Framework** : Aucun (utilisation de PHP natif)
- **Base de Données** : MySQL

