package main

import (
	"encoding/json"
	"fmt"
	"forum/Database"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetPostsAPI(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	filterValues := r.URL.Query().Get("filters")

	var filters []int
	if filterValues != "" {
		filterStrings := strings.Split(filterValues, ",")
		for _, filterString := range filterStrings {
			filterInt, err := strconv.Atoi(filterString)
			if err == nil {
				filters = append(filters, filterInt)
			}
		}
	}

	// Convert page and limit values to integers
	pageNum, _ := strconv.Atoi(page)
	limitNum, _ := strconv.Atoi(limit)

	// Calculate the offset based on the requested page and limit
	offset := (pageNum - 1) * limitNum

	// Fetch posts from the database with pagination
	posts, err := Database.GetPosts(offset, limitNum, filters)
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

func CommentLikeDislike(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody struct {
		PostID    string `json:"postId"`
		LikeValue int    `json:"likeValue"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		// Handle JSON decoding error
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		// Gérer l'erreur si nécessaire
		fmt.Println("Erreur lors de la récupération du cookie de session :", err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	UserID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		// Gérer l'erreur si nécessaire
		fmt.Println("Erreur lors de la récupération de la valeur du cookie :", err)
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// Access the values received
	CommentID := requestBody.PostID
	likeValue := requestBody.LikeValue
	commentIDint, _ := strconv.Atoi(CommentID)

	if likeValue == -1 {
		fmt.Println("testdislike")
		fmt.Println(Database.CheckCommentDislikes(UserID, commentIDint))
		if Database.CheckCommentDislikes(UserID, commentIDint) {
			Database.AddDislikecomment(UserID, commentIDint)
		}
		if !Database.CheckCommentLikes(UserID, commentIDint) {
			Database.RemoveLikeComment(UserID, commentIDint)
		}
	} else if likeValue == 1 {
		if Database.CheckCommentLikes(UserID, commentIDint) {
			Database.AddLikeComment(UserID, commentIDint)
		}
		if !Database.CheckCommentDislikes(UserID, commentIDint) {
			Database.RemoveDislikeComment(UserID, commentIDint)
		}
	} else {
		fmt.Println("testcase0")
		Database.RemoveDislikeComment(UserID, commentIDint)
		Database.RemoveLikeComment(UserID, commentIDint)
	}

	fmt.Println(commentIDint)
	fmt.Println(likeValue)
}

func PostLikeDislike(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody struct {
		PostID    string `json:"postId"`
		LikeValue int    `json:"likeValue"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		// Handle JSON decoding error
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		// Gérer l'erreur si nécessaire
		fmt.Println("Erreur lors de la récupération du cookie de session :", err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	UserID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		// Gérer l'erreur si nécessaire
		fmt.Println("Erreur lors de la récupération de la valeur du cookie :", err)
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// Access the values received
	postID := requestBody.PostID
	likeValue := requestBody.LikeValue
	postIDint, _ := strconv.Atoi(postID)

	if likeValue == -1 {
		fmt.Println("testdislike")
		fmt.Println(Database.CheckDislikes(UserID, postIDint))
		if Database.CheckDislikes(UserID, postIDint) {
			Database.AddDislikePost(UserID, postIDint)
		}
		if !Database.CheckLikes(UserID, postIDint) {
			Database.RemoveLikePost(UserID, postIDint)
		}
	} else if likeValue == 1 {
		if Database.CheckLikes(UserID, postIDint) {
			Database.AddLikePost(UserID, postIDint)
		}
		if !Database.CheckDislikes(UserID, postIDint) {
			Database.RemoveDislikePost(UserID, postIDint)
		}
	} else {
		fmt.Println("testcase0")
		Database.RemoveDislikePost(UserID, postIDint)
		Database.RemoveLikePost(UserID, postIDint)
	}

	fmt.Println(postID)
	fmt.Println(likeValue)
}

func GetUserlikedPostsAPI(w http.ResponseWriter, r *http.Request) {
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
	posts, err := Database.GetUserLikedPosts(UserID)
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

func GetUserPostsAPI(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error retrieving session cookie:", err)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	UserID, err := strconv.Atoi(cookie.Value)
	if err != nil {
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

func addCommentAPI(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		// Gérer l'erreur si nécessaire
		fmt.Println("Erreur lors de la récupération du cookie de session :", err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	UserID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		// Gérer l'erreur si nécessaire
		fmt.Println("Erreur lors de la récupération de la valeur du cookie :", err)
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}
	// Décoder le corps de la requête
	var data struct {
		PostID         string `json:"postID"`
		CommentContent string `json:"commentContent"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du corps de la requête :", err)
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Erreur lors du décodage du corps de la requête :", err)
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}
	// Ajouter le commentaire à la base de données
	Database.AddComment(data.CommentContent, UserID, data.PostID)
	// Envoyer la réponse avec les commentaires au format JSON
	sendCommentsResponse(w, data.PostID)
}

func getCommentsAPI(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("postId")
	if postID == "" {
		http.Error(w, "Paramètre postId manquant", http.StatusBadRequest)
		return
	}
	sendCommentsResponse(w, postID)
}

func sendCommentsResponse(w http.ResponseWriter, postID string) {
	comments, err := Database.GetComment(postID)
	fmt.Println(Database.GetComment(postID))
	if err != nil {
		log.Println("Erreur lors de la récupération des commentaires :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.MarshalIndent(comments, "", "  ")
	if err != nil {
		log.Println("Erreur lors de la conversion en JSON :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
