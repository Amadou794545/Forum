package main

import (
	"fmt"
	"forum/Database"
	"net/http"
	"regexp"
	"strings"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if CheckSessionCookie(r) {
			cookie, err := r.Cookie("session")
			if err != nil {
				fmt.Println("Error retrieving session cookie:", err)
				return
			}

			userID := cookie.Value
			username, err := Database.GetUserUsername(userID)
			if err != nil {
				fmt.Println("Error retrieving username:", err)
				return
			}

			fmt.Println("Bienvenue", username)
		}

		http.ServeFile(w, r, "template/index.html")
	}
}

func handlerInscription(w http.ResponseWriter, r *http.Request) {
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

func handlerConnexion(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" {
		if r.Method == "GET" {
			http.ServeFile(w, r, "template/login.html")
		} else if r.Method == "POST" {

			// Vérification des informations de connexion
			username := r.FormValue("username")
			password := r.FormValue("password")

			if Database.CheckLogin(username, password) {
				fmt.Println("ok log")
				userID, err := Database.GetUserID(username)
				if err != nil {
					// Handle the error appropriately
					fmt.Println("Error:", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				HandlerCookie(w, r, userID)   // Ajout du cookie de session
				UpdateSessionExpiration(w, r) // Mise à jour de sa durée
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			fmt.Println("ko")
			http.ServeFile(w, r, "template/login.html")
		}
	}
}
