<?php
// Vérifier si le formulaire a été soumis
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // Récupérer les données du formulaire
    $title = $_POST["title"];
    $platform = $_POST["maison_edition"];
    $release_date = $_POST["release_date"];

    // Valider les données (vous pouvez ajouter des vérifications supplémentaires ici)

    // Connexion à la base de données (remplacez les valeurs par vos propres paramètres de connexion)
    $host = "localhost";
    $dbname = "votre_base_de_donnees";
    $username = "votre_nom_utilisateur";
    $password = "votre_mot_de_passe";

    try {
        $pdo = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);
        $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

        // Préparer la requête d'insertion
        $stmt = $pdo->prepare("INSERT INTO jeux_video (titre, maison_edition, date_sortie) VALUES (:titre, :maison_edition, :date_sortie)");

        // Exécuter la requête avec les données du formulaire
        $stmt->execute(array(
            ':titre' => $title,
            ':maison_edition' => $maison_edition,
            ':date_sortie' => $release_date
        ));

        // Rediriger l'utilisateur vers la page de liste des jeux
        header("Location: index.php");
        exit;
    } catch (PDOException $e) {
        // En cas d'erreur, afficher un message d'erreur
        echo "Erreur : " . $e->getMessage();
    }
}
?>
