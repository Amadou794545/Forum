package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/inscription", handlerInscription)
	http.HandleFunc("/login", handlerConnexion)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	port := ":3030"
	fmt.Printf("Serveur en cours d'exécution sur le port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}
