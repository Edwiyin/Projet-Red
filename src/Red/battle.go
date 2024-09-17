package gokemon

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func afficherBarrePV(pvActuels, pvMax int) string {
	barLength := 20
	filledLength := int(float64(pvActuels) / float64(pvMax) * float64(barLength))
	bar := strings.Repeat("█", filledLength) + strings.Repeat("░", barLength-filledLength)
	return fmt.Sprintf("[%s] %d/%d", bar, pvActuels, pvMax)
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

func Combat(joueur *Dresseur) {
	if len(joueur.Equipe) == 0 {
		fmt.Println(Jaune("\nVous n'avez pas de Pokémon pour combattre. Créez d'abord votre dresseur."))
		return
	}
	pokemonJoueur := ChoisirPokemon(joueur)
	ennemi := GenerateWildPokemon(joueur)

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
			fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
			fmt.Println(Jaune("5. Retour"))

			fmt.Scan(&choix)
			switch choix {
			case 1:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
				state = true
			case 2:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
				state = true
			case 3:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
				state = true
			case 4:
				damage := pokemonJoueur.Attaquer(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
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
				fmt.Printf("%s a maintenant %d PV\n", pokemonJoueur.Nom, pokemonJoueur.PVActuels)
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
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
				state = true
			case 2:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
				state = true
			case 3:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
				state = true
			case 4:
				damage := pokemonJoueur.AttaqueSpec(&ennemi)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
				fmt.Println(Jaune(afficherBarrePV(ennemi.PVActuels, ennemi.PVMax)))
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
				fmt.Printf("%s a maintenant %d PV\n", pokemonJoueur.Nom, pokemonJoueur.PVActuels)
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
				fmt.Printf("%s a maintenant %d PV\n", pokemonJoueur.Nom, pokemonJoueur.PVActuels)
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
				pokemonJoueur = &joueur.Equipe[0] // On utilise le premier Pokémon ressuscité
				fmt.Printf(Jaune("\nVous continuez le combat avec %s!\n"), pokemonJoueur.Nom)
			}
		}

	}
	if pokemonJoueur.EstVivant() {
		expGained := ennemi.Niveau * 10
		moneyGained := ennemi.Niveau * 50
		fmt.Printf(Jaune("\nVous avez gagné le combat! %s gagne %d points d'expérience.\n"), pokemonJoueur.Nom, expGained)
		fmt.Printf(Jaune("Vous avez gagné %d PokéDollars!\n"), moneyGained)
		joueur.Argent += moneyGained
		time.Sleep(6 * time.Second)
		if pokemonJoueur.GainExperience(expGained) {
			fmt.Printf(Jaune("%s passe au niveau %d!\n"), pokemonJoueur.Nom, pokemonJoueur.Niveau)
		}
	} else {
		fmt.Println(Jaune("Vous avez perdu le combat..."))
		time.Sleep(5 * time.Second)
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
					time.Sleep(5 * time.Second)
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
func poisonPot(enemy *Pokemon) {
	poisonDamage := 7
	enemy.PVActuels -= poisonDamage
	if enemy.PVActuels < 0 {
		enemy.PVActuels = 0
	}
	fmt.Printf(Jaune("\n%s a été empoisonné et perd %d PV. PV actuels : %d/%d\n"), enemy.Nom, poisonDamage, enemy.PVActuels, enemy.PVMax)
}

func UsePoisonPotion(joueur *Dresseur, pokemon *Pokemon) bool {
	for i, item := range joueur.Inventaire {
		if item.Nom == "Potion de Poison" {
			if item.Quantite > 0 {
				joueur.Inventaire[i].Quantite--
				poisonPot(pokemon)
				return true
			} else {
				fmt.Println(Jaune("\nVous n'avez plus de Potions de Poison."))
				return false
			}
		}
	}
	fmt.Println(Jaune("\nVous n'avez pas de Potions de Poison dans votre inventaire."))
	return false
}

func UsePotion(joueur *Dresseur, pokemon *Pokemon) {
	for i, item := range joueur.Inventaire {
		if item.Nom == "Potion" {
			if item.Quantite > 0 {
				healAmount := 20
				pokemon.PVActuels += healAmount
				if pokemon.PVActuels >= pokemon.PVMax {
					if pokemon.PVActuels == pokemon.PVMax {
						fmt.Printf(Jaune("\nVotre Pokémon est déjà en bonne santé \n: %d/%d"), pokemon.Nom, healAmount, pokemon.PVActuels, pokemon.PVMax)
					} else {
						joueur.Inventaire[i].Quantite--
					}
					fmt.Printf(Jaune("\nVous avez utilisé une Potion. %s a récupéré %d PV. PV actuels : %d/%d\n"), pokemon.Nom, healAmount, pokemon.PVActuels, pokemon.PVMax)
					time.Sleep(3 * time.Second)
				} else {
					fmt.Println(Jaune("\nVous n'avez plus de Potions."))
				}
				return
			}
		}
		fmt.Println(Jaune("\nVous n'avez pas de Potions dans votre inventaire."))
	}
}
