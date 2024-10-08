package gokemon

import (
	"fmt"
	"strings"
)

func VisiterForgeron(joueur *Dresseur) {
	for {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		largeur := 155
		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("                                                              FORGERON", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu(fmt.Sprintf("  Votre Porte-Monnaie: %d PokéDollars", joueur.Argent), largeur)
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("  Capacité d'inventaire: "+fmt.Sprintf("%d/%d", len(joueur.Inventaire), joueur.CapaciteInventaire), largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu(" Votre inventaire :", largeur)
		for i, item := range joueur.Inventaire {
			AfficherLigneMenu(fmt.Sprintf("%d. %s (x%d)", i+1, item.Nom, item.Quantite), largeur)
		}

		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("1. Fabriquer un Casque (50 PokéDollars)[Fourrure:2 , Écaille:1]", largeur)
		AfficherLigneMenu("2. Fabriquer une Armure (50 PokéDollars)[Écaille: 3, Charbon: 1]", largeur)
		AfficherLigneMenu("3. Fabriquer des Bottes (50 PokéDollars)[Plume: 2, Carapace: 1]", largeur)
		AfficherLigneMenu("4. Retour au menu principal", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix (1-4): "))
		var choix string
		Wrap(func() { fmt.Scanln(&choix) })

		switch choix {
		case "1":
			FabriquerEquipement(joueur, "Casque", "Tête")
		case "2":
			FabriquerEquipement(joueur, "Armure", "Torse")
		case "3":
			FabriquerEquipement(joueur, "Bottes", "Pieds")
		case "4":
			return
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		if choix != "4" {
			fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
			Wrap(func() { fmt.Scanln() })
		}
	}
}

func FabriquerEquipement(joueur *Dresseur, nom string, emplacement string) {

	equipements := []Equipment{joueur.Equipement.Tete, joueur.Equipement.Torse, joueur.Equipement.Pieds}
	for _, equip := range equipements {
		if equip.Emplacement == emplacement {
			fmt.Printf(Jaune("\nVous avez déjà équipé un %s. Vous ne pouvez pas en fabriquer un autre.\n"), nom)
			return
		}
	}

	recette, existe := CraftingRecipes[nom]
	if !existe {
		fmt.Println(Jaune("\nRecette introuvable."))
		return
	}

	if joueur.Argent < recette.CoutArgent {
		fmt.Println(Jaune("\nVous n'avez pas assez d'argent pour fabriquer cet équipement."))
		return
	}

	for ressource, quantiteRequise := range recette.Ressources {
		quantitePossedee := 0
		for _, item := range joueur.Inventaire {
			if item.Nom == ressource {
				quantitePossedee = item.Quantite
				break
			}
		}
		if quantitePossedee < quantiteRequise {
			fmt.Printf(Jaune("\nVous n'avez pas assez de %s pour fabriquer cet équipement.\n"), ressource)
			return
		}
	}

	if len(joueur.Inventaire) >= joueur.CapaciteInventaire {
		fmt.Println(Jaune("\nVotre inventaire est plein. Vous ne pouvez pas fabriquer cet équipement."))
		return
	}

	joueur.Argent -= recette.CoutArgent
	for ressource, quantiteRequise := range recette.Ressources {
		for i := range joueur.Inventaire {
			if joueur.Inventaire[i].Nom == ressource {
				joueur.Inventaire[i].Quantite -= quantiteRequise
				if joueur.Inventaire[i].Quantite == 0 {
					joueur.Inventaire = append(joueur.Inventaire[:i], joueur.Inventaire[i+1:]...)
				}
				break
			}
		}
	}

	var bonusPV, bonusAttack int
	switch emplacement {
	case "Tête":
		bonusPV = 15
		bonusAttack = 5
	case "Torse":
		bonusPV = 25
		bonusAttack = 10
	case "Pieds":
		bonusPV = 10
		bonusAttack = 15
	}

	nouvelEquipement := Equipment{Nom: nom, Emplacement: emplacement, BonusPV: bonusPV, BonusAttack: bonusAttack}
	joueur.EquiperEquipement(nouvelEquipement)

	fmt.Printf(Jaune("\nVous avez fabriqué et équipé un(e) %s pour %d PokéDollars et les ressources nécessaires.\n"), nom, recette.CoutArgent)
	fmt.Printf(Jaune("Bonus : PV +%d, Attaque +%d\n"), bonusPV, bonusAttack)
}
