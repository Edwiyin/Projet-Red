package gokemon

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func afficherBarrePV(pokemon Pokemon) string {
	barreLength := 20
	pvRatio := float64(pokemon.PVActuels) / float64(pokemon.PVMax)
	pvActuels := int(pvRatio * float64(barreLength))

	barre := strings.Repeat("█", pvActuels) + strings.Repeat("░", barreLength-pvActuels)
	return fmt.Sprintf("%s [%s] %d/%d", pokemon.Nom, barre, pokemon.PVActuels, pokemon.PVMax)
}

func (p *Pokemon) Attaquer(cible *Pokemon) int {
	damage := rand.Intn(10) + 1
	cible.PVActuels -= damage
	if cible.PVActuels < 0 {
		cible.PVActuels = 0
	}
	return damage
}

func (p *Pokemon) AttaqueSpec(cible *Pokemon) int {
	damage := SkillDamage[p.Type][cible.Type]
	cible.PVActuels -= damage
	if cible.PVActuels < 0 {
		cible.PVActuels = 0
	}
	return damage
}

func (p *Pokemon) EstVivant() bool {
	return p.PVActuels > 0
}

type PokemonType string

const ()

type Item struct {
	Nom      string
	Quantite int
}

func Combat(joueur *Dresseur, isTraining bool) {
	if len(joueur.Equipe) == 0 {
		fmt.Println(Jaune("\nVous n'avez pas de Pokémon pour combattre. Créez d'abord votre dresseur."))
		return
	}
	audioManager.StopMusic()
	audioManager.PlayBattleMusic()
	pokemonJoueur := ChoisirPokemon(joueur)
	ennemi := GenerateWildPokemon(joueur)

	printBattleGround(PokeArt[pokemonJoueur.Nom], PokeArt[ennemi.Nom])
	fmt.Printf(Jaune("\nUn %s sauvage de niveau %d apparaît!\n"), ennemi.Nom, ennemi.Niveau)
	fmt.Printf(Jaune("Vous envoyez %s au combat!\n"), pokemonJoueur.Nom)

	for pokemonJoueur.EstVivant() && ennemi.EstVivant() {
		fmt.Println(Jaune("\nQue voulez-vous faire ?"))
		fmt.Println(Jaune("1. Attaque Normale"))
		fmt.Println(Jaune("2. Attaque Spéciale"))
		fmt.Println(Jaune("3. Utiliser une Potion"))
		fmt.Println(Jaune("4. Lancer une Pokéball"))
		fmt.Println(Jaune("5. Changer de Pokémon"))
		fmt.Println(Jaune("6. Fuir"))

		state := true
		var choix int
		Wrap(func() { fmt.Scan(&choix) })

		switch choix {
		case 1:
			fmt.Println(Jaune("\nChoisissez une attaque :"))
			fmt.Println(Jaune("1. Coup de Poing"))
			fmt.Println(Jaune("2. Griffe"))
			fmt.Println(Jaune("3. Charge"))
			fmt.Println(Jaune("4. Morsure"))
			fmt.Println(Jaune("5. Retour"))

			fmt.Scan(&choix)
			switch choix {
			case 1:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 2:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 3:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 4:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 5:
				fmt.Println(Jaune("Retour au combat..."))
				state = false
				time.Sleep(1 * time.Second)

			default:
				fmt.Println(Jaune("Choix invalide."))

			}
			if ennemi.EstVivant() && state {
				damage := ennemi.Attaquer(pokemonJoueur)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", ennemi.Nom, pokemonJoueur.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(*pokemonJoueur)))
			}
		case 2:
			fmt.Println(Jaune("\nChoisissez une attaque spéciale :"))
			fmt.Println(Jaune(fmt.Sprintf("1. %s", SkillName[pokemonJoueur.Type][0])))
			fmt.Println(Jaune(fmt.Sprintf("2. %s", SkillName[pokemonJoueur.Type][1])))
			fmt.Println(Jaune(fmt.Sprintf("3. %s", SkillName[pokemonJoueur.Type][2])))
			fmt.Println(Jaune(fmt.Sprintf("4. %s", SkillName[pokemonJoueur.Type][3])))
			fmt.Println(Jaune("5. Retour"))

			fmt.Scan(&choix)
			switch choix {
			case 1:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 2:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 3:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 4:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Rouge(afficherBarrePV(ennemi)))
				state = true
			case 5:
				fmt.Println(Jaune("Retour au combat..."))
				time.Sleep(1 * time.Second)
				state = false

			default:
				fmt.Println(Jaune("Choix invalide."))
			}
			if ennemi.EstVivant() && state {
				damage := ennemi.Attaquer(pokemonJoueur)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", ennemi.Nom, pokemonJoueur.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(*pokemonJoueur)))
			}

		case 3:

			fmt.Println(Jaune("\nChoisissez une potion :"))
			fmt.Println(Jaune("1. Potion de Soin"))
			fmt.Println(Jaune("2. Potion de Poison"))
			fmt.Println(Jaune("3. Retour"))

			fmt.Scan(&choix)
			switch choix {
			case 1:
				UsePotion(joueur, pokemonJoueur)
				state = true
			case 2:
				if UsePoisonPotion(joueur, &ennemi) {
				}
				state = true
			case 3:
				fmt.Println(Jaune("Retour au combat..."))
				time.Sleep(1 * time.Second)
				state = false
			default:
				fmt.Println(Jaune("Choix invalide."))
				state = false
			}

		case 4:
			if TryToCatch(joueur, &ennemi) {
				return
			}
		case 5:
			nouveauPokemon := ChoisirPokemon(joueur)
			if nouveauPokemon != pokemonJoueur {
				pokemonJoueur = nouveauPokemon
				fmt.Printf(Jaune("\nVous rappelez %s et envoyez %s au combat!\n"), pokemonJoueur.Nom, nouveauPokemon.Nom)
				damage := ennemi.Attaquer(pokemonJoueur)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", ennemi.Nom, pokemonJoueur.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(*pokemonJoueur)))
			} else {
				fmt.Println(Jaune("\nVous avez choisi le même Pokémon. Le combat continue."))
				time.Sleep(3 * time.Second)
			}
		case 6:
			fmt.Println(Jaune("Vous avez fui le combat!"))
			time.Sleep(3 * time.Second)
			return
		default:
			fmt.Println(Jaune("Choix invalide."))
		}
		if !pokemonJoueur.EstVivant() {
			nouveauPokemon := ChoisirPokemonVivant(joueur)
			if nouveauPokemon != nil {
				pokemonJoueur = nouveauPokemon
				fmt.Printf(Jaune("\nVotre Pokemon est K.O. ! Vous envoyez %s au combat!\n"), nouveauPokemon.Nom)
			} else {
				if Dead(joueur) {
					fmt.Println(Jaune("\nVous avez perdu le combat..."))
					return
				}
				pokemonJoueur = &joueur.Equipe[0]
				fmt.Printf(Jaune("\nVous continuez le combat avec %s!\n"), pokemonJoueur.Nom)
			}
		}

	}

	if pokemonJoueur.EstVivant() {
		expGained := ennemi.Niveau * 20
		moneyGained := ennemi.Niveau * 50
		fmt.Printf(Jaune("\nVous avez gagné le combat! %s gagne %d points d'expérience.\n"), pokemonJoueur.Nom, expGained)
		fmt.Printf(Jaune("Vous avez gagné %d PokéDollars!\n"), moneyGained)
		joueur.Argent += moneyGained
		time.Sleep(3 * time.Second)
		if !isTraining {
			if pokemonJoueur.GainExperience(expGained) {
				fmt.Printf(Jaune("%s passe au niveau %d!\n"), pokemonJoueur.Nom, pokemonJoueur.Niveau)
			}
			ressourceWon := TypeToResource[ennemi.Type]
			quantiteWon := rand.Intn(3) + 1

			joueur.AddResource(ressourceWon, quantiteWon)
			fmt.Printf(Jaune("Vous avez gagné %d %s!\n"), quantiteWon, ressourceWon)
		}
		time.Sleep(3 * time.Second)

	} else {
		fmt.Println(Jaune("Vous avez perdu le combat..."))
		time.Sleep(4 * time.Second)
	}
}
func printBattleGround(pokemonJoueurAscii string, pokemonEnnemiAscii string) {
	splittedPJA := strings.Split(pokemonJoueurAscii, "\n")
	splittedPEA := strings.Split(pokemonEnnemiAscii, "\n")
	maxSize := max(len(splittedPJA), len(splittedPEA))
	minSize := min(len(splittedPJA), len(splittedPEA))
	filler := make([]string, maxSize-minSize)
	for i := 0; i < cap(filler); i++ {
		filler[i] = ""
	}
	if len(splittedPJA) < len(splittedPEA) {
		splittedPJA = append(filler, splittedPJA...)
	} else if len(splittedPJA) > len(splittedPEA) {
		splittedPEA = append(filler, splittedPEA...)
	}
	for i := 0; i < maxSize; i++ {
		fmt.Printf("%-40s%20s%-40s\n", splittedPJA[i], "", splittedPEA[i])
	}
}
func TryToCatch(joueur *Dresseur, pokemon *Pokemon) bool {
	for i, item := range joueur.Inventaire {
		if item.Nom == "Pokéball" {
			if item.Quantite > 0 {
				joueur.Inventaire[i].Quantite--
				catchChance := float64(pokemon.PVMax-pokemon.PVActuels) / float64(pokemon.PVMax)
				fmt.Printf(Jaune("\nVous lancez une Pokéball... Chance de capture : %.2f\n"), catchChance)
				if rand.Float64() < catchChance {
					joueur.Equipe = append(joueur.Equipe, *pokemon)
					fmt.Printf(Jaune("\nFélicitations! Vous avez capturé %s!\n"), pokemon.Nom)
					time.Sleep(4 * time.Second)
					return true
				} else {
					fmt.Println(Jaune("\nLe Pokémon s'est échappé de la Pokéball!"))
				}
			} else {
				fmt.Println(Jaune("\nVous n'avez plus de Pokéballs."))
			}
			return false
		}
	}
	fmt.Println(Jaune("\nVous n'avez pas de Pokéballs dans votre inventaire."))
	return false
}

func ChoisirPokemon(joueur *Dresseur) *Pokemon {
	if len(joueur.Equipe) == 1 {
		return &joueur.Equipe[0]
	}

	fmt.Println(Jaune("\nChoisissez votre Pokémon pour ce combat :"))
	for i, pokemon := range joueur.Equipe {
		fmt.Printf(Jaune("%d. %s (Niveau: %d, PV: %d/%d)\n"), i+1, pokemon.Nom, pokemon.Niveau, pokemon.PVActuels, pokemon.PVMax)
	}

	var choix int
	for {
		fmt.Print(Vert("\nEntrez le numéro du Pokémon que vous voulez utiliser : "))
		Wrap(func() { fmt.Scanln(&choix) })
		if choix > 0 && choix <= len(joueur.Equipe) {
			return &joueur.Equipe[choix-1]
		}
		fmt.Println(Jaune("Choix invalide. Veuillez réessayer."))
	}
}

func ChoisirPokemonVivant(joueur *Dresseur) *Pokemon {
	for i := range joueur.Equipe {
		if joueur.Equipe[i].EstVivant() {
			return &joueur.Equipe[i]
		}
	}
	return nil
}
func Dead(joueur *Dresseur) bool {
	allDead := true
	for i := range joueur.Equipe {
		if joueur.Equipe[i].EstVivant() {
			allDead = false
			break
		}
	}

	if allDead {
		fmt.Println(Jaune("\nTous vos Pokémon sont K.O. ! Vous êtes transporté au centre Pokémon le plus proche..."))
		time.Sleep(4 * time.Second)

		for i := range joueur.Equipe {
			joueur.Equipe[i].PVActuels = joueur.Equipe[i].PVMax / 2
			fmt.Printf(Jaune("%s a été ressuscité avec %d PV.\n"), joueur.Equipe[i].Nom, joueur.Equipe[i].PVActuels)
		}

		fmt.Println(Jaune("\nVos Pokémon ont été soignés. Prenez soin d'eux !"))
		time.Sleep(4 * time.Second)
		return true
	}

	return false
}
