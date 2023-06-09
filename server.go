package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	if r.Method == "POST" {
		titre := r.FormValue("titre")
		description := r.FormValue("description")
		//image, _, err := r.FormFile("image")
		//if err != nil {
		//	fmt.Println("Erreur lors de la récupération du fichier d'image :", err)
		//} else {
		//defer image.Close()
		fmt.Println("Titre :", titre)
		fmt.Println("Description :", description)
		//fmt.Println("Chemin de l'image :", image.Filename)

	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(2 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the file from the form data
	file, handler, err := r.FormFile("image")
	if err != nil {
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
