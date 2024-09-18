package gokemon

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func Jaune(texte string) string {
	return "\033[33m" + texte + "\033[0m"
}

func Vert(texte string) string {
	return "\033[32m" + texte + "\033[0m"
}

func Rouge(texte string) string {
	return "\033[31m" + texte + "\033[0m"
}

func Cyan(texte string) string {
	return "\033[36m" + texte + "\033[0m"
}

func Bleu(texte string) string {
	return "\033[34m" + texte + "\033[0m"
}

func AfficherTitre() {
	titre := `
						 ██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗
						██╔════╝ ██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║
						██║  ███╗██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║
						██║   ██║██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║
						╚██████╔╝╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║
 						╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝`
	fmt.Println(Rouge(titre))
}

func AfficherCadre(largeur int) {
	fmt.Println(Jaune(strings.Repeat("═", largeur)))
}

func AfficherLigneMenu(texte string, largeur int) {
	espaces := largeur - len([]rune(texte)) - 6
	fmt.Printf("%s║ %s%s ║%s\n", Jaune("║"), Jaune(texte), strings.Repeat(" ", espaces), Jaune("║"))
}

func MessageRapide(message string, vitesse int, nomCouleur string) {
	c := color.New(color.FgWhite)
	switch nomCouleur {
	case "vert":
		c = color.New(color.FgGreen)
	case "rouge":
		c = color.New(color.FgRed)
	case "bleu":
		c = color.New(color.FgBlue)
	case "cyan":
		c = color.New(color.FgCyan)
	case "jaune":
		c = color.New(color.FgYellow)
	}

	for _, char := range message {
		fmt.Print(c.Sprint(string(char)))
		time.Sleep(time.Duration(vitesse) * time.Millisecond)
	}

	fmt.Println()
}
