package gokemon

import (
	"math/rand"
	"fmt"
)

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

func GenerateWildPokemon(joueur *Dresseur) Pokemon {
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
        {"Vaporeon", Water},
        {"Jolteon", Electric},
        {"Flareon", Fire},
        {"Omanyte", Water},
        {"Kabuto", Water},
        {"Aerodactyl", Flying},
        {"Snorlax", Normal},
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
        {"Magikarp", Water},        
        {"Scyther", Bug},
        {"Electabuzz", Electric},
        {"Magmar", Fire},
        {"Ronflex", Normal},
        {"Dracaufeu", Fire},
        {"Tortank", Water},
        {"Florizarre", Grass},
        
    }

    maxLevel := 1
    for _, pokemon := range joueur.Equipe {
        if pokemon.Niveau > maxLevel {
            maxLevel = pokemon.Niveau
        }
    }

    var level int
    if maxLevel <= 5 {
        level = rand.Intn(3) + 1
    } else {
        minLevel := maxLevel - 2
        maxLevel := maxLevel + 2
        level = rand.Intn(maxLevel-minLevel+1) + minLevel
    }

    randomPokemon := wildPokemons[rand.Intn(len(wildPokemons))]
    
    return Pokemon{
        Nom:        randomPokemon.name,
        PVMax:      level * 10,
        PVActuels:  level * 10,
        Niveau:     level,
        Type:       randomPokemon.pokemonType,
        Experience: 0,
        Attaque:    level * 5,
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