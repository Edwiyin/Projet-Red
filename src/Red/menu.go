package gokemon

import (
	"fmt"
	"os"
	"strings"
)

const LIMITE_INVENTAIRE = 10

var audioManager *AudioManager

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

func choixPokemonFunc(choixPokemon string) Pokemon {
	switch choixPokemon {
	case "1":
		pokemon := NewPokemon("Bulbizarre", Grass, 1)
        pokemon.PVActuels = pokemon.PVMax / 2
return pokemon
	case "2":
		pokemon := NewPokemon("Salamèche", Fire, 1)
		pokemon.PVActuels = pokemon.PVMax / 2
		return pokemon
	case "3":
		pokemon := NewPokemon("Carapuce", Water, 1)
		pokemon.PVActuels = pokemon.PVMax / 2
		return pokemon
	default:
		fmt.Println(Jaune("Choix invalide. Pokémon par défaut : Pikachu"))
		pokemon := NewPokemon("Pikachu", Electric, 1)
		pokemon.PVActuels = pokemon.PVMax / 2
		return pokemon
	    
	}
}

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

		AfficherLigneMenu(fmt.Sprintf("Solde: %d PokéDollars", joueur.Argent), largeur)
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

func MenuPrincipal(joueur *Dresseur, newAudioManager *AudioManager) {
	audioManager = newAudioManager
	largeur := 50
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	AfficherTitre()

	fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
	AfficherLigneMenu("", largeur)
	AfficherLigneMenu("        NEW GAME", largeur)
	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

	fmt.Print(Vert("\nAppuyez sur Entrée pour commencer..."))
	Wrap(func() { fmt.Scanln() })
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
		AfficherLigneMenu("5. Visiter le marchand", largeur)
		AfficherLigneMenu("6. Quitter", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix (1-6): "))
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
			AccessInventory(joueur)
		case "4":
			Combat(joueur)
		case "5":
			VisiteMarchand(joueur)
		case "6":
			fmt.Println(Jaune("\nMerci d'avoir joué. Au revoir!"))
			os.Exit(0)
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		Wrap(func() { fmt.Scanln() })
	}
}

func VisiteMarchand(joueur *Dresseur) {
	for {
		largeur := 60
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("        BOUTIQUE DU MARCHAND", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("1. Acheter une Potion (50 PokéDollars)", largeur)
		AfficherLigneMenu("2. Acheter une Pokéball (100 PokéDollars)", largeur)
		AfficherLigneMenu("3. Acheter une Potion de Poison (75 PokéDollars)", largeur)
		AfficherLigneMenu("4. Vendre un objet", largeur)
		AfficherLigneMenu("5. Vendre un Pokémon", largeur)
        AfficherLigneMenu("6. Retour au menu principal", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Printf(Jaune("\nVotre solde: %d PokéDollars\n"), joueur.Argent)
		fmt.Print(Vert("\nEntrez votre choix (1-5): "))
		var choix string
		Wrap(func() { fmt.Scanln(&choix) })

		switch choix {
		case "1":
			AcheterObjet(joueur, "Potion", 50)
		case "2":
			AcheterObjet(joueur, "Pokéball", 100)
		case "3":
			AcheterObjet(joueur, "Potion de Poison", 75)
		case "4":
			VendreObjet(joueur)
	    case "5":
			VendrePokemon(joueur)
		case "6":
			return
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		Wrap(func() { fmt.Scanln() })
	}
}

func AcheterObjet(joueur *Dresseur, nomObjet string, prix int) {
    if joueur.Argent >= prix {
        totalItems := 0
        for _, item := range joueur.Inventaire {
            totalItems += item.Quantite
        }
        if totalItems >= LIMITE_INVENTAIRE {
            fmt.Println(Jaune("\nVotre inventaire est plein. Vous ne pouvez pas acheter plus d'objets."))
            return
        }
        
        joueur.Argent -= prix
        for i := range joueur.Inventaire {
            if joueur.Inventaire[i].Nom == nomObjet {
                joueur.Inventaire[i].Quantite++
                fmt.Printf(Jaune("\nVous avez acheté un(e) %s pour %d PokéDollars.\n"), nomObjet, prix)
                return
            }
        }
        joueur.Inventaire = append(joueur.Inventaire, InventoryItem{Nom: nomObjet, Quantite: 1})
        fmt.Printf(Jaune("\nVous avez acheté un(e) %s pour %d PokéDollars.\n"), nomObjet, prix)
    } else {
        fmt.Println(Jaune("\nVous n'avez pas assez d'argent pour acheter cet objet."))
    }
}

func VendreObjet(joueur *Dresseur) {
	fmt.Println(Jaune("\nQuels objets voulez-vous vendre ?"))
	for i, item := range joueur.Inventaire {
		fmt.Printf(Jaune("%d. %s (x%d) - Prix de vente: %d PokéDollars\n"), i+1, item.Nom, item.Quantite, GetPrixVente(item.Nom))
	}
	fmt.Printf(Jaune("%d. Annuler\n"), len(joueur.Inventaire)+1)

	var choix int
	fmt.Print(Vert("\nEntrez votre choix : "))
	Wrap(func() { fmt.Scanln(&choix) })

	if choix > 0 && choix <= len(joueur.Inventaire) {
		item := &joueur.Inventaire[choix-1]
		if item.Quantite > 0 {
			prixVente := GetPrixVente(item.Nom)
			joueur.Argent += prixVente
			item.Quantite--
			fmt.Printf(Jaune("\nVous avez vendu un(e) %s pour %d PokéDollars.\n"), item.Nom, prixVente)
			if item.Quantite == 0 {
				joueur.Inventaire = append(joueur.Inventaire[:choix-1], joueur.Inventaire[choix:]...)
			}
		} else {
			fmt.Println(Jaune("\nVous n'avez plus de cet objet dans votre inventaire."))
		}
	} else if choix != len(joueur.Inventaire)+1 {
		fmt.Println(Jaune("\nChoix invalide."))
	}
}

func VendrePokemon(joueur *Dresseur) {
    if len(joueur.Equipe) <= 1 {
        fmt.Println(Jaune("\nVous ne pouvez pas vendre votre dernier Pokémon !"))
        return
    }

    fmt.Println(Jaune("\nQuels Pokémon voulez-vous vendre ?"))
    for i, pokemon := range joueur.Equipe {
        prix := pokemon.Niveau * 100
        fmt.Printf(Jaune("%d. %s (Niveau %d) - Prix de vente: %d PokéDollars\n"), i+1, pokemon.Nom, pokemon.Niveau, prix)
    }
    fmt.Printf(Jaune("%d. Annuler\n"), len(joueur.Equipe)+1)

    var choix int
    fmt.Print(Vert("\nEntrez votre choix : "))
    Wrap(func() { fmt.Scanln(&choix) })

    if choix > 0 && choix < len(joueur.Equipe)+1 {
        pokemonVendu := joueur.Equipe[choix-1]
        prixVente := pokemonVendu.Niveau * 100
        joueur.Argent += prixVente
        joueur.Equipe = append(joueur.Equipe[:choix-1], joueur.Equipe[choix:]...)
        fmt.Printf(Jaune("\nVous avez vendu %s pour %d PokéDollars.\n"), pokemonVendu.Nom, prixVente)
    } else if choix != len(joueur.Equipe)+1 {
        fmt.Println(Jaune("\nChoix invalide."))
    }
}

func GetPrixVente(nomObjet string) int {
	switch nomObjet {
	case "Potion":
		return 25
	case "Pokéball":
		return 50
	case "Potion de Poison":
		return 35
	default:
		return 10
	}
}
