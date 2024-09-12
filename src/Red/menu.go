package gokemon

import (
	"fmt"
	"os"
	"strings"
)

func creerDresseur(joueur *Dresseur) {
	if joueur.Nom == "" {
		fmt.Print(Vert("Entrez votre nom de dresseur : "))
		fmt.Scanln(&joueur.Nom)

		fmt.Println(Jaune("\nChoisissez votre Pokémon de départ :"))
		fmt.Println(Jaune("1. Bulbizarre (Type: Plante)"))
		fmt.Println(Jaune("2. Salamèche (Type: Feu)"))
		fmt.Println(Jaune("3. Carapuce (Type: Eau)"))

		var choixPokemon string
		fmt.Print(Vert("Entrez votre choix (1-3) : "))
		fmt.Scanln(&choixPokemon)

		var pokemon Pokemon
		switch choixPokemon {
		case "1":
			pokemon = Pokemon{Nom: "Bulbizarre", Type: "Plante", Niveau: 5, PVMax: 45, PVActuels: 45, Attaque: 10}
		case "2":
			pokemon = Pokemon{Nom: "Salamèche", Type: "Feu", Niveau: 5, PVMax: 39, PVActuels: 39, Attaque: 11}
		case "3":
			pokemon = Pokemon{Nom: "Carapuce", Type: "Eau", Niveau: 5, PVMax: 44, PVActuels: 44, Attaque: 9}
		default:
			fmt.Println(Jaune("Choix invalide. Pokémon par défaut : Pikachu"))
			pokemon = Pokemon{Nom: "Pikachu", Type: "Électrik", Niveau: 5, PVMax: 35, PVActuels: 35, Attaque: 12}
		}

		joueur.Equipe = append(joueur.Equipe, pokemon)
		fmt.Printf(Jaune("Félicitations, %s ! Vous avez choisi %s comme Pokémon de départ!\n"), joueur.Nom, pokemon.Nom)
	} else {
		fmt.Println(Jaune("Vous avez déjà créé votre dresseur."))
	}
}

func MenuPrincipal(joueur *Dresseur) {
	for {
		largeur := 50
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("        MENU PRINCIPAL", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("1. Créer Dresseur", largeur)
		AfficherLigneMenu("2. Afficher les informations du dresseur", largeur)
		AfficherLigneMenu("3. Accéder à l'inventaire", largeur)
		AfficherLigneMenu("4. Combat", largeur)
		AfficherLigneMenu("5. Quitter", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix (1-5): "))
		var choix string
		fmt.Scanln(&choix)

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
			AccessInventory(joueur)
		case "4":
			Combat(joueur)
		case "5":
			fmt.Println(Jaune("\nMerci d'avoir joué. Au revoir!"))
			os.Exit(0)
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		fmt.Scanln()
	}
}
