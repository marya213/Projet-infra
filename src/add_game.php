<?php
if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    $db_file = '../config/db.php';
    if (file_exists($db_file) && is_readable($db_file)) {
        require $db_file;

        // Connexion à la base de données
        try {
            $pdo = BDD_Connect_LAN();
        } catch (PDOException $e) {
            echo "Erreur de connexion à la base de données: " . $e->getMessage();
            exit();
        }

        $title = htmlspecialchars(trim($_POST['title']));
        $brand = htmlspecialchars(trim($_POST['maison_d_edition']));
        $evaluation = filter_var($_POST['evaluation'], FILTER_VALIDATE_INT);
        $date_evaluation = date('Y-m-d');

        $image = $_FILES['image']['name'];
        $imageTmpName = $_FILES['image']['tmp_name'];
        $uploadDir = 'uploads/';

        // Générer un nom de fichier unique pour éviter les conflits
        $uploadFile = $uploadDir . uniqid() . '_' . basename($image);

        // Vérifier si le fichier image est une image réelle
        $check = getimagesize($imageTmpName);
        if ($check !== false) {
            if (move_uploaded_file($imageTmpName, $uploadFile)) {
                $stmt = $pdo->prepare('INSERT INTO jeux (title, maison_d_edition, image, evaluation, date_evaluation) VALUES (?, ?, ?, ?, ?)');
                if ($stmt->execute([$title, $brand, $uploadFile, $evaluation, $date_evaluation])) {
                    header('Location: index.php');
                    exit();
                } else {
                    echo "Erreur: Impossible d'exécuter la requête.";
                }
            } else {
                echo "Erreur: Impossible de télécharger l'image.";
            }
        } else {
            echo "Le fichier n'est pas une image.";
        }
    } else {
        echo "Database configuration file is missing or not readable.";
    }
}
?>
<!DOCTYPE html>
<html>
<head>
    <title>Ajouter un Jeu</title>
    <link rel="stylesheet" type="text/css" href="../assets/css/styles.css">
</head>
<body>
    <h1>Ajouter un Jeu</h1>
    <form method="post" enctype="multipart/form-data">
        <label>Titre: <input type="text" name="title" required></label><br>
        <label>Maison edition: <input type="text" name="maison_d_edition" required></label><br>
        <label>Image: <input type="file" name="image" accept="image/*" required></label><br>
        <label>Évaluation: <input type="number" name="evaluation" required></label><br>
        <button type="submit">Ajouter</button>
    </form>
</body>
</html>
