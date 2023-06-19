package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

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

var db *sql.DB

func main() {
	http.HandleFunc("/api/posts", GetPostsAPI)

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/inscription", handlerInscription)
	http.HandleFunc("/login", handlerConnexion)

	http.HandleFunc("/upload", uploadFile)

	http.HandleFunc("/created", handlerCreated)
	http.HandleFunc("/api/user/posts", GetUserPostsAPI)

	router := mux.NewRouter()

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":3030", router))

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("uploads"))))

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
	var table string
	var likeColumn string
	var action string
	commentID := 1
	like := true

	if commentID != 0 {
		table = "comments"
		likeColumn = "comment_likes"
	} else {
		table = "posts"
		likeColumn = "likes"
	}

	if like {
		action = "like"
	} else {
		action = "dislike"
	}

	fmt.Println(table)
	fmt.Println(likeColumn)
	fmt.Println(action)
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

	// Set the maximum file size to 10 MB
	maxFileSize := int64(10 * 1024 * 1024)
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Access the form values after parsing the multipart form data
	titre := r.FormValue("titre")
	description := r.FormValue("description")
	hobbies, err := strconv.Atoi(r.FormValue("dada"))

	// Check if an image file is present
	file, handler, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		// Other error occurred while retrieving the file, handle accordingly
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	// If an image file is provided, handle it
	var imagePath string
	if file != nil {
		defer file.Close()

		// Generate a unique number or timestamp to append to the filename
		uniqueNumber := time.Now().UnixNano()

		// Get the file extension
		extension := filepath.Ext(handler.Filename)

		// Generate the new filename
		newFilename := fmt.Sprintf("Post-%d%s", uniqueNumber, extension)

		// Save the file on the server with the new filename
		filePath := filepath.Join("./uploads", newFilename)
		outFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Error saving the file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Error(w, "Error saving the file", http.StatusInternalServerError)
			return
		}

		imagePath = "images/" + newFilename
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		// Handle the error if needed
		fmt.Println("Error retrieving session cookie:", err)
		return
	}

	userId, err := strconv.Atoi(cookie.Value)
	if err != nil {
		// Handle the error if needed
		fmt.Println("Error retrieving cookie value:", err)
		return
	}

	Database.AddPost(titre, imagePath, description, userId, hobbies)
	fmt.Fprintln(w, "File uploaded successfully!")
}

func handlerCreated(w http.ResponseWriter, r *http.Request) {
	if cookies.CheckSessionCookie(r) {
		_, err := r.Cookie("session")
		if err != nil {
			// Handle the error if needed
			fmt.Println("Error retrieving session cookie:", err)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	http.ServeFile(w, r, "template/userPage.html")
}

func GetUserPostsAPI(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		// Handle the error if needed
		fmt.Println("Error retrieving session cookie:", err)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	UserID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		// Handle the error if needed
		fmt.Println("Error retrieving cookie value:", err)
		return
	}

	// Fetch posts by user ID from the database
	posts, err := Database.GetUserPosts(UserID)
	if err != nil {
		log.Println("Error retrieving user posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(posts)
	if err != nil {
		log.Println("Error converting to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
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

	http.ServeFile(w, r, "/template/index.html")
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

func init() {
	// Connect to the PostgreSQL database
	connStr := "user=yourusername password=yourpassword dbname=yourdbname sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
