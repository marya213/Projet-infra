# Nom de la classe principale
MAIN_CLASS=src/TextFieldTest.java

# Dossier pour les fichiers compilés
BIN_DIR := build

# Dossier pour les librairies
LIB_DIR := lib

# Séparateur de classpath (change this to : on Unix/Linux)
CP_SEP := ;

# Classpath incluant le dossier bin et les fichiers JAR dans lib
CLASSPATH := $(BIN_DIR)$(CP_SEP)$(LIB_DIR)/*

# Liste explicite des fichiers source
SOURCES := src/*.java

# Commande de build
build:
	@echo "Construction du projet..."
	@mkdir -p $(BIN_DIR)
	javac -cp "$(CLASSPATH)" -encoding UTF-8 -Xlint:deprecation -d $(BIN_DIR) $(SOURCES)

# Commande pour nettoyer le projet
clean:
	@echo "Nettoyage..."
	rm -rf $(BIN_DIR)

# Commande pour exécuter le programme
run: build
	@echo "Exécution du programme..."
	java -cp "$(CLASSPATH)" $(MAIN_CLASS)

# Option 'phony' pour indiquer que 'clean', 'run', 'build' ne sont pas des fichiers
.PHONY: build clean run
