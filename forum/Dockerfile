# Utiliser une image de base Go
FROM golang:1.16-alpine

# Répertoire de travail
WORKDIR /app

# Copier les fichiers du projet
COPY . .

# Installer les dépendances
RUN go mod download

# Construire l'application
RUN go build -o forum

# Exposer le port
EXPOSE 8080

# Démarrer l'application
CMD ["./forum"]
