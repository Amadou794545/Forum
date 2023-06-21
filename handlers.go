package main

import (
	"fmt"
	"forum/Database"
	"forum/cookies"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/index.html"))
	imgPathDta := ImgPathData{
		ImgPath: "",
	}
	if cookies.CheckSessionCookie(r) {
		cookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Error retrieving session cookie :", err)
			return
		}
		userID := cookie.Value
		username, err := Database.GetUserUsername(userID)
		if err != nil {
			fmt.Println("Error retrieving username :", err)
			return
		}
		cookies.UpdateSessionExpiration(w, r) // Reset la date de péremption du cookie
		fmt.Println("Bienvenue", username)
		imgPathDta.ImgPath, err = Database.GetUserImg(userID)
		if err != nil {
			fmt.Println("Error retrieving imgPath :", err)
			return
		}
	}
	tmpl.Execute(w, imgPathDta)
}

func HandlerInscription(w http.ResponseWriter, r *http.Request) {
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
			imgPath := "Pictures/Profil/anonyme.jpg"
			Database.AddUser(email, username, password, imgPath)
			userID, err := Database.GetUserID(username)
			if err != nil {
				fmt.Println("Error getting user ID:", err)
				return
			}
			num, err := strconv.Atoi(userID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cookies.HandlerSessionCookie(w, r, num)
			http.Redirect(w, r, "/inscriptionPicture", http.StatusFound)
		} else {
			InscriptionData.Username = username
			InscriptionData.Email = email
			InscriptionData.ErrorMessage = errorMessage
		}

		userID, err := Database.GetUserID(username)
		if err != nil {
			fmt.Println("Error getting user ID:", err)
			return
		}
		num, err := strconv.Atoi(userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		cookies.HandlerSessionCookie(w, r, num) // Ajout du cookie de session
		tmpl.Execute(w, InscriptionData)
	}
}

func HandlerInscriptionPicture(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/inscription_picture.html")
}

func HandlerUserPicture(w http.ResponseWriter, r *http.Request) {
	maxFileSize := int64(10 * 1024 * 1024)
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("uploadInput")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	if file != nil { //img uplodée
		defer file.Close()
		// Unique nbr
		uniqueNumber := time.Now().UnixNano()
		// Get file extension
		extension := filepath.Ext(handler.Filename)
		// Generate the new filename
		newFilename := fmt.Sprintf("Img-%d%s", uniqueNumber, extension)
		// Get new path
		filePath := filepath.Join("./Pictures/uploads", newFilename)
		// Save the file on the server with the new filename
		outFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Error saving the file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()
		imgPath := "Pictures/uploads/" + newFilename
		cookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Error retrieving session cookie:", err)
			return
		}
		userID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			fmt.Println("Error retrieving userID:", err)
			return
		}
		Database.UpdateImgProfile(imgPath, userID)
	} else { //img par défault
		path := r.FormValue("selectedPicture")
		filename := filepath.Base(path)
		cookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Error retrieving session cookie:", err)
			return
		}
		userID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			fmt.Println("Error retrieving userID:", err)
			return
		}
		Database.UpdateImgProfile("Pictures/Profil/"+filename, userID)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/login.html"))
	loginData := LoginData{
		Username:     "",
		ErrorMessage: "",
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		if Database.CheckLogin(username, password) {
			userID, err := Database.GetUserID(username)
			if err != nil {
				fmt.Println("Error:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			num, err := strconv.Atoi(userID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cookies.HandlerSessionCookie(w, r, num) // Ajout du cookie de session
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

func HandlerDeconnect(w http.ResponseWriter, r *http.Request) {
	cookies.DeleteAllCookies(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandlerCreated(w http.ResponseWriter, r *http.Request) {
	if cookies.CheckSessionCookie(r) {
		_, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Error retrieving session cookie:", err)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	http.ServeFile(w, r, "template/userPage.html")
}

func HandlerLiked(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/liked.html")
}

func HandlerSettings(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/settings.html"))
	imgPathDta := ImgPathData{
		ImgPath: "",
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error retrieving session cookie :", err)
		return
	}
	userID := cookie.Value
	imgPathDta.ImgPath, err = Database.GetUserImg(userID)
	if err != nil {
		fmt.Println("Error retrieving imgPath :", err)
		return
	}
	username := r.FormValue("username")
	if username == "" {
		imgPathDta.UsernameMessage = "Entrez un pseudo pour lancer la vérification"
	} else if Database.CheckUsername(username) {
		imgPathDta.UsernameMessage = "Pseudo déjà utilisé"
	} else {
		imgPathDta.UsernameMessage = "Pseudo modifié avec succès"
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			fmt.Println("Error convert from string to int :", err)
			return
		}
		Database.UpdateUsername(username, userIDInt)
	}
	tmpl.Execute(w, imgPathDta)
}
