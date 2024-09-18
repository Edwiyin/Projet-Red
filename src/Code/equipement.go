package gokemon

import (
	"fmt"
	"strings"
)

func AfficherEquipements(joueur *Dresseur) {
	largeur := 155
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	AfficherTitre()

	fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
	AfficherLigneMenu("", largeur)
	AfficherLigneMenu("                                                              ÉQUIPEMENTS", largeur)
	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))

	if joueur.Equipement.Tete == (Equipment{}) && joueur.Equipement.Torse == (Equipment{}) && joueur.Equipement.Pieds == (Equipment{}) {
		AfficherLigneMenu("Vous n'avez pas d'équipements pour le moment.", largeur)
	} else {
		equipements := []Equipment{joueur.Equipement.Tete, joueur.Equipement.Torse, joueur.Equipement.Pieds}
		for _, equip := range equipements {
			if equip != (Equipment{}) {
				AfficherLigneMenu(fmt.Sprintf("%s - Emplacement: %s, Bonus PV: %d, Bonus Attaque: %d",
					equip.Nom, equip.Emplacement, equip.BonusPV, equip.BonusAttack), largeur)
			}
		}
	}
	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

	fmt.Print(Vert("\nAppuyez sur Entrée pour revenir au menu principal..."))
	Wrap(func() { fmt.Scanln() })
}
