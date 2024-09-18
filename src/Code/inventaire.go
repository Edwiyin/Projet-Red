package gokemon

import (
	"fmt"
	"strings"
)

const LimiteInv = 10

func TakePot(item *Item, joueur *Dresseur) {
    if item.Quantite > 0 {
        if len(joueur.Equipe) == 1 {
            healPokemon(&joueur.Equipe[0], item)
        } else {
            fmt.Println(Jaune("\nChoisissez le Pokémon à soigner :"))
            for i, pokemon := range joueur.Equipe {
                fmt.Printf(Jaune("%d. %s (PV: %d/%d)\n"), i+1, pokemon.Nom, pokemon.PVActuels, pokemon.PVMax)
            }
            
            var choix int
            fmt.Print(Vert("\nEntrez votre choix : "))
            Wrap(func() { fmt.Scanln(&choix) })
            
            if choix > 0 && choix <= len(joueur.Equipe) {
                healPokemon(&joueur.Equipe[choix-1], item)
            } else {
                fmt.Println(Jaune("\nChoix invalide. Aucun Pokémon n'a été soigné."))
            }
        }
    } else {
        fmt.Println(Jaune("\nVous n'avez plus de Potions."))
    }
}

func healPokemon(pokemon *Pokemon, item *Item) {
    if pokemon.PVActuels < pokemon.PVMax {
        pokemon.PVActuels += 20
        if pokemon.PVActuels > pokemon.PVMax {
            pokemon.PVActuels = pokemon.PVMax
        }
        item.Quantite--
        fmt.Printf(Jaune("\nVous avez utilisé une Potion sur %s. PV actuels: %d/%d\n"), pokemon.Nom, pokemon.PVActuels, pokemon.PVMax)
    } else {
        fmt.Printf(Jaune("\n%s a déjà ses PV au maximum.\n"), pokemon.Nom)
    }
}

func AccessInventory(joueur *Dresseur) {
	for {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		largeur := 155
		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("                                                                  INVENTAIRE", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))

		AfficherLigneMenu(fmt.Sprintf("                                                      4  Porte-Monnaie: %d PokéDollars", joueur.Argent), largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))

		for i, item := range joueur.Inventaire {
			AfficherLigneMenu(fmt.Sprintf("%d. %s (x%d)", i+1, item.Nom, item.Quantite), largeur)
		}
		AfficherLigneMenu(fmt.Sprintf("%d. Retour au menu principal", len(joueur.Inventaire)+1), largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix : "))
		var choix int
		Wrap(func() { fmt.Scanln(&choix) })

		if choix == len(joueur.Inventaire)+1 {
			return
		} else if choix > 0 && choix <= len(joueur.Inventaire) {
			item := &joueur.Inventaire[choix-1]
			if item.Nom == "Potion" {
				TakePot((*Item)(item), joueur)
			} else if item.Nom == "Casque" || item.Nom == "Armure" || item.Nom == "Bottes" {
				EquiperObjet(joueur, item)
			} else {
				fmt.Printf(Jaune("\nVous ne pouvez pas utiliser %s pour le moment.\n"), item.Nom)
			}
		} else {
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		Wrap(func() { fmt.Scanln() })
	}
}

func EquiperObjet(joueur *Dresseur, item *InventoryItem) {
	var emplacement string
	switch item.Nom {
	case "Casque":
		emplacement = "Tête"
	case "Armure":
		emplacement = "Torse"
	case "Bottes":
		emplacement = "Pieds"
	}

	equipement := Equipment{Nom: item.Nom, Emplacement: emplacement, BonusPV: 10}
	joueur.EquiperEquipement(equipement)
	fmt.Printf(Jaune("\nVous avez équipé %s.\n"), item.Nom)
}