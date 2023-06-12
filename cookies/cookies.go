package cookies

import (
	"net/http"
	"time"
)

func HandlerCookie(w http.ResponseWriter, r *http.Request, userID string) {
	cookie := &http.Cookie{
		Name:    "session",
		Value:   userID,
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
