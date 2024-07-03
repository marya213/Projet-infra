<?php
require '../config/db.php';

// Vérifiez que l'ID est présent et valide
$id = filter_input(INPUT_GET, 'id', FILTER_VALIDATE_INT);
if ($id === false) {
    die('ID invalide.');
}

try {
    // Connexion à la base de données
    $pdo = BDD_Connect_LAN();

    // Préparez et exécutez la requête pour obtenir les données du jeu
    $stmt = $pdo->prepare('SELECT * FROM jeux WHERE id = ?');
    $stmt->execute([$id]);
    $game = $stmt->fetch();

    if (!$game) {
        die('Jeu non trouvé.');
    }

    if ($_SERVER['REQUEST_METHOD'] == 'POST') {
        $title = filter_input(INPUT_POST, 'title', FILTER_SANITIZE_STRING);
        $brand = filter_input(INPUT_POST, 'brand', FILTER_SANITIZE_STRING);
        $evaluation = filter_input(INPUT_POST, 'evaluation', FILTER_VALIDATE_INT);

        $image = $game['image']; // valeur par défaut
        if (isset($_FILES['image']) && $_FILES['image']['error'] == UPLOAD_ERR_OK) {
            $imageTmpPath = $_FILES['image']['tmp_name'];
            $imageName = basename($_FILES['image']['name']);
            $imageSize = $_FILES['image']['size'];
            $imageType = $_FILES['image']['type'];
            
            // Liste des types de fichiers acceptés
            $allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];

            if (in_array($imageType, $allowedTypes) && $imageSize < 2000000) { // Limite de taille à 2MB
                $imagePath = "uploads/$imageName";
                if (move_uploaded_file($imageTmpPath, $imagePath)) {
                    $image = $imageName;
                } else {
                    die('Erreur lors du téléchargement de l\'image.');
                }
            } else {
                die('Type de fichier ou taille non valide.');
            }
        }

        $date_evaluation = date('Y-m-d');

        // Mise à jour des données du jeu dans la base de données
        $stmt = $pdo->prepare('UPDATE jeux SET title = ?, brand = ?, image = ?, evaluation = ?, date_evaluation = ? WHERE id = ?');
        $stmt->execute([$title, $brand, $image, $evaluation, $date_evaluation, $id]);

        // Redirection vers la page principale
        header('Location: index.php');
        exit;
    }
} catch (PDOException $e) {
    echo 'Connection failed: ' . $e->getMessage();
    exit;
}
?>
<!DOCTYPE html>
<html>
<head>
    <title>Edit Game</title>
    <link rel="stylesheet" type="text/css" href="../assets/css/styles.css">
</head>
<body>
    <h1>Edit Game</h1>
    <form method="post" enctype="multipart/form-data">
        <label>Title: <input type="text" name="title" value="<?php echo htmlspecialchars($game['title']); ?>"></label><br>
        <label>Brand: <input type="text" name="brand" value="<?php echo htmlspecialchars($game['maison_d_edition']); ?>"></label><br>
        <label>Image: <input type="file" name="image"></label><br>
        <label>Evaluation: <input type="number" name="evaluation" value="<?php echo htmlspecialchars($game['evaluation']); ?>"></label><br>
        <button type="submit">Save</button>
    </form>
</body>
</html>
