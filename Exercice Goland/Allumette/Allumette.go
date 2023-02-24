package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Demande au joueur de choisir le nombre d'allumettes
	var n int
	fmt.Print("Combien d'allumettes voulez-vous ? (minimum 4) ")
	fmt.Scanln(&n)

	// Vérifie que le nombre d'allumettes est supérieur ou égal à 4
	if n < 4 {
		fmt.Println("Le nombre d'allumettes doit être au moins égal à 4.")
		return
	}

	// Boucle de jeu
	for {
		// Affichage du nombre d'allumettes restantes
		fmt.Println("Il reste", n, "allumettes.")

		// Tour du joueur
		var p int
		for {
			fmt.Print("Combien d'allumettes voulez-vous prendre ? (1-3) ")
			fmt.Scanln(&p)
			if p >= 1 && p <= 3 && p <= n {
				break
			}
			fmt.Println("Vous devez prendre entre 1 et 3 allumettes, et il doit en rester suffisamment.")
		}
		n -= p

		// Vérifie si le joueur a perdu
		if n == 0 {
			fmt.Println("Vous avez pris la dernière allumette. Vous avez perdu !")
			break
		}

		// Tour de l'ordinateur
		var c int
		if n == 1 {
			c = 1
		} else {
			c = rand.Intn(3) + 1
			if c > n {
				c = n
			}
		}
		fmt.Println("L'ordinateur prend", c, "allumettes.")
		n -= c

		// Vérifie si l'ordinateur a perdu
		if n == 0 {
			fmt.Println("L'ordinateur a pris la dernière allumette. Vous avez gagné !")
			break
		}
	}
}
