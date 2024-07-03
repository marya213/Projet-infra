<?php
require '../config/db.php';

try {
    // Connexion à la base de données
    $pdo = BDD_Connect_LAN();

    // Suppression d'un jeu si l'ID est fourni dans l'URL
    if (isset($_GET['id']) && is_numeric($_GET['id'])) {
        $id = $_GET['id'];

        // Préparation et exécution de la requête SQL pour supprimer l'enregistrement
        $stmt = $pdo->prepare('DELETE FROM jeux WHERE id = ?');
        if ($stmt->execute([$id])) {
            // Redirection vers index.php après la suppression
            header('Location: index.php');
            exit;
        } else {
            // Gestion de l'erreur en cas d'échec de la suppression
            echo "Error: Unable to delete the record.";
        }
    }

    // Récupération des jeux depuis la base de données
    $sql = "SELECT id, title, image, maison_d_edition, evaluation, date_evaluation FROM jeux";
    $stmt = $pdo->query($sql);
    $games = $stmt->fetchAll(PDO::FETCH_ASSOC);
} catch (PDOException $e) {
    echo 'Connection failed: ' . $e->getMessage();
    exit;
}
?>
 <!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Video Games List</title>
    <link rel="stylesheet" type="text/css" href="../assets/css/styles.css">
    <style>
        .game-item {
            display: flex;
            margin-bottom: 20px;
        }

        .game-item img {
            width: 150px; /* Ajustez la largeur selon vos besoins */
            margin-right: 20px; /* Espacement entre l'image et les détails du jeu */
        }

        .game-details {
            flex: 1;
        }
    </style>
</head>
<body>
    <header>
        <div class="container">
            <h1>Video Games List</h1>
            <nav>
                <ul>
                    <li><a href="index.php">Accueil</a></li>
                    <li><a href="add_game.php">Ajouter un jeu</a></li>
                </ul>
            </nav>
        </div>
    </header>

    <div class="container">
        <h2>Liste des Jeux</h2>
        <ul class="game-list">
            <?php foreach ($games as $game): ?>
                <li class="game-item">
                    <div class="game-details">
                        <img src="<?php echo htmlspecialchars($game['image']); ?>" alt="<?php echo htmlspecialchars($game['title']); ?>">
                        <h3><?php echo htmlspecialchars($game['title']); ?></h3>
                        <p><strong>Maison d'édition :</strong> <?php echo htmlspecialchars($game['maison_d_edition']); ?></p>
                        <p><strong>Évaluation :</strong> <?php echo htmlspecialchars($game['evaluation']); ?></p>
                        <p><strong>Date d'évaluation :</strong> <?php echo htmlspecialchars($game['date_evaluation']); ?></p>
                        <div class="action-buttons">
                            <a href="edit_game.php?id=<?php echo $game['id']; ?>">Modifier</a>
                            <a href="index.php?id=<?php echo $game['id']; ?>" onclick="return confirm('Voulez-vous vraiment supprimer ce jeu?')">Supprimer</a>
                        </div>
                    </div>
                </li>
            <?php endforeach; ?>
        </ul>
    </div>
</body>
</html>
