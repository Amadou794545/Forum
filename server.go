package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"forum/Database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/inscription", HandlerInscription)
	http.HandleFunc("/inscriptionPicture", HandlerInscriptionPicture)
	http.HandleFunc("/userPicture", HandlerUserPicture)
	http.HandleFunc("/login", HandlerLogin)
	http.HandleFunc("/deconnect", HandlerDeconnect)
	http.HandleFunc("/created", HandlerCreated)
	http.HandleFunc("/liked", HandlerLiked)
	http.HandleFunc("/settings", HandlerSettings)

	http.HandleFunc("/comment", addCommentAPI)
	http.HandleFunc("/api/comments", getCommentsAPI)
	http.HandleFunc("/upload", uploadFile)

	http.HandleFunc("/api/posts", GetPostsAPI)
	http.HandleFunc("/api/user/posts", GetUserPostsAPI)

	fs := http.FileServer(http.Dir("/go/bin/CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))
	http.Handle("/java-script/", http.StripPrefix("/java-script", http.FileServer(http.Dir("java-script"))))

	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("uploads"))))
	http.Handle("/Pictures/", http.StripPrefix("/Pictures", http.FileServer(http.Dir("Pictures"))))

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	maxFileSize := int64(10 * 1024 * 1024) // Set the maximum file size to 10 MB
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	titre := r.FormValue("titre")
	description := r.FormValue("description")
	hobbies, err := strconv.Atoi(r.FormValue("dada"))
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	// Check if img is present
	file, handler, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	// If img present, handle it
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
		fmt.Println("Error retrieving session cookie:", err)
		return
	}

	userId, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Println("Error retrieving cookie value:", err)
		return
	}

	Database.AddPost(titre, imagePath, description, userId, hobbies)
	fmt.Fprintln(w, "File uploaded successfully!")
}
