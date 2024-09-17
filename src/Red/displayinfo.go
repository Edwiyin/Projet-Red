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

func ViewTeam(joueur *Dresseur) {
	if len(joueur.Equipe) == 0 {
		fmt.Println(Jaune("\nVous n'avez pas encore de Pokémon dans votre équipe."))
		return
	}

	fmt.Println(Jaune("\nVotre équipe Pokémon :"))
	for i, pokemon := range joueur.Equipe {
		fmt.Printf(Jaune("%d. %s (Type: %s, Niveau: %d, PV: %d/%d, Attaque: %d, Exp: %d/%d)\n"),
			i+1, pokemon.Nom, pokemon.Type, pokemon.Niveau, pokemon.PVActuels, pokemon.PVMax,
			pokemon.Attaque, pokemon.Experience, pokemon.ExperienceToNextLevel)
	}
}