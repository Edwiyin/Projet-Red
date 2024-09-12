package gokemon

import (
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Nom       string
	Type      string
	Niveau    int
	PVMax     int
	PVActuels int
	Attaque   int
}



type Item struct {
	Nom      string
	Quantite int
}

type Dresseur struct {
	Nom        string
	Equipe     []Pokemon
	Inventaire []Item
}

func (p *Pokemon) EstVivant() bool {
	return p.PVActuels > 0
}

func (p *Pokemon) Attaquer(cible *Pokemon) {
	degats := rand.Intn(p.Attaque) + 1
	cible.PVActuels -= degats
	if cible.PVActuels < 0 {
		cible.PVActuels = 0
	}
	fmt.Printf("%s attaque %s et lui inflige %d dégâts!\n", p.Nom, cible.Nom, degats)
	fmt.Printf("%s a maintenant %d PV\n", cible.Nom, cible.PVActuels)
}