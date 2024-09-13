package gokemon

import (
	"fmt"
	"math/rand"
)

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

const (
	
)


type Item struct {
	Nom      string
	Quantite int
}


func Combat(joueur *Dresseur) {
	if len(joueur.Equipe) == 0 {
		fmt.Println(Jaune("\nVous n'avez pas de Pokémon pour combattre. Créez d'abord votre dresseur."))
		return
	}

	pokemonJoueur := &joueur.Equipe[0]
	ennemi := GenerateWildPokemon()

	fmt.Printf(Jaune("\nUn %s sauvage de niveau %d apparaît!\n"), ennemi.Nom, ennemi.Niveau)

	for pokemonJoueur.EstVivant() && ennemi.EstVivant() {
		fmt.Println(Jaune("\nQue voulez-vous faire ?"))
		fmt.Println(Jaune("1. Attaquer"))
		fmt.Println(Jaune("2. Utiliser une Potion"))
		fmt.Println(Jaune("3. Lancer une Pokéball"))
		fmt.Println(Jaune("4. Fuir"))

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
			fmt.Println(Jaune("Vous avez fui le combat!"))
			return
		default:
			fmt.Println(Jaune("Choix invalide."))
		}
	}

<<<<<<< HEAD
	if EstVivant(pokemonJoueur) {
		fmt.Println(Jaune("\nChoisissez votre nouveau Pokémon :"))
		fmt.Println(Jaune("1. Bulbizarre (Type: Plante)"))
		fmt.Println(Jaune("2. Salamèche (Type: Feu)"))
		fmt.Println(Jaune("3. Carapuce (Type: Eau)"))
		fmt.Println(Jaune("4. Vibrannif (Type: Eau)"))
		fmt.Println(Jaune("5. Roucool (Type: Vol)"))
		fmt.Println(Jaune("6. Rattata (Type: Normal)"))
		fmt.Println(Jaune("7. Aspicot (Type: Insecte)"))

		var choixPokemon string
		fmt.Print(Vert("Entrez votre choix (1-7) : "))
		fmt.Scanln(&choixPokemon)

		// Vérifier si le choix est valide
		for choixPokemon < "1" || choixPokemon > "7" {
			fmt.Println(Vert("Choix invalide. Veuillez choisir entre 1 et 7."))
			fmt.Print(Vert("Entrez votre choix (1-7) : "))
			fmt.Scanln(&choixPokemon)
		}

		choixPokemonFunc(choixPokemon, joueur)

		fmt.Println(Jaune("Vous avez gagné le combat!"))
=======
	if pokemonJoueur.EstVivant() {
		expGained := ennemi.Niveau * 10
		fmt.Printf(Jaune("\nVous avez gagné le combat! %s gagne %d points d'expérience.\n"), pokemonJoueur.Nom, expGained)
		if pokemonJoueur.GainExperience(expGained) {
			fmt.Printf(Jaune("%s passe au niveau %d!\n"), pokemonJoueur.Nom, pokemonJoueur.Niveau)
		}
>>>>>>> 54389c1d0dee910e79e8c4a7cbff2eece45d7b51
	} else {
		fmt.Println(Jaune("Vous avez perdu le combat..."))
	}
}

func GenerateWildPokemon() Pokemon {
	wildPokemons := []struct {
		name string
		pokemonType PokemonType
	}{
		{"Rattata", Normal},
		{"Pidgey", Flying},
		{"Caterpie", Bug},
		{"Weedle", Bug},
		{"Pikachu", Electric},
	}

	randomPokemon := wildPokemons[rand.Intn(len(wildPokemons))]
	level := rand.Intn(5) + 1
	return Pokemon{
		Nom:       randomPokemon.name,
		PVActuels: level * 10,
		PVMax:     level * 10,
		Niveau:    level,
		Type:      randomPokemon.pokemonType,
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
				if rand.Float64() < catchChance {
					joueur.Equipe = append(joueur.Equipe, *pokemon)
					fmt.Printf(Jaune("\nFélicitations! Vous avez capturé %s!\n"), pokemon.Nom)
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
