package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for {
		// Affichage du menu
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1. Récupérer tout le texte contenu dans un fichier")
		fmt.Println("2. Ajouter du texte dans un fichier")
		fmt.Println("3. Supprimer tout le contenu d'un fichier")
		fmt.Println("4. Remplacer le contenu d'un fichier par du texte donné par l'utilisateur")
		fmt.Println("0. Quitter")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// Récupération du texte contenu dans un fichier
			fmt.Print("Entrez le nom du fichier : ")
			var nomFichier string
			fmt.Scanln(&nomFichier)
			contenuFichier, err := ioutil.ReadFile(nomFichier)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier :", err)
				continue
			}
			fmt.Println("Contenu du fichier :", string(contenuFichier))
		case 2:
			// Ajout de texte dans un fichier
			fmt.Print("Entrez le nom du fichier : ")
			var nomFichier string
			fmt.Scanln(&nomFichier)
			fmt.Print("Entrez le texte à ajouter : ")
			var texte string
			fmt.Scanln(&texte)
			err := ioutil.WriteFile(nomFichier, []byte(texte), os.ModeAppend)
			if err != nil {
				fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
				continue
			}
			fmt.Println("Le texte a été ajouté dans le fichier.")
		case 3:
			// Suppression du contenu d'un fichier
			fmt.Print("Entrez le nom du fichier : ")
			var nomFichier string
			fmt.Scanln(&nomFichier)
			err := ioutil.WriteFile(nomFichier, []byte(""), 0)
			if err != nil {
				fmt.Println("Erreur lors de la suppression du contenu du fichier :", err)
				continue
			}
			fmt.Println("Le contenu du fichier a été supprimé.")
		case 4:
			// Remplacement du contenu d'un fichier par du texte donné par l'utilisateur
			fmt.Print("Entrez le nom du fichier : ")
			var nomFichier string
			fmt.Scanln(&nomFichier)
			fmt.Print("Entrez le texte à ajouter : ")
			var texte string
			fmt.Scanln(&texte)
			err := ioutil.WriteFile(nomFichier, []byte(texte), 0)
			if err != nil {
				fmt.Println("Erreur lors du remplacement du contenu du fichier :", err)
				continue
			}
			fmt.Println("Le contenu du fichier a été remplacé.")
		case 0:
			// Sortie du programme
			fmt.Println("Au revoir !")
			return
		default:
			// Choix invalide
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}
