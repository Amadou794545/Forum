package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"forum/Database"
	"forum/cookies"

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
	//testPost()

	http.HandleFunc("/api/posts", GetPostsAPI)

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/inscription", handlerInscription)
	http.HandleFunc("/inscriptionPicture", handlerInscriptionPicture)
	http.HandleFunc("/inscriptionDada", handlerInscriptionDada)
	http.HandleFunc("/login", handlerConnexion)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures", http.FileServer(http.Dir("Pictures"))))
	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func testPost() {
	for i := 0; i < 150; i++ {
		if i < 10 {
			Database.AddPost("test"+strconv.Itoa(i), "", "test", 2, 1)
		}
		if i < 25 && i > 9 {
			Database.AddPost("test2"+strconv.Itoa(i), "", "test", 2, 2)
		}
		if i > 24 && i < 50 {
			Database.AddPost("test3"+strconv.Itoa(i), "", "test", 2, 3)
		}
		if i > 49 && i < 75 {
			Database.AddPost("test4"+strconv.Itoa(i), "", "test", 2, 4)
		}
		if i > 74 && i < 100 {
			Database.AddPost("test5"+strconv.Itoa(i), "", "test", 2, 5)
		}
		if i > 99 && i < 125 {
			Database.AddPost("test6"+strconv.Itoa(i), "", "test", 2, 6)
		}
		if i > 124 && i < 150 {
			Database.AddPost("test7"+strconv.Itoa(i), "", "test", 2, 7)
		}
	}
}

func GetPostsAPI(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	// Convert page and limit values to integers
	pageNum, _ := strconv.Atoi(page)
	limitNum, _ := strconv.Atoi(limit)

	// Calculate the offset based on the requested page and limit
	offset := (pageNum - 1) * limitNum

	// Fetch posts from the database with pagination
	posts, err := Database.GetPosts(offset, limitNum)
	if err != nil {
		log.Println("Erreur lors de la récupération des posts :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(posts)
	if err != nil {
		log.Println("Erreur lors de la conversion en JSON :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if cookies.CheckSessionCookie(r) {
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

			cookies.UpdateSessionExpiration(w, r) // Reset la date de péremption du cookie
			fmt.Println("Bienvenue", username)
		}

		http.ServeFile(w, r, "template/index.html")
	}
}

func handlerInscription(w http.ResponseWriter, r *http.Request) {
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
			cookie := http.Cookie{
				Name:  "username",
				Value: username,
			}
			http.SetCookie(w, &cookie)
			imgPath := "pictures/Profil/anonyme.jpg"
			Database.AddUser(email, username, password, imgPath)
			http.Redirect(w, r, "/inscriptionPicture", http.StatusFound)
		} else {
			InscriptionData.Username = username
			InscriptionData.Email = email
			InscriptionData.ErrorMessage = errorMessage
		}

		tmpl.Execute(w, InscriptionData)
	}
}

func handlerInscriptionPicture(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		imgPath := r.Form.Get("selectedPicture")
		filename := "Pictures/Profil/" + path.Base(imgPath)

		cookie, err := r.Cookie("username")
		if err == nil {
			username := cookie.Value
			userID, err := Database.GetUserID(username)
			if err != nil {
				fmt.Println("Error getting user ID:", err)
				http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
				return
			}

			Database.UpdateImgProfile(filename, userID)
			jsonResponse := struct {
				PhotoProfil string `json:"PhotoProfil"`
			}{
				PhotoProfil: filename,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(jsonResponse)
			return
		}

		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Serve the HTML file for GET request
	http.ServeFile(w, r, "template/inscription_picture.html")
}

func handlerInscriptionDada(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/inscription_dada.html")
}

func handlerConnexion(w http.ResponseWriter, r *http.Request) {
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
			cookies.HandlerCookie(w, r, userID) // Ajout du cookie de session
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
