package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Serve les fichiers JavaScript
	fs := http.FileServer(http.Dir("java-script/"))
	http.Handle("/java-script/", http.StripPrefix("/java-script", fs))

	// Serve les fichiers CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Définit les routes pour les pages
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/", Index)
	http.HandleFunc("/upload", Upload)

	// Lance le serveur
	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// Récupérer le fichier image uploadé
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du fichier", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ouvrir un nouveau fichier pour écrire l'image
	f, err := os.OpenFile("assets/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Erreur lors de l'ouverture du fichier", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copier le contenu du fichier uploadé dans le fichier de destination
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Erreur lors de l'enregistrement du fichier", http.StatusInternalServerError)
		return
	}

	// L'image a été enregistrée avec succès
	// Vous pouvez effectuer d'autres actions, comme enregistrer le chemin de l'image dans une base de données, etc.

	// Répondre avec une confirmation
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("L'image a été enregistrée avec succès"))
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/inscription" {
		if r.Method == "GET" {
			http.ServeFile(w, r, "template/inscription.html")
		} else if r.Method == "POST" {
			// Récupérer les données du formulaire d'inscription (username, email, password)
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")
			// Vérifier si le mot de passe contient au moins un chiffre, une majuscule et un caractère spécial
			hasDigit := regexp.MustCompile(`\d`).MatchString(password)
			hasUpper := strings.ToUpper(password) != password
			hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)

			if !hasDigit || !hasUpper || !hasSpecial {
				errorMessage := "Le mot de passe doit contenir au moins un chiffre, une majuscule et un caractère spécial"
				http.Error(w, errorMessage, http.StatusBadRequest)
				http.ServeFile(w, r, "template/inscription.html")
				return
			} else {
				http.Redirect(w, r, "/", http.StatusFound)
			}

			// Afficher les données d'inscription dans la console
			fmt.Println("Nouvel utilisateur enregistré :")
			fmt.Println("Username :", username)
			fmt.Println("Email :", email)
			fmt.Println("Password :", password)

			// Enregistrer l'utilisateur dans la base de données

			// Rediriger vers la page de connexion ou afficher un message de succès
		}

	}
}

func Connexion(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" {
		if r.Method == "GET" {
			http.ServeFile(w, r, "template/login.html")
		} else if r.Method == "POST" {
			// Récupérer les données du formulaire de connexion (username, password)
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Vérifier les informations de connexion
			if username == "john" && password == "secret" {
				// Informations de connexion valides
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				// Informations de connexion invalides
				http.ServeFile(w, r, "template/login.html")
			}
		}
	}
}
