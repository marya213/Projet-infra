# Utilise l'image officielle de Go pour construire l'application
FROM golang:1.18-alpine

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier les fichiers de module Go et télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le contenu actuel dans le répertoire de travail dans le conteneur
COPY . .

# Construire l'application
RUN go build -o main .

# Démarrer l'application
CMD ["./main"]
