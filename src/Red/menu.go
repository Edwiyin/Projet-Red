package gokemon

import (
	"fmt"
	"os"
	"strings"
)

func choixPokemonFunc(choixPokemon string) Pokemon {
	switch choixPokemon {
	case "1":
		return NewPokemon("Bulbizarre", Grass, 5)
	case "2":
		return NewPokemon("Salamèche", Fire, 5)
	case "3":
		return NewPokemon("Carapuce", Water, 5)
	default:
		fmt.Println(Jaune("Choix invalide. Pokémon par défaut : Pikachu"))
		return NewPokemon("Pikachu", Electric, 5)
	}
}

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

		pokemon := choixPokemonFunc(choixPokemon)
		joueur.Equipe = append(joueur.Equipe, pokemon)
		fmt.Printf(Jaune("Félicitations, %s ! Vous avez choisi %s comme Pokémon de départ!\n"), joueur.Nom, pokemon.Nom)
	} else {
		fmt.Println(Jaune("Vous avez déjà créé votre dresseur."))
	}
}

func ViewTeam(joueur *Dresseur) {
	if len(joueur.Equipe) == 0 {
		fmt.Println(Jaune("\nVous n'avez pas encore de Pokémon dans votre équipe."))
		return
	}

	fmt.Println(Jaune("\nVotre équipe Pokémon :"))
	for i, pokemon := range joueur.Equipe {
		fmt.Printf(Jaune("%d. %s (Type: %s, Niveau: %d, PV: %d/%d, Attaque: %d, Exp: %d/%d)\n"),
			i+1, pokemon.Nom, pokemon.Type, pokemon.Niveau, pokemon.PVActuels, pokemon.PVMax,
			pokemon.Attaque, pokemon.Experience, pokemon.ExperienceToNextLevel)
	}
}

func AccessInventory(joueur *Dresseur) {

	fmt.Println(Jaune("\nAccès à l'inventaire..."))
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
		AfficherLigneMenu("1. Créer Dresseur  ", largeur)
		AfficherLigneMenu("2. Afficher les informations du dresseur", largeur)
		AfficherLigneMenu("3. Accéder à l'inventaire  ", largeur)
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
