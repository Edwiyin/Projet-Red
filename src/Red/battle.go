package gokemon

import (
	"fmt"
	"math/rand"
	"strings"
)

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
				TakePot(item, joueur)
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

func TakePot(potion *Item, joueur *Dresseur) {
	if potion.Quantite > 0 {
		if len(joueur.Equipe) > 0 {
			pokemon := &joueur.Equipe[0]
			if pokemon.PVActuels < pokemon.PVMax {
				pokemon.PVActuels += 50
				if pokemon.PVActuels > pokemon.PVMax {
					pokemon.PVActuels = pokemon.PVMax
				}
				potion.Quantite--
				fmt.Printf(Jaune("\nVous avez utilisé une Potion. %s a récupéré 50 PV. PV actuels : %d/%d\n"), pokemon.Nom, pokemon.PVActuels, pokemon.PVMax)
			} else {
				fmt.Println(Jaune("\nVotre Pokémon a déjà tous ses PV."))
			}
		} else {
			fmt.Println(Jaune("\nVous n'avez pas de Pokémon dans votre équipe."))
		}
	} else {
		fmt.Println(Jaune("\nVous n'avez plus de Potions."))
	}
}

func EstVivant(p *Pokemon) bool {
	return p.PVActuels > 0
}

func Attaquer(p *Pokemon, cible *Pokemon) {
	degats := rand.Intn(p.Attaque) + 1
	cible.PVActuels -= degats
	if cible.PVActuels < 0 {
		cible.PVActuels = 0
	}
	fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", p.Nom, cible.Nom, degats)
	fmt.Printf("%s a maintenant %d PV\n", cible.Nom, cible.PVActuels)
}

func Combat(joueur *Dresseur) {
	if len(joueur.Equipe) == 0 {
		fmt.Println(Jaune("\nVous n'avez pas de Pokémon pour combattre. Créez d'abord votre dresseur."))
		return
	}

	pokemonJoueur := &joueur.Equipe[0]
	ennemi := &Pokemon{Nom: "Roucool", Type: "Vol", Niveau: 5, PVMax: 40, PVActuels: 40, Attaque: 8}

	fmt.Println(Jaune("\nUn Roucool sauvage apparaît!"))

	for EstVivant(pokemonJoueur) && EstVivant(ennemi) {
		fmt.Println(Jaune("\nQue voulez-vous faire ?"))
		fmt.Println(Jaune("1. Attaquer"))
		fmt.Println(Jaune("2. Utiliser une Potion"))
		fmt.Println(Jaune("3. Fuir"))

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			Attaquer(pokemonJoueur, ennemi)
			if EstVivant(ennemi) {
				Attaquer(ennemi, pokemonJoueur)
			}
		case 2:
			// Logique pour utiliser une potion
		case 3:
			fmt.Println(Jaune("Vous avez fui le combat!"))
			return
		default:
			fmt.Println(Jaune("Choix invalide."))
		}
	}

	if EstVivant(pokemonJoueur) {
		fmt.Println(Jaune("Vous avez gagné le combat!"))
	} else {
		fmt.Println(Jaune("Vous avez perdu le combat..."))
	}
}
