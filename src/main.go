package main

import (
	"fmt"
	gokemon "gokemon/Red"
	"math/rand"
	"time"
)

var joueur gokemon.Dresseur
var audioManager *gokemon.AudioManager

func init() {
	joueur = gokemon.Dresseur{
		Nom:    "",
		Equipe: []gokemon.Pokemon{},
		Inventaire: []gokemon.InventoryItem{
			{Nom: "Potion", Quantite: 3},
			{Nom: "Pokéball", Quantite: 5},
		},
	}

	audioManager = gokemon.NewAudioManager()

	err := audioManager.LoadBackgroundMusic("assets/music/gokemon.mp3")
	if err != nil {
		fmt.Println("Erreur lors du chargement de la musique de fond:", err)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	audioManager.PlayBackgroundMusic()

	gokemon.MenuPrincipal(&joueur, audioManager)

}

fmt.Println("Fin du jeu")