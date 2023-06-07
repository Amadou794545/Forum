package main

import (
	"fmt"
	"html/template"
	"net/http"

	"forum/Database"

	_ "github.com/mattn/go-sqlite3"
)

type LoginData struct {
	Username     string
	ErrorMessage string
}

func main() {
	//page
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/", Index)

	//css
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	//js

	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	//server
	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func Connexion(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/login.html"))
	loginData := LoginData{
		Username:     "",
		ErrorMessage: "",
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	println(username)
	println(password)
	if username != "" && password != "" { //a retirer quand premier check ok en js
		if Database.CheckLogin(username, password) {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {

			errorMessage := "Nom d'utilisateur ou Mot de passe invalide"

			loginData.Username = username
			loginData.ErrorMessage = errorMessage
		}
	}
	tmpl.Execute(w, loginData)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "template/inscription.html")
	} else if r.Method == "POST" {
		// Récupérer les données du formulaire d'inscription (username, email, password)
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		Database.AddUser(email, username, password, "test")
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

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}
