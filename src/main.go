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
			{Nom: "Potion", Quantite: 5},
			{Nom: "Pokéball", Quantite: 5},
		},
	}

	audioManager = gokemon.NewAudioManager()

	err := audioManager.LoadBackgroundMusic("assets/music/titlescreen1.mp3")
	if err != nil {
		fmt.Println("Erreur lors du chargement de la musique de fond:", err)
	}
	err = audioManager.LoadSoundEffect("select", "assets/music/select1.mp3")
	if err != nil {
		fmt.Println("Erreur lors du chargement de la musique de fond:", err)
	}
	err = audioManager.LoadBattleMusic("assets/music/battle.mp3")
	if err != nil {
		fmt.Println("Erreur lors du chargement de la musique de fond:", err)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	audioManager.PlayBackgroundMusic()

	gokemon.MenuPrincipal(&joueur, audioManager)

}
