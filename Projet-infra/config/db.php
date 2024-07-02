<?php

function BDD_Connect_LAN() {
    $host = 'localhost'; // Adresse IP du serveur de base de données
    $dbname = 'marya'; // Nom de la base de données
    $username = 'root'; // Nom d'utilisateur correct
    $password = ''; // Mot de passe correct

    $dsn = "mysql:host=$host;dbname=$dbname";
    $options = [
        PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION,
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
        PDO::ATTR_EMULATE_PREPARES   => false,
    ];

    try {
        $pdo = new PDO($dsn, $username, $password, $options);
        return $pdo;
    } catch (PDOException $e) {
        throw new PDOException("Database connection failed: " . $e->getMessage());
    }
}
