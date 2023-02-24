package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Définition des fonctions de rappel pour chaque page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenue sur la page d'accueil !")
	})

	http.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenue sur la page 1 !")
	})

	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenue sur la page 2 !")
	})

	// Fonction de rappel pour la page d'erreur 404
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erreur 404 - Page non trouvée")
	})

	// Redirection des pages
	http.HandleFunc("/redirect1", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/page1", http.StatusSeeOther)
	})

	http.HandleFunc("/redirect2", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/page2", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/redirect3", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
	})

	// Démarrage du serveur web
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
