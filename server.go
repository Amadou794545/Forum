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
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/inscription", handlerInscription)
	http.HandleFunc("/login", handlerConnexion)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if CheckSessionCookie(r) {
			cookie, err := r.Cookie("session")
			if err != nil {
				// Handle the error if needed
				fmt.Println("Error retrieving session cookie:", err)
				return
			}

			userID := cookie.Value
			username, err := Database.GetUserUsername(userID)
			if err != nil {
				// Handle the error if needed
				fmt.Println("Error retrieving username:", err)
				return
			}

			UpdateSessionExpiration(w, r) // Reset la date de péremption du cookie
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
	}
}

func handlerConnexion(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" {
		if r.Method == "GET" {
			http.ServeFile(w, r, "template/login.html")
		} else if r.Method == "POST" {
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
					userID, err := Database.GetUserID(username)
					if err != nil {
						fmt.Println("Error:", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
						return
					}
					HandlerCookie(w, r, userID) // Ajout du cookie de session
					http.Redirect(w, r, "/", http.StatusFound)
					return
				} else {
					errorMessage := "Nom d'utilisateur ou Mot de passe invalide"
					loginData.Username = username
					loginData.ErrorMessage = errorMessage
				}
			}
			tmpl.Execute(w, loginData)
		}
	}
}
