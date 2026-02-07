package main

import (
	"fmt"

	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
)

// isFlying returns true if the reindeer is flying at the given time, false if it is resting.
func isFlying(time int, flyTime int, cycleTime int) bool {
	flyingTime := (time + cycleTime - 1) % cycleTime
	if flyingTime < flyTime {
		return true
	}
	return false
}

func main() {
	day := 14

	// Grab the command line parameters (file path and part number)
	_, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	cometSpeed := 14
	cometFlyTime := 10
	cometRestTime := 127

	dancerSpeed := 16
	dancerFlyTime := 11
	dancerRestTime := 162

	cometCycleTime := cometFlyTime + cometRestTime
	cometCycleDistance := cometSpeed * cometFlyTime

	dancerCycleTime := dancerFlyTime + dancerRestTime
	dancerCycleDistance := dancerSpeed * dancerFlyTime

	time := 2503

	if part == 1 {
		cometTotalCycles := time / cometCycleTime
		cometRemainingTime := time % cometCycleTime
		cometRemainingDistance := min(cometFlyTime, cometRemainingTime) * cometSpeed
		cometTotalDistance := cometTotalCycles*cometCycleDistance + cometRemainingDistance

		dancerTotalCycles := time / dancerCycleTime
		dancerRemainingTime := time % dancerCycleTime
		dancerRemainingDistance := min(dancerFlyTime, dancerRemainingTime) * dancerSpeed
		dancerTotalDistance := dancerTotalCycles*dancerCycleDistance + dancerRemainingDistance

		fmt.Printf("Winner distance: %d\n", max(cometTotalDistance, dancerTotalDistance))
	}

	if part == 2 {
		cometPoints := 0
		dancerPoints := 0
		cometDistance := 0
		dancerDistance := 0

		for t := 1; t <= time; t++ {
			if isFlying(t, cometFlyTime, cometCycleTime) {
				cometDistance += cometSpeed
			}

			if isFlying(t, dancerFlyTime, dancerCycleTime) {
				dancerDistance += dancerSpeed
			}

			if cometDistance >= dancerDistance {
				cometPoints++
			}
			if dancerDistance >= cometDistance {
				dancerPoints++
			}

			cf := isFlying(t, cometFlyTime, cometCycleTime)
			df := isFlying(t, dancerFlyTime, dancerCycleTime)
			fmt.Printf("%d: comet: %v %d %d, dancer: %v %d %d\n", t, cf, cometDistance, cometPoints, df, dancerDistance, dancerPoints)
		}
		fmt.Printf("Comet points: %d\n", cometPoints)
		fmt.Printf("Dancer points: %d\n", dancerPoints)
		fmt.Printf("Winner points: %d\n", max(cometPoints, dancerPoints))
	}
}
