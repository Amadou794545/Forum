package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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
	http.HandleFunc("/api/posts", GetPostsAPI)

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/inscription", handlerInscription)
	http.HandleFunc("/login", handlerConnexion)

	http.HandleFunc("/upload", uploadFile)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	// Parse the multipart form data
	titre := r.FormValue("titre")
	description := r.FormValue("description")

	fmt.Println(titre + " " + description)
	// Set the maximum file size to 10 MB
	maxFileSize := int64(10 * 1024 * 1024)
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if an image file is present
	file, handler, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			// No image file provided, handle accordingly
			fmt.Fprintln(w, "No image file provided.")
			return
		}
		// Other error occurred while retrieving the file, handle accordingly
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Read the file content
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}
	// Save the file on the server (you can change the path as per your requirement)
	filepath := "./uploads/" + handler.Filename
	err = ioutil.WriteFile(filepath, fileBytes, 0644)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	// Handle the successful file upload here (e.g., show a success message)
	fmt.Fprintln(w, "File uploaded successfully!")
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if cookies.CheckSessionCookie(r) {
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
