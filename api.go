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

func GetCommentAPI(w http.ResponseWriter, r *http.Request) {
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

	// Decode the request body
	var data struct {
		PostID         string `json:"postID"`
		CommentContent string `json:"commentContent"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Access the comment data
	fmt.Println("Post ID:", data.PostID)
	fmt.Println("Comment Content:", data.CommentContent)
	Database.AddComment(data.CommentContent, UserID, data.PostID)
}
