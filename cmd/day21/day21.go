package main

import (
	"fmt"
	"math"

	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}
var armor = []Item{
	{"None", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"None", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

var maxConfigurations = len(weapons) * len(armor) * len(rings) * len(rings)

func configuration(id int) (weaponId, armorId, ring1Id, ring2Id int) {
	weaponId = id % len(weapons)
	id /= len(weapons)

	armorId = id % len(armor)
	id /= len(armor)

	ring1Id = id % len(rings)
	id /= len(rings)

	ring2Id = id % len(rings)
	id /= len(rings)

	return weaponId, armorId, ring1Id, ring2Id
}

func battle(playerDamage, playerArmor int) bool {
	bossHitPoints := 104
	bossDamage := 8
	bossArmor := 1

	playerHitPoints := 100

	for playerHitPoints > 0 {
		playerDamageDealt := playerDamage - bossArmor
		if playerDamageDealt < 1 {
			playerDamageDealt = 1
		}
		bossHitPoints -= playerDamageDealt

		if bossHitPoints <= 0 {
			return true
		}

		bossDamageDealt := bossDamage - playerArmor
		if bossDamageDealt < 1 {
			bossDamageDealt = 1
		}
		playerHitPoints -= bossDamageDealt
	}

	return false
}

func main() {
	day := 21

	_, part := setup.Parameters(day)
	setup.Banner(day, part)

	if part == 1 {
		minCost := math.MaxInt
		for id := 0; id < maxConfigurations; id++ {
			weaponId, armorId, ring1Id, ring2Id := configuration(id)
			playerCost := weapons[weaponId].cost + armor[armorId].cost + rings[ring1Id].cost + rings[ring2Id].cost
			if playerCost < minCost {
				playerDamage := weapons[weaponId].damage + armor[armorId].damage + rings[ring1Id].damage + rings[ring2Id].damage
				playerArmor := weapons[weaponId].armor + armor[armorId].armor + rings[ring1Id].armor + rings[ring2Id].armor
				if battle(playerDamage, playerArmor) {
					minCost = playerCost
				}
			}
		}
		fmt.Printf("Minimum Cost: %d\n", minCost)
	}

	if part == 2 {
		maxCost := math.MinInt
		for id := 0; id < maxConfigurations; id++ {
			weaponId, armorId, ring1Id, ring2Id := configuration(id)
			playerCost := weapons[weaponId].cost + armor[armorId].cost + rings[ring1Id].cost + rings[ring2Id].cost
			if playerCost > maxCost {
				playerDamage := weapons[weaponId].damage + armor[armorId].damage + rings[ring1Id].damage + rings[ring2Id].damage
				playerArmor := weapons[weaponId].armor + armor[armorId].armor + rings[ring1Id].armor + rings[ring2Id].armor
				if !battle(playerDamage, playerArmor) {
					maxCost = playerCost
				}
			}
		}
		fmt.Printf("Minimum Cost: %d\n", maxCost)
	}
}
