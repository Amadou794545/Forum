package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"forum/Database"

	_ "github.com/mattn/go-sqlite3"
)

/*func main() {
	//page
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/", Index)

	//css
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	//server
	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}*/

func Connexion(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/login" {
		if r.Method == "GET" {
			http.ServeFile(w, r, "template/login.html")
		} else if r.Method == "POST" {
			// Récupérer les données du formulaire de connexion (username, password)
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Vérifier les informations de connexion

			if Database.CheckLogin(username, password) {
				println("ok log")
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				println("ko")
				http.ServeFile(w, r, "template/login.html")
			}
		}
	}
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
				Database.AddUser(email, username, password, "test")
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

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}
