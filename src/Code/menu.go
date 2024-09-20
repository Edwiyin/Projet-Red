package gokemon

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

var audioManager *AudioManager

func charCreation() *Dresseur {
	var nom string
	var nomValide bool

	for !nomValide {
		fmt.Print(Vert("Entrez votre nom de dresseur (lettres uniquement) : "))
		Wrap(func() { fmt.Scanln(&nom) })

		nomValide = true
		for _, char := range nom {
			if !unicode.IsLetter(char) {
				nomValide = false
				fmt.Println(Jaune("Le nom doit contenir uniquement des lettres. Veuillez réessayer."))
				break
			}
		}

		if nomValide && len(nom) > 0 {
			nom = strings.ToLower(nom)
			nom = strings.Title(nom)
		} else if len(nom) == 0 {
			nomValide = false
			fmt.Println(Jaune("Le nom ne peut pas être vide. Veuillez réessayer."))
		}
	}

	joueur := &Dresseur{Nom: nom}

	fmt.Print(Vert("Entrez votre choix (1-3) : "))
	fmt.Println(Jaune("\nChoisissez votre Pokémon de départ :"))
	fmt.Println(Jaune("1. Bulbizarre (Type: Plante)"))
	fmt.Println(Jaune("2. Salamèche (Type: Feu)"))
	fmt.Println(Jaune("3. Carapuce (Type: Eau)"))
	var choixPokemon string
	Wrap(func() { fmt.Scanln(&choixPokemon) })

	pokemon := choixPokemonFunc(choixPokemon)
	joueur.Equipe = append(joueur.Equipe, *pokemon)
	joueur.Argent = 100
	fmt.Printf(Jaune("Félicitations, %s ! Vous avez choisi %s comme Pokémon de départ!\n"), joueur.Nom, pokemon.Nom)
	InitialiserCapaciteInventaire(joueur)
	return joueur
}

func createCharacter(joueur *Dresseur) {
	if joueur.Nom == "" {
		*joueur = *charCreation()
	} else {
		fmt.Println(Jaune("Vous avez déjà créé votre dresseur."))
	}
}

func MenuPrincipal(joueur *Dresseur, newAudioManager *AudioManager) {
	audioManager = newAudioManager
	largeur := 155
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	AfficherTitre()

	fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
	AfficherLigneMenu("", largeur)
	AfficherLigneMenu("                                                                  NEW GAME", largeur)
	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))
	fmt.Println(PokeArt["Laggron"])
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
		AfficherLigneMenu("8. Entraînement", largeur)
		AfficherLigneMenu("9. Qui sont-ils", largeur)
		AfficherLigneMenu("10. Quitter le Jeu", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Print(Vert("\nEntrez votre choix (1-10): "))
		var choix string
		Wrap(func() { fmt.Scanln(&choix) })

		switch choix {
		case "1":
			createCharacter(joueur)

		case "2":
			if joueur.Nom == "" {
				fmt.Println(Jaune("\nVeuillez d'abord créer votre dresseur."))
			} else {
				DisplayInfo(joueur)
			}
		case "3":
			AfficherEquipements(joueur)
		case "4":
			AccessInventory(joueur)
		case "5":
			Combat(joueur, false)
			audioManager.StopMusic()
			audioManager.PlayBattleMusic()
		case "6":
			VisiteMarchand(joueur)
		case "7":
			VisiterForgeron(joueur)
		case "8":
			trainigFight(joueur)
			audioManager.StopMusic()
			audioManager.PlayBackgroundMusic()
		case "9":
			MessageRapide(("Abba"), 40, "bleu")
			time.Sleep(1 * time.Second)
			MessageRapide(("Steven Spielberg"), 40, "bleu")
			time.Sleep(1 * time.Second)
			MessageRapide(("Les développeurs de ce jeu sont: Massinissa Ahfir, Edwin Wehbe, Michel Mustafaov"), 40, "bleu")
		case "10":
			fmt.Println(Jaune("\nMerci d'avoir joué. Au revoir!"))
			os.Exit(0)
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		Wrap(func() { fmt.Scanln() })
	}
}
