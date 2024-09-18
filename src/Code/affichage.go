package gokemon

import (
	"fmt"
	"strings"
)

func Jaune(texte string) string {
	return "\033[33m" + texte + "\033[0m"
}

func Vert(texte string) string {
	return "\033[32m" + texte + "\033[0m"
}

func AfficherTitre() {
	titre := `
						 ██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗
						██╔════╝ ██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║
						██║  ███╗██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║
						██║   ██║██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║
						╚██████╔╝╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║
 						╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝`
	fmt.Println(Jaune(titre))
}

func AfficherCadre(largeur int) {
	fmt.Println(Jaune(strings.Repeat("═", largeur)))
}

func AfficherLigneMenu(texte string, largeur int) {
	espaces := largeur - len([]rune(texte)) - 6
	fmt.Printf("%s║ %s%s ║%s\n", Jaune("║"), Jaune(texte), strings.Repeat(" ", espaces), Jaune("║"))
}
