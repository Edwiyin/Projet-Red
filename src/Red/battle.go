package gokemon

import (
	"fmt"
	"math/rand"
	"time"
)

func poisonPot(pokemon *Pokemon) {
	fmt.Printf(Jaune("\nLa Potion de Poison affecte %s !\n"), pokemon.Nom)
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		pokemon.PVActuels -= 10
		if pokemon.PVActuels < 0 {
			pokemon.PVActuels = 0
		}
		fmt.Printf(Jaune("%s subit 10 points de dégâts. PV actuels : %d/%d\n"), pokemon.Nom, pokemon.PVActuels, pokemon.PVMax)
		if pokemon.PVActuels == 0 {
			fmt.Printf(Jaune("%s est K.O. !\n"), pokemon.Nom)
			break
		}
	}
}
func (p *Pokemon) GainExperience(exp int) bool {
	p.Experience += exp
	if p.Experience >= p.Niveau*100 {
		p.Experience = 0
		p.Niveau++
		return true
	}
	return false
}

func (p *Pokemon) Attaquer(cible *Pokemon) int {
	damage := rand.Intn(10) + 1
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
	ennemi := GenerateWildPokemon()

	fmt.Printf(Jaune("\nUn %s sauvage de niveau %d apparaît!\n"), ennemi.Nom, ennemi.Niveau)
	fmt.Printf(Jaune("Vous envoyez %s au combat!\n"), pokemonJoueur.Nom)

	for pokemonJoueur.EstVivant() && ennemi.EstVivant() {
		fmt.Println(Jaune("\nQue voulez-vous faire ?"))
		fmt.Println(Jaune("1. Attaquer"))
		fmt.Println(Jaune("2. Utiliser une Potion"))
		fmt.Println(Jaune("3. Lancer une Pokéball"))
		fmt.Println(Jaune("4. Utiliser une Potion de Poison"))
		fmt.Println(Jaune("5. Changer de Pokémon"))
		fmt.Println(Jaune("6. Fuir"))

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			damage := pokemonJoueur.Attaquer(&ennemi)
			fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", pokemonJoueur.Nom, ennemi.Nom, damage)
			fmt.Printf("%s a maintenant %d PV\n", ennemi.Nom, ennemi.PVActuels)
			if ennemi.EstVivant() {
				damage := ennemi.Attaquer(pokemonJoueur)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", ennemi.Nom, pokemonJoueur.Nom, damage)
				fmt.Printf("%s a maintenant %d PV\n", pokemonJoueur.Nom, pokemonJoueur.PVActuels)
			}
		case 2:
			UsePotion(joueur, pokemonJoueur)
			if ennemi.EstVivant() {
				damage := ennemi.Attaquer(pokemonJoueur)
				fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", ennemi.Nom, pokemonJoueur.Nom, damage)
				fmt.Printf("%s a maintenant %d PV\n", pokemonJoueur.Nom, pokemonJoueur.PVActuels)
			}
		case 3:
			if TryToCatch(joueur, &ennemi) {
				return
			}
		case 4:
			if UsePoisonPotion(joueur, &ennemi) {
				if ennemi.EstVivant() {
					damage := ennemi.Attaquer(pokemonJoueur)
					fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", ennemi.Nom, pokemonJoueur.Nom, damage)
					fmt.Printf("%s a maintenant %d PV\n", pokemonJoueur.Nom, pokemonJoueur.PVActuels)
				}
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
			}
		case 6:
			fmt.Println(Jaune("Vous avez fui le combat!"))
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

func GenerateWildPokemon() Pokemon {
	wildPokemons := []struct {
		name        string
		pokemonType PokemonType
	}{
		{"Rattata", Normal},
		{"Pidgey", Flying},
		{"Caterpie", Bug},
		{"Weedle", Bug},
		{"Pikachu", Electric},
		{"Eevee", Normal},
		{"Vulpix", Fire},
		{"Jigglypuff", Normal},
		{"Zubat", Flying},
		{"Oddish", Grass},
		{"Paras", Bug},
		{"Venonat", Bug},
		{"Meowth", Normal},
		{"Psyduck", Water},
		{"Mankey", Normal},
		{"Growlithe", Fire},
		{"Poliwag", Water},
		{"Horsea", Water},
		{"Goldeen", Water},
		{"Staryu", Water},
		{"Scyther", Bug},
		{"Electabuzz", Electric},
		{"Magmar", Fire},
		{"Ronflex", Normal},
		{"Dracaufeu", Fire},
		{"Tortank", Water},
		{"Florizarre", Grass},
	}

	randomPokemon := wildPokemons[rand.Intn(len(wildPokemons))]
	level := rand.Intn(3) + 1
	return Pokemon{
		Nom:        randomPokemon.name,
		PVActuels:  level * 10,
		PVMax:      level * 10,
		Niveau:     level,
		Type:       randomPokemon.pokemonType,
		Experience: 0,
	}
}

func UsePotion(joueur *Dresseur, pokemon *Pokemon) {
	for i, item := range joueur.Inventaire {
		if item.Nom == "Potion" {
			if item.Quantite > 0 {
				healAmount := 20
				pokemon.PVActuels += healAmount
				if pokemon.PVActuels > pokemon.PVMax {
					pokemon.PVActuels = pokemon.PVMax
				}
				joueur.Inventaire[i].Quantite--
				fmt.Printf(Jaune("\nVous avez utilisé une Potion. %s a récupéré %d PV. PV actuels : %d/%d\n"), pokemon.Nom, healAmount, pokemon.PVActuels, pokemon.PVMax)
			} else {
				fmt.Println(Jaune("\nVous n'avez plus de Potions."))
			}
			return
		}
	}
	fmt.Println(Jaune("\nVous n'avez pas de Potions dans votre inventaire."))
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
