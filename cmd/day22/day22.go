package main

import (
	"fmt"
	"math"

	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

type State struct {
	playerHitPoints int
	playerMana      int
	playerArmor     int
	bossHitPoints   int
	shieldTimer     int
	poisonTimer     int
	rechargeTimer   int
}

type Spell struct {
	name string
	cast func(state *State, spell Spell) int
	cost int
}

var spells = []Spell{
	{name: "Magic Missile", cast: magicMissile, cost: 53},
	{name: "Drain", cast: drain, cost: 73},
	{name: "Shield", cast: shield, cost: 113},
	{name: "Poison", cast: poison, cost: 173},
	{name: "Recharge", cast: recharge, cost: 229},
}

var bossHitPoints = 51
var bossDamage = 9

var playerHitPoints = 50
var playerMana = 500

type CacheValue struct {
	manaSpent int
	playerWon bool
}
type Cache map[State]CacheValue

func main() {
	day := 22

	_, part := setup.Parameters(day)
	setup.Banner(day, part)

	state := State{
		playerHitPoints: playerHitPoints,
		playerMana:      playerMana,
		bossHitPoints:   bossHitPoints,
	}

	cache := Cache{}
	manaSpent, playerWon := nextTurn(state, cache, 1, part)
	if playerWon {
		fmt.Printf("Mana spent: %d\n", manaSpent)
	} else {
		fmt.Println("Player lost")
	}
}

func nextTurn(state State, cache Cache, round int, part int) (int, bool) {
	// Check if this state has already been computed
	if cacheValue, found := cache[state]; found {
		return cacheValue.manaSpent, cacheValue.playerWon
	}

	startingState := state

	// For part 2, deduct 1 hit point from the player at the start of each of their turns.
	if part == 2 {
		state.playerHitPoints--
		if state.playerHitPoints <= 0 {
			cache[startingState] = CacheValue{manaSpent: 0, playerWon: false}
			return 0, false
		}
	}

	//	fmt.Printf("%s%d. Player's turn -- State: %v\n", indent(round), round, state)
	// Apply effects at the start of each turn
	applyEffects(&state)
	//	fmt.Printf("%s|   After effects -- State: %v\n", indent(round), state)

	// Check if the boss is dead after applying effects. Effects cost no mana.
	if state.bossHitPoints <= 0 {
		//		fmt.Printf("%s|   Boss is dead from effects -- State: %v\n", indent(round), state)
		cache[startingState] = CacheValue{manaSpent: 0, playerWon: true}
		return 0, true
	}

	// if the player doesn't have enough mana to cast any spell, they lose
	if state.playerMana < 53 {
		//		fmt.Printf("%s|   Player is out of mana -- State: %v\n", indent(round), state)
		cache[startingState] = CacheValue{manaSpent: 0, playerWon: false}
		return 0, false
	}

	// Find the minimum mana spent to win from this state
	minManaSpent := math.MaxInt
	playerWon := false // Assume the player hasn't won yet

	for _, spell := range spells {
		if state.playerMana >= spell.cost {
			nextState := state
			manaSpent := spell.cast(&nextState, spell)
			if manaSpent == 0 {
				continue // If the spell couldn't be cast, skip to the next spell
			}
			//			fmt.Printf("%s+--%s -- State: %v\n", indent(round), spell.name, nextState)
			if nextState.bossHitPoints > 0 {
				bossAttack(&nextState, round)
				if nextState.bossHitPoints > 0 {
					if nextState.playerHitPoints <= 0 {
						//						fmt.Printf("%s|   Player is dead -- State: %v\n", indent(round), nextState)
						continue // If the player lost, skip to the next spell
					}
					manaSpentNext, playerWonNext := nextTurn(nextState, cache, round+1, part) // Recursively continue to the next turn
					if !playerWonNext {
						continue // If the player lost in the end, skip to the next spell
					}
					manaSpent += manaSpentNext
					playerWon = playerWonNext
				} else {
					//					fmt.Printf("%s|   Boss is dead -- State: %v\n", indent(round), nextState)
					playerWon = true
				}
			} else {
				//				fmt.Printf("%s|   Boss is dead -- State: %v\n", indent(round), nextState)
				playerWon = true
			}

			minManaSpent = min(minManaSpent, manaSpent)
		}
	}
	//	if playerWon {
	//		fmt.Printf("%s|   Mana spent: %d\n", indent(round), minManaSpent)
	//	}

	// Cache the result for this state
	cache[startingState] = CacheValue{manaSpent: minManaSpent, playerWon: playerWon}
	return minManaSpent, playerWon
}

func applyEffects(state *State) {
	if state.shieldTimer > 0 {
		state.playerArmor = 7
		state.shieldTimer--
	} else {
		state.playerArmor = 0
	}

	if state.poisonTimer > 0 {
		state.bossHitPoints -= 3
		state.poisonTimer--
	}

	if state.rechargeTimer > 0 {
		state.playerMana += 101
		state.rechargeTimer--
	}
}

//func indent(round int) string {
//	return strings.Repeat("| ", round)
//}

func magicMissile(state *State, spell Spell) int {
	state.playerMana -= spell.cost
	state.bossHitPoints -= 4
	return spell.cost
}

func drain(state *State, spell Spell) int {
	state.playerMana -= spell.cost
	state.bossHitPoints -= 2
	state.playerHitPoints += 2
	return spell.cost
}

func shield(state *State, spell Spell) int {
	if state.shieldTimer > 0 {
		return 0 // Can't cast Shield if it's already active
	}
	state.playerMana -= spell.cost
	state.shieldTimer = 6
	return spell.cost
}

func poison(state *State, spell Spell) int {
	if state.poisonTimer > 0 {
		return 0 // Can't cast Poison if it's already active
	}
	state.playerMana -= spell.cost
	state.poisonTimer = 6
	return spell.cost
}

func recharge(state *State, spell Spell) int {
	if state.rechargeTimer > 0 {
		return 0 // Can't cast Recharge if it's already active
	}
	state.playerMana -= spell.cost
	state.rechargeTimer = 5
	return spell.cost
}

func bossAttack(state *State, round int) {
	//	fmt.Printf("%s|   Boss's turn\n", indent(round))
	// Apply effects at the start of each half turn
	applyEffects(state)
	//	fmt.Printf("%s|   After effects -- State: %v\n", indent(round), state)

	if state.bossHitPoints > 0 {
		damage := bossDamage - state.playerArmor
		if damage < 1 {
			damage = 1
		}
		state.playerHitPoints -= damage
	}
	//	fmt.Printf("%s|   After boss attack -- State: %v\n", indent(round), state)

}
