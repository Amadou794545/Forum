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

type InscriptionData struct {
	Username     string
	Email        string
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
	tmpl := template.Must(template.ParseFiles("./template/inscription.html"))
	InscriptionData := InscriptionData{
		Username:     "",
		Email:        "",
		ErrorMessage: "",
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	errorMessage := ""

	if username == "" || email == "" || password == "" {
		tmpl.Execute(w, InscriptionData)
	} else {

		if Database.CheckUsername(username) {
			errorMessage = "Username deja utilisé"
		}
		if Database.CheckEmail(email) {
			errorMessage += " Email deja utilisé"
		}

		if !Database.CheckUsername(username) && !Database.CheckEmail(email) {
			Database.AddUser(email, username, password, "test")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			InscriptionData.Username = username
			InscriptionData.Email = email
			InscriptionData.ErrorMessage = errorMessage
		}

		tmpl.Execute(w, InscriptionData)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}
