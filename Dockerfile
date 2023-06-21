FROM golang:1.20

# Copie du code source
WORKDIR /go/src/app
COPY . .

# Compilation de l'application Go
RUN go build -o /go/bin/app

# Exposition du port sur lequel le serveur web écoute
EXPOSE 3030

# Démarrage de l'application
CMD ["/go/bin/app"]


# Build docker img
    # docker build -t forum .
# Check si l'img a été crée
    # docker images
# Run docker img
    # docker run -p 3030:3030 --name my-forum forum
# Run docker img sous forme de montage de volume (pour mise à jour auto du container)
	# docker run -p 3030:3030 -v C:/chemin/vers/dossier/du/projet:/go/src/app --name my-forum forum
# Logs de notre container
    # docker logs my-forum
