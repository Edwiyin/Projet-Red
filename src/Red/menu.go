package gokemon

import (
	"fmt"
	"os"
	"strings"
)

func TakePot(item *Item, joueur *Dresseur) {
	if item.Quantite > 0 {
		for i := range joueur.Equipe {
			if joueur.Equipe[i].PVActuels < joueur.Equipe[i].PVMax {
				joueur.Equipe[i].PVActuels += 20
				if joueur.Equipe[i].PVActuels > joueur.Equipe[i].PVMax {
					joueur.Equipe[i].PVActuels = joueur.Equipe[i].PVMax
				}
				item.Quantite--
				fmt.Printf(Jaune("\nVous avez utilisé une Potion sur %s. PV actuels: %d/%d\n"), joueur.Equipe[i].Nom, joueur.Equipe[i].PVActuels, joueur.Equipe[i].PVMax)
				return
			}
		}
		fmt.Println(Jaune("\nTous vos Pokémon ont déjà leurs PV au maximum."))
	} else {
		fmt.Println(Jaune("\nVous n'avez plus de Potions."))
	}
}



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
	for {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		largeur := 50
		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("        INVENTAIRE", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))

		for i, item := range joueur.Inventaire {
			AfficherLigneMenu(fmt.Sprintf("%d. %s (x%d)", i+1, item.Nom, item.Quantite), largeur)
		}
		AfficherLigneMenu(fmt.Sprintf("%d. Retour au menu principal", len(joueur.Inventaire)+1), largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix : "))
		var choix int
		fmt.Scanln(&choix)

		if choix == len(joueur.Inventaire)+1 {
			return
		} else if choix > 0 && choix <= len(joueur.Inventaire) {
			item := &joueur.Inventaire[choix-1]
			if item.Nom == "Potion" {
				TakePot((*Item)(item), joueur)
			} else {
				fmt.Printf(Jaune("\nVous ne pouvez pas utiliser %s pour le moment.\n"), item.Nom)
			}
		} else {
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		fmt.Scanln()
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
