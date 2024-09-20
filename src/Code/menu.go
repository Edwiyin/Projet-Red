package gokemon

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var audioManager *AudioManager

func (joueur *Dresseur) charCreation() {
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
			nom = Capitalize(nom)
			joueur.Nom = nom
		} else if len(nom) == 0 {
			nomValide = false
			fmt.Println(Jaune("Le nom ne peut pas être vide. Veuillez réessayer."))
		}
	}

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
}

func (joueur *Dresseur) createCharacter() {
	joueur.CapaciteInventaire = 10
	if joueur.Nom == "" {
		joueur.charCreation()
	} else {
		fmt.Println(Jaune("Vous avez déjà créé votre dresseur."))
	}

}

func (joueur *Dresseur) MenuPrincipal(newAudioManager *AudioManager) {
	audioManager = newAudioManager
	audioManager.PlayBackgroundMusic()
	largeur := 155
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	AfficherTitre()

	fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
	AfficherLigneMenu("", largeur)
	AfficherLigneMenu("                                                                  NEW GAME", largeur)
	AfficherLigneMenu("", largeur)
	fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

	MessageRapide(("Appuyez sur Entrée pour commencer..."), 3, "vert")
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

		fmt.Print(Vert("\nEntrez votre choix (1-9): "))
		var choix string
		Wrap(func() { fmt.Scanln(&choix) })

		switch choix {
		case "1":
			joueur.createCharacter()
		case "2":
			if joueur.Nom == "" {
				fmt.Println(Jaune("\nVeuillez d'abord créer votre dresseur."))
			} else {
			DisplayInfo(joueur)
			}
		case "3":
			AfficherEquipements(joueur)
		case "4":
			joueur.AccessInventory()
		case "5":
			Combat(joueur, audioManager)
			audioManager.StopMusic()
			audioManager.PlayBackgroundMusic()

		case "6":
			VisiteMarchand(joueur)
		case "7":
			VisiterForgeron(joueur)
		case "8":
			MessageRapide(("Abba"), 3, "bleu")
			MessageRapide(("Steven Spielberg"), 3, "bleu")
			MessageRapide(("Les développeurs de ce jeu sont: Massinissa Ahfir, Edwin Wehbe, Michel Mustafaov"), 3, "bleu")
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

func Capitalize(s string) string {
	slice := []rune(s)
	if slice[0] >= 'a' && slice[0] <= 'z' {
		slice[0] = slice[0] - 32
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			if (slice[i-1] >= 'a' && slice[i-1] <= 'z') || (slice[i-1] >= 'A' && slice[i-1] <= 'Z') || (slice[i-1] >= '0' && slice[i-1] <= '9') {
				slice[i] = slice[i] + 32
			}
		} else if slice[i] >= 'a' && slice[i] <= 'z' {
			if (slice[i-1] < 'a' || slice[i-1] > 'z') && (slice[i-1] < 'A' || slice[i-1] > 'Z') && (slice[i-1] < '0' || slice[i-1] > '9') {
				slice[i] = slice[i] - 32
			}
		}
	}
	return string(slice)
}
