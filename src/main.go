package main

import (
	"fmt"
	gokemon "gokemon/Code"
	"math/rand"
	"time"
)

const LimiteInvInitiale = 10

var joueur gokemon.Dresseur
var audioManager *gokemon.AudioManager

func init() {
	joueur = gokemon.Dresseur{
		Nom:                "",
		Argent:             0,
		CapaciteInventaire: LimiteInvInitiale,
		Equipe:             []gokemon.Pokemon{},
		Inventaire:         make([]gokemon.InventoryItem, 0, LimiteInvInitiale),
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
		fmt.Println("Erreur lors du chargement de l'effet sonore de bataille:", err)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	audioManager.PlayBackgroundMusic()

	gokemon.MenuPrincipal(&joueur, audioManager)

}
