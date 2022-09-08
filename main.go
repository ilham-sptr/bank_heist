package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true

	eludedGuards := rand.Intn(100)

	if eludedGuards > 50 {
		fmt.Println("Look like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better disguise next time?")
	}

	fmt.Println("is HeistOn is currently", isHeistOn)

	openVault := rand.Intn(100)

	if openVault > 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openVault < 70 {
		fmt.Println("What's the combo to this lock again?")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		default:
			fmt.Println("Start the gateway car!")
		}
	}
}
