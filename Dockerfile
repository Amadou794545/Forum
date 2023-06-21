FROM golang:1.20

# Installation du client MariaDB
RUN apt-get update && apt-get install -y mariadb-client

# Copie du code source
WORKDIR /go/src/app
COPY . .

# Compilation de l'application Go
RUN go build -o /go/bin/app

# Copie des contenus statiques
COPY CSS /go/bin/CSS
COPY java-script /go/bin/java-script
COPY Pictures /go/bin/Pictures
COPY template /go/bin/template

# Exposition du port sur lequel le serveur web écoute
EXPOSE 3030

# Démarrage de l'application
CMD ["/go/bin/app"]


# Lancer le serveur via docker
    # docker run -p 3030:3030 --name my-forum forum
# Lancer le serveur via docker sous forme de montage de volume (pour mise à jour auto du container)
	# docker run -p 3030:3030 -v C:/Users/fanny/Desktop/Forum:/go/src/app --name my-forum forum
# Logs de notre container
    # docker logs my-forum
