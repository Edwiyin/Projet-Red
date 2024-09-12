package main

import (
	gokemon "gokemon/Red"
	"math/rand"
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

func main() {
	rand.Seed(time.Now().UnixNano())

	gokemon.MenuPrincipal(&joueur)
}
