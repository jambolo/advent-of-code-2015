package main

import (
	"fmt"
	"log"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseIngredients(lines []string) []ingredient {
	ingredients := make([]ingredient, len(lines))

	for i, line := range lines {
		var name string
		var capacity, durability, flavor, texture, calories int

		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name, &capacity, &durability, &flavor, &texture, &calories)

		ingredients[i] = ingredient{
			name:       name,
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		}
	}

	return ingredients
}

func main() {
	day := 15

	path, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	ingredients := parseIngredients(lines)

	ingredientCount := len(ingredients)

	compositions := utils.Compositions(100, ingredientCount)
	bestScore := 0
	for _, c := range compositions {
		totalCapacity := 0
		totalDurability := 0
		totalFlavor := 0
		totalTexture := 0
		totalCalories := 0

		for i, amount := range c {
			totalCapacity += amount * ingredients[i].capacity
			totalDurability += amount * ingredients[i].durability
			totalFlavor += amount * ingredients[i].flavor
			totalTexture += amount * ingredients[i].texture
			totalCalories += amount * ingredients[i].calories
		}

		totalCapacity = max(totalCapacity, 0)
		totalDurability = max(totalDurability, 0)
		totalFlavor = max(totalFlavor, 0)
		totalTexture = max(totalTexture, 0)
		totalCalories = max(totalCalories, 0)

		score := totalCapacity * totalDurability * totalFlavor * totalTexture

		if (part == 1) || (part == 2 && totalCalories == 500) {
			bestScore = max(bestScore, score)
		}
	}

	fmt.Printf("Best score: %d\n", bestScore)
}
