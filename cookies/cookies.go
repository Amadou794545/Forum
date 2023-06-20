package cookies

import (
	"net/http"
	"strconv"
	"time"
)

func HandlerSessionCookie(w http.ResponseWriter, r *http.Request, userID int) {
	cookie := &http.Cookie{
		Name:    "session",
		Value:   strconv.Itoa(userID),
		Expires: time.Now().Add(24 * time.Hour), //24h
		Path:    "/",
	}
	http.SetCookie(w, cookie)
}

func CheckSessionCookie(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false // Cookie pas présent
	}
	if cookie.Value != "" {
		return true //Cookie présent avec la bonne valeur
	}
	return false
}

func UpdateSessionExpiration(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		println("The cookie doesn't exist")
		return
	}
	// Mise à jour expiration du cookie +24h à partir de l'heure actuelle
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// Réécriture du cookie dans la réponse HTTP
	http.SetCookie(w, cookie)
}

func DeleteAllCookies(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}
