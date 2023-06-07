package main

import (
	"fmt"
	"net/http"
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

func Inscription(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "template/inscription.html")
	} else if r.Method == "POST" {
		// Récupérer les données du formulaire d'inscription (username, email, password)
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		http.Redirect(w, r, "/", http.StatusFound)

		// Afficher les données d'inscription dans la console
		fmt.Println("Nouvel utilisateur enregistré :")
		fmt.Println("Username :", username)
		fmt.Println("Email :", email)
		fmt.Println("Password :", password)

		// Enregistrer l'utilisateur dans la base de données

		// Rediriger vers la page de connexion ou afficher un message de succès
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
