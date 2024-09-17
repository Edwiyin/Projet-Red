package gokemon

import "fmt"

func DisplayInfo(joueur Dresseur) {
	fmt.Println(Jaune("\nInformations du dresseur :"))
	fmt.Printf(Jaune("Nom : %s\n"), joueur.Nom)
	fmt.Printf(Jaune("Solde : %d PokéDollars\n"), joueur.Argent)
	fmt.Println(Jaune("\nÉquipe Pokémon :"))
	for _, pokemon := range joueur.Equipe {
		fmt.Printf(Jaune("- %s (Type: %s, Niveau: %d, Exp: %d/%d, PV: %d/%d, Attaque: %d)\n"), pokemon.Nom, pokemon.Type, pokemon.Niveau, pokemon.Experience, pokemon.Niveau*100, pokemon.PVActuels, pokemon.PVMax, pokemon.Attaque)
	}

}
