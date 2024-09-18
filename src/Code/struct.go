package gokemon

import (
	"math/rand"
)

const CapaciteInventaireInitiale = 10
const LimiteInvInitiale = 10
const AugmentationInv = 10
const MaxAugmentationsInv = 5

const (
	Normal   PokemonType = "Normal"
	Fire     PokemonType = "Fire"
	Water    PokemonType = "Water"
	Grass    PokemonType = "Grass"
	Electric PokemonType = "Electric"
	Flying   PokemonType = "Flying"
	Bug      PokemonType = "Bug"
)

var typeEffectiveness = map[PokemonType]map[PokemonType]float64{
	Normal:   {Normal: 1, Fire: 1, Water: 1, Grass: 1, Electric: 1, Flying: 1, Bug: 1},
	Fire:     {Normal: 1, Fire: 0.5, Water: 0.5, Grass: 2, Electric: 1, Flying: 1, Bug: 2},
	Water:    {Normal: 1, Fire: 2, Water: 0.5, Grass: 0.5, Electric: 1, Flying: 1, Bug: 1},
	Grass:    {Normal: 1, Fire: 0.5, Water: 2, Grass: 0.5, Electric: 1, Flying: 0.5, Bug: 0.5},
	Electric: {Normal: 1, Fire: 1, Water: 2, Grass: 0.5, Electric: 0.5, Flying: 2, Bug: 1},
	Flying:   {Normal: 1, Fire: 1, Water: 1, Grass: 2, Electric: 0.5, Flying: 1, Bug: 2},
	Bug:      {Normal: 1, Fire: 0.5, Water: 1, Grass: 2, Electric: 1, Flying: 0.5, Bug: 1},
}

var SkillName = map[PokemonType][]string{
	Normal:   {"Charge", "Coup d'Boule", "Vive-Attaque", "Morsure"},
	Fire:     {"Flammèche", "Déflagration", "Lance-Flamme", "Surchauffe"},
	Water:    {"Pistolet à O", "Hydrocanon", "Cascade", "Surf"},
	Grass:    {"Tranch'Herbe", "Fouet Lianes", "Tempête Verte", "Lance-Soleil"},
	Electric: {"Éclair", "Cage-Éclair", "Tonnerre", "Fatal-Foudre"},
	Flying:   {"Picpic", "Cru-Aile", "Aéropique", "Hurle-Temps"},
	Bug:      {"Piqûre", "Toile", "Morsure", "Giga-Sangsue"},
}

var SkillDamage = map[PokemonType]map[PokemonType]int{
	Normal: {
		Normal: 10,
		Fire:   5,
		Water:  5,
		Grass:  5,
		Electric: 5,
	},
	Fire: {
		Normal: 10,
		Fire:   5,
		Water:  2,
		Grass:  15,
		Electric: 10,
	},
	Water: {
		Normal: 10,
		Fire:   15,
		Water:  5,
		Grass:  2,
		Electric: 5,
	},
	Grass: {
		Normal: 10,
		Fire:   2,
		Water:  15,
		Grass:  5,
		Electric: 10,
	},
	Electric: {
		Normal: 10,
		Fire:   10,
		Water:  15,
		Grass:  5,
		Electric: 5,
	},
}

type Pokemon struct {
	Nom                   string
	PVActuels             int
	PVMax                 int
	Niveau                int
	Type                  PokemonType
	Experience            int
	Attaque               int
	ExperienceToNextLevel int
	Skills                []string
}

type InventoryItem struct {
	Nom      string
	Quantite int
}

type Dresseur struct {
	Nom        string
	Equipe     []Pokemon
	Inventaire []InventoryItem
	Argent     int
	Ressources []Resource
	Equipement struct {
		Tete  Equipment
		Torse Equipment
		Pieds Equipment
	}
	CapaciteInventaire int
	NombreAugmentationsInv int
}

type Equipment struct {
	Nom        string
	Emplacement string
	BonusPV    int
	BonusAttack int
}

func (p *Pokemon) IsAlive() bool {
	return p.PVActuels > 0
}

func (p *Pokemon) Attack(cible *Pokemon) int {
	baseDamage := rand.Intn(p.Attaque) + 1
	effectiveness := typeEffectiveness[p.Type][cible.Type]
	damage := int(float64(baseDamage) * effectiveness)
	cible.PVActuels -= damage
	if cible.PVActuels < 0 {
		cible.PVActuels = 0
	}
	return damage
}

func (p *Pokemon) LevelUp() {
	p.Niveau++
	p.PVMax += rand.Intn(3) + 2
	p.Attaque += rand.Intn(2) + 1
	p.PVActuels = p.PVMax
	p.ExperienceToNextLevel = int(float64(p.ExperienceToNextLevel) * 1.2)
	p.UnlockSkills()
}

func (p *Pokemon) UnlockSkills() {
	if p.Niveau >= 2 {
		p.Skills = SkillName[p.Type]
	}
}

func NewPokemon(nom string, pokemonType PokemonType, niveau int) Pokemon {
	baseHP := 30 + rand.Intn(10)
	baseAttack := 5 + rand.Intn(5)
	return Pokemon{
		Nom:                   nom,
		Type:                  pokemonType,
		Niveau:                niveau,
		PVMax:                 baseHP + (niveau-1)*3,
		PVActuels:             baseHP + (niveau-1)*3,
		Attaque:               baseAttack + (niveau-1)*2,
		Experience:            0,
		ExperienceToNextLevel: 100,
	}
}

func (d *Dresseur) EquiperEquipement(equipement Equipment) {
	var ancienEquipement Equipment
	switch equipement.Emplacement {
	case "Tête":
		ancienEquipement = d.Equipement.Tete
		d.Equipement.Tete = equipement
	case "Torse":
		ancienEquipement = d.Equipement.Torse
		d.Equipement.Torse = equipement
	case "Pieds":
		ancienEquipement = d.Equipement.Pieds
		d.Equipement.Pieds = equipement
	}
	for i, item := range d.Inventaire {
		if item.Nom == equipement.Nom && item.Quantite > 0 {
			d.Inventaire[i].Quantite--
			if d.Inventaire[i].Quantite == 0 {
				d.Inventaire = append(d.Inventaire[:i], d.Inventaire[i+1:]...)
			}
			break
		}
	}

	if ancienEquipement.Nom != "" {
		d.Inventaire = append(d.Inventaire, InventoryItem{Nom: ancienEquipement.Nom, Quantite: 1})
	}

	for i := range d.Equipe {
		d.Equipe[i].PVMax += equipement.BonusPV
		d.Equipe[i].Attaque += equipement.BonusAttack
		if d.Equipe[i].PVActuels > d.Equipe[i].PVMax {
			d.Equipe[i].PVActuels = d.Equipe[i].PVMax
		}
	}
}

type Resource struct {
    Nom      string
    Quantite int
}

var TypeToResource = map[PokemonType]string{
    Normal:   "Fourrure",
    Fire:     "Charbon",
    Water:    "Écaille",
    Grass:    "Feuille",
    Electric: "Batterie",
    Flying:   "Plume",
    Bug:      "Carapace",
}

type CraftingRecipe struct {
    Nom         string
    Ressources  map[string]int
    CoutArgent  int
}

var CraftingRecipes = map[string]CraftingRecipe{
    "Casque": {
        Nom:        "Casque",
        Ressources: map[string]int{"Fourrure": 2, "Écaille": 1},
        CoutArgent: 50,
    },
    "Armure": {
        Nom:        "Armure",
        Ressources: map[string]int{"Écaille": 3, "Charbon": 1},
        CoutArgent: 50,
    },
    "Bottes": {
        Nom:        "Bottes",
        Ressources: map[string]int{"Plume": 2, "Carapace": 1},
        CoutArgent: 50,
    },
}

func (d *Dresseur) AddResource(resource string, quantite int) {
    for i, item := range d.Inventaire {
        if item.Nom == resource {
            d.Inventaire[i].Quantite += quantite
            return
        }
    }
   
    d.Inventaire = append(d.Inventaire, InventoryItem{Nom: resource, Quantite: quantite})
}
