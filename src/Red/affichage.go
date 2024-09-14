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

func DisplayInfo(joueur Dresseur) {
	fmt.Println(Jaune("\nInformations du dresseur :"))
	fmt.Printf(Jaune("Nom : %s\n"), joueur.Nom)
	fmt.Printf(Jaune("Solde : %d PokéDollars\n"), joueur.Argent)
	fmt.Println(Jaune("\nÉquipe Pokémon :"))
	for _, pokemon := range joueur.Equipe {
		fmt.Printf(Jaune("- %s (Type: %s, Niveau: %d, PV: %d/%d, Attaque: %d)\n"), pokemon.Nom, pokemon.Type, pokemon.Niveau, pokemon.PVActuels, pokemon.PVMax, pokemon.Attaque)
	}
	fmt.Println(Jaune("\nInventaire :"))
	for _, item := range joueur.Inventaire {
		fmt.Printf(Jaune("- %s (x%d)\n"), item.Nom, item.Quantite)
	}
}
