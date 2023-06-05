package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/inscription", Inscription)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", LoginPageHandler)

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
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

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, error := r.Cookie("session")
		if error == nil {
			(http.Redirect(w, r, "/", http.StatusSeeOther))
			return
		}
		t, _ := template.ParseFiles("./template/LoginPage.html")
		t.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		if r.FormValue("register") != ("") {
			if r.FormValue("username") != ("") && r.FormValue("password") != ("") {
				temp := GetPassword()
				if _, ok := temp[r.FormValue("username")]; !ok {
					temp[r.FormValue("username")] = r.FormValue("password")
					SetPassword(temp)
					(http.Redirect(w, r, "/", http.StatusSeeOther))
				} else {
					type PageInfos struct {
						ErrorMessage string
					}
					t, _ := template.ParseFiles("./template/LoginPage.html")
					t.Execute(w, PageInfos{"This username is already in use!"})
				}
			}
		}
		if r.FormValue("login") != ("") {
			temp := GetPassword()
			if temp[r.FormValue("username")] == r.FormValue("password") {
				generateSession(w, r.FormValue("username"))
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				type PageInfos struct {
					ErrorMessage string
				}
				t, _ := template.ParseFiles("./template/LoginPage.html")
				t.Execute(w, PageInfos{"Invalid password or username."})
			}
		}
	}
}

func generateSession(w http.ResponseWriter, username string) {
	cookie := http.Cookie{Name: "session", Value: username}
	http.SetCookie(w, &cookie)
}

func SetPassword(data map[string]string) {
	value, error := json.Marshal(data)
	if error == nil {
		os.WriteFile("./package.json", value, 0644)
	}
}

func GetPassword() map[string]string {
	file, error := os.ReadFile("./package.json")
	if error == nil {
		content := map[string]string{}
		json.Unmarshal(file, &content)
		return content
	}
	return map[string]string{}
}
