package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func randomTableau(dimensions []int) [][]int {
	tableau := make([][]int, dimensions[0])
	if len(dimensions) == 1 {
		for i := range tableau {
			tableau[i] = []int{rand.Intn(100)}
		}
	} else {
		for i := range tableau {
			tableau[i] = randomTableau(dimensions[1:])
		}
	}
	return tableau
}

func tableauToHTML(tableau [][]int) string {
	var sb strings.Builder
	sb.WriteString("<table>")
	for _, ligne := range tableau {
		sb.WriteString("<tr>")
		for _, element := range ligne {
			sb.WriteString("<td>")
			sb.WriteString(strconv.Itoa(element))
			sb.WriteString("</td>")
		}
		sb.WriteString("</tr>")
	}
	sb.WriteString("</table>")
	return sb.String()
}

func main() {
	// Génération d'un tableau aléatoire à n dimensions
	dimensions := []int{3, rand.Intn(10) + 1, rand.Intn(10) + 1}
	tableau := randomTableau(dimensions)
	fmt.Println(tableau)

	// Affichage du tableau sous forme de tableau HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, tableauToHTML(tableau))
	})
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
