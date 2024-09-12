package main

import (
	"fmt"
	gokemon "gokemon/Red"
	"math/rand"
	"os"
	"strings"
	"time"
)

var joueur gokemon.Dresseur

func init() {
	joueur = gokemon.Dresseur{
		Nom:    "",
		Equipe: []gokemon.Pokemon{},
		Inventaire: []gokemon.Item{
			{Nom: "Potion", Quantite: 3},
			{Nom: "Pokéball", Quantite: 5},
			{Nom: "Baie", Quantite: 2},
		},
	}
}

func creerDresseur() {
	if joueur.Nom == "" {
		fmt.Print(gokemon.Vert("Entrez votre nom de dresseur : "))
		fmt.Scanln(&joueur.Nom)

		fmt.Println(gokemon.Jaune("\nChoisissez votre Pokémon de départ :"))
		fmt.Println(gokemon.Jaune("1. Bulbizarre (Type: Plante)"))
		fmt.Println(gokemon.Jaune("2. Salamèche (Type: Feu)"))
		fmt.Println(gokemon.Jaune("3. Carapuce (Type: Eau)"))

		var choixPokemon string
		fmt.Print(gokemon.Vert("Entrez votre choix (1-3) : "))
		fmt.Scanln(&choixPokemon)

		var pokemon gokemon.Pokemon
		switch choixPokemon {
		case "1":
			pokemon = gokemon.Pokemon{Nom: "Bulbizarre", Type: "Plante", Niveau: 5, PVMax: 45, PVActuels: 45, Attaque: 10}
		case "2":
			pokemon = gokemon.Pokemon{Nom: "Salamèche", Type: "Feu", Niveau: 5, PVMax: 39, PVActuels: 39, Attaque: 11}
		case "3":
			pokemon = gokemon.Pokemon{Nom: "Carapuce", Type: "Eau", Niveau: 5, PVMax: 44, PVActuels: 44, Attaque: 9}
		default:
			fmt.Println(gokemon.Jaune("Choix invalide. Pokémon par défaut : Pikachu"))
			pokemon = gokemon.Pokemon{Nom: "Pikachu", Type: "Électrik", Niveau: 5, PVMax: 35, PVActuels: 35, Attaque: 12}
		}

		joueur.Equipe = append(joueur.Equipe, pokemon)
		fmt.Printf(gokemon.Jaune("Félicitations, %s ! Vous avez choisi %s comme Pokémon de départ!\n"), joueur.Nom, pokemon.Nom)
	} else {
		fmt.Println(gokemon.Jaune("Vous avez déjà créé votre dresseur."))
	}
}

func accessInventory() {
	for {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		gokemon.AfficherTitre()

		largeur := 50
		fmt.Println(gokemon.Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		gokemon.AfficherLigneMenu("", largeur)
		gokemon.AfficherLigneMenu("        INVENTAIRE", largeur)
		gokemon.AfficherLigneMenu("", largeur)
		fmt.Println(gokemon.Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))

		for i, item := range joueur.Inventaire {
			gokemon.AfficherLigneMenu(fmt.Sprintf("%d. %s (x%d)", i+1, item.Nom, item.Quantite), largeur)
		}
		gokemon.AfficherLigneMenu(fmt.Sprintf("%d. Retour au menu principal", len(joueur.Inventaire)+1), largeur)
		gokemon.AfficherLigneMenu("", largeur)
		fmt.Println(gokemon.Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(gokemon.Vert("\nEntrez votre choix : "))
		var choix int
		fmt.Scanln(&choix)

		if choix == len(joueur.Inventaire)+1 {
			return
		} else if choix > 0 && choix <= len(joueur.Inventaire) {
			item := &joueur.Inventaire[choix-1]
			if item.Nom == "Potion" {
				takePot(item)
			} else {
				fmt.Printf(gokemon.Jaune("\nVous ne pouvez pas utiliser %s pour le moment.\n"), item.Nom)
			}
		} else {
			fmt.Println(gokemon.Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(gokemon.Vert("\nAppuyez sur Entrée pour continuer..."))
		fmt.Scanln()
	}
}

func takePot(potion *gokemon.Item) {
	if potion.Quantite > 0 {
		if len(joueur.Equipe) > 0 {
			pokemon := &joueur.Equipe[0]
			if pokemon.PVActuels < pokemon.PVMax {
				pokemon.PVActuels += 50
				if pokemon.PVActuels > pokemon.PVMax {
					pokemon.PVActuels = pokemon.PVMax
				}
				potion.Quantite--
				fmt.Printf(gokemon.Jaune("\nVous avez utilisé une Potion. %s a récupéré 50 PV. PV actuels : %d/%d\n"), pokemon.Nom, pokemon.PVActuels, pokemon.PVMax)
			} else {
				fmt.Println(gokemon.Jaune("\nVotre Pokémon a déjà tous ses PV."))
			}
		} else {
			fmt.Println(gokemon.Jaune("\nVous n'avez pas de Pokémon dans votre équipe."))
		}
	} else {
		fmt.Println(gokemon.Jaune("\nVous n'avez plus de Potions."))
	}
}

func estVivant(p *gokemon.Pokemon) bool {
	return p.PVActuels > 0
}

func attaquer(p *gokemon.Pokemon, cible *gokemon.Pokemon) {
	degats := rand.Intn(p.Attaque) + 1
	cible.PVActuels -= degats
	if cible.PVActuels < 0 {
		cible.PVActuels = 0
	}
	fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", p.Nom, cible.Nom, degats)
	fmt.Printf("%s a maintenant %d PV\n", cible.Nom, cible.PVActuels)
}

func combat() {
	if len(joueur.Equipe) == 0 {
		fmt.Println(gokemon.Jaune("\nVous n'avez pas de Pokémon pour combattre. Créez d'abord votre dresseur."))
		return
	}

	pokemonJoueur := &joueur.Equipe[0]
	ennemi := &gokemon.Pokemon{Nom: "Roucool", Type: "Vol", Niveau: 5, PVMax: 40, PVActuels: 40, Attaque: 8}

	fmt.Println(gokemon.Jaune("\nUn Roucool sauvage apparaît!"))

	for pokemonJoueur.EstVivant() && ennemi.EstVivant() {
		fmt.Println(gokemon.Jaune("\nQue voulez-vous faire ?"))
		fmt.Println(gokemon.Jaune("1. Attaquer"))
		fmt.Println(gokemon.Jaune("2. Utiliser une Potion"))
		fmt.Println(gokemon.Jaune("3. Fuir"))

		var choix int
		fmt.Print(gokemon.Vert("Entrez votre choix (1-3) : "))
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			pokemonJoueur.Attaquer(ennemi)
			if ennemi.EstVivant() {
				ennemi.Attaquer(pokemonJoueur)
			}
		case 2:
			for i, item := range joueur.Inventaire {
				if item.Nom == "Potion" {
					takePot(&joueur.Inventaire[i])
					ennemi.Attaquer(pokemonJoueur)
					break
				}
			}
		case 3:
			fmt.Println(gokemon.Jaune("Vous avez fui le combat!"))
			return
		default:
			fmt.Println(gokemon.Jaune("Choix invalide, vous perdez votre tour!"))
			ennemi.Attaquer(pokemonJoueur)
		}
	}

	if pokemonJoueur.EstVivant() {
		fmt.Println(gokemon.Jaune("Félicitations! Vous avez vaincu le Roucool sauvage!"))
	} else {
		fmt.Println(gokemon.Jaune("Votre Pokémon est K.O. ! Vous avez perdu le combat."))
	}
}

func menuPrincipal() {
	for {
		largeur := 50
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		gokemon.AfficherTitre()

		fmt.Println(gokemon.Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		gokemon.AfficherLigneMenu("", largeur)
		gokemon.AfficherLigneMenu("        MENU PRINCIPAL", largeur)
		gokemon.AfficherLigneMenu("", largeur)
		fmt.Println(gokemon.Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		gokemon.AfficherLigneMenu("1. Créer Dresseur", largeur)
		gokemon.AfficherLigneMenu("2. Afficher les informations du dresseur", largeur)
		gokemon.AfficherLigneMenu("3. Accéder à l'inventaire", largeur)
		gokemon.AfficherLigneMenu("4. Combat", largeur)
		gokemon.AfficherLigneMenu("5. Quitter", largeur)
		gokemon.AfficherLigneMenu("", largeur)
		fmt.Println(gokemon.Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(gokemon.Vert("\nEntrez votre choix (1-5): "))
		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			creerDresseur()
		case "2":
			if joueur.Nom == "" {
				fmt.Println(gokemon.Jaune("\nVeuillez d'abord créer votre dresseur."))
			} else {
				gokemon.DisplayInfo(joueur)
			}
		case "3":
			accessInventory()
		case "4":
			combat()
		case "5":
			fmt.Println(gokemon.Jaune("\nMerci d'avoir joué. Au revoir!"))
			os.Exit(0)
		default:
			fmt.Println(gokemon.Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(gokemon.Vert("\nAppuyez sur Entrée pour continuer..."))
		fmt.Scanln()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	menuPrincipal()
}
