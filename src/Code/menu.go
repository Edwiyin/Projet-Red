package gokemon

import (
	"fmt"
	"os"
	"strings"
)

var audioManager *AudioManager

func creerDresseur(joueur *Dresseur) {
	if joueur.Nom == "" {
		fmt.Print(Vert("Entrez votre nom de dresseur : "))
		Wrap(func() { fmt.Scanln(&joueur.Nom) })

		fmt.Println(Jaune("\nChoisissez votre Pokémon de départ :"))
		fmt.Println(Jaune("1. Bulbizarre (Type: Plante)"))
		fmt.Println(Jaune("2. Salamèche (Type: Feu)"))
		fmt.Println(Jaune("3. Carapuce (Type: Eau)"))

		var choixPokemon string
		fmt.Print(Vert("Entrez votre choix (1-3) : "))
		Wrap(func() { fmt.Scanln(&choixPokemon) })

		pokemon := choixPokemonFunc(choixPokemon)
		joueur.Equipe = append(joueur.Equipe, pokemon)
		fmt.Printf(Jaune("Félicitations, %s ! Vous avez choisi %s comme Pokémon de départ!\n"), joueur.Nom, pokemon.Nom)
	} else {
		fmt.Println(Jaune("Vous avez déjà créé votre dresseur."))
	}
}

func MenuPrincipal(joueur *Dresseur, newAudioManager *AudioManager) {
	audioManager = newAudioManager
	largeur := 155
	joueur.Argent += 100
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	AfficherTitre()

	fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
	AfficherLigneMenu("", largeur)
	AfficherLigneMenu("                                                                  NEW GAME", largeur)
	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

	fmt.Print(Vert("\nAppuyez sur Entrée pour commencer..."))
	Wrap(func() { fmt.Scanln() })
	for {
		largeur := 155
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("                                                              MENU PRINCIPAL", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("1. Créer Dresseur  ", largeur)
		AfficherLigneMenu("2. Afficher les informations du dresseur", largeur)
		AfficherLigneMenu("3. Afficher les équipements", largeur)
		AfficherLigneMenu("4. Accéder à l'inventaire  ", largeur)
		AfficherLigneMenu("5. Combatre un Pokémon Sauvage", largeur)
		AfficherLigneMenu("6. Visiter le Marchand", largeur)
		AfficherLigneMenu("7. Visiter le Forgeron", largeur)
		AfficherLigneMenu("8. Qui sont-ils", largeur)
		AfficherLigneMenu("9. Quitter le Jeu", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix (1-8): "))
		var choix string
		Wrap(func() { fmt.Scanln(&choix) })

		switch choix {
		case "1":
			creerDresseur(joueur)
		case "2":
			if joueur.Nom == "" {
				fmt.Println(Jaune("\nVeuillez d'abord créer votre dresseur."))
			} else {
				DisplayInfo(*joueur)
			}
		case "3":
			AfficherEquipements(joueur)
		case "4":
			AccessInventory(joueur)
		case "5":
			Combat(joueur)
		case "6":
			VisiteMarchand(joueur)
		case "7":
			VisiterForgeron(joueur)
		case "8":
			MessageRapide(("Abba"), 3, "bleu")
			MessageRapide(("Steven Spielberg"), 3, "bleu")
			MessageRapide(("Les développeurs de ce jeu sont: Massinissa Ahfir, Edwin Wehbe, Michel Mustafaov"), 3,	"bleu")
		case "9":
			fmt.Println(Jaune("\nMerci d'avoir joué. Au revoir!"))
			os.Exit(0)
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		Wrap(func() { fmt.Scanln() })
	}
}

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
			AfficherLigneMenu(fmt.Sprintf("%s - Emplacement: %s, Bonus PV: %d, Bonus Attaque: %d", 
				equip.Nom, equip.Emplacement, equip.BonusPV, equip.BonusAttack), largeur)
		}
	}

	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

	fmt.Print(Vert("\nAppuyez sur Entrée pour revenir au menu principal..."))
	Wrap(func() { fmt.Scanln() })
}
