package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Si la méthode HTTP est POST, cela signifie que le formulaire a été soumis
		if r.Method == "POST" {
			// Récupération des données saisies dans le formulaire
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Affichage des données saisies dans la console en Go
			fmt.Printf("Nom d'utilisateur : %s\n", username)
			fmt.Printf("Mot de passe : %s\n", password)

			// Réponse HTTP pour confirmer la soumission du formulaire
			w.Write([]byte("Le formulaire a été soumis avec succès !"))
			return
		}

		// Si la méthode HTTP est GET, cela signifie que la page doit être affichée pour la première fois
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<form method="POST">
				<label for="username">Nom d'utilisateur :</label>
				<input type="text" name="username"><br>

				<label for="password">Mot de passe :</label>
				<input type="password" name="password"><br>

				<button type="submit">Soumettre</button>
			</form>
		`)
	})

	// Démarrage du serveur web
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
