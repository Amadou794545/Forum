FROM golang:1.20
# Copie du code source
WORKDIR /app
COPY . .
# Compilation de l'application Go
RUN go build -o run
# Exposition du port sur lequel le serveur web écoute
EXPOSE 3030
# Démarrage de l'application
CMD ["/app/run"]

# Supprimer tous les containers et toutes les images
    # docker system prune -a
# Build docker img
    # docker build -t forum .
# Check si docker img bien créée
    # docker images
# Run docker img (utiliser commande suivante à la place)
    # docker run -p 3030:3030 --name my-forum forum
# Run docker img sous forme de montage de volume (pour mise à jour auto du container)
	# docker run -p 3030:3030 -v C:/chemin/vers/dossier/Forum:/go/src/app --name my-forum forum
# Logs de notre container
    # docker logs my-forum