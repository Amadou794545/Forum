package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Serve les fichiers JavaScript
	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	// Serve les fichiers CSS
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Définit les routes pour les pages
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/", Index)
	http.HandleFunc("/upload", uploadFile)

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
