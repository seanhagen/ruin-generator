package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/seanhagen/ruin-generator/ruins"
)

/**
 * File: main.go
 * Date: 2021-11-03 15:16:33
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func main() {
	seedFlag := flag.Int64("seed", -1, "Seed value to use, values < 0 means use current time")
	flag.Parse()

	seed := time.Now().UnixNano()
	if *seedFlag > 0 {
		seed = *seedFlag
	}
	rand.Seed(seed)

	ent, corr, err := ruins.GetEntrance()
	if err != nil {
		fmt.Printf("not able to get entrance: %v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("ruin before sorting out exits: \n\n")
	// fmt.Printf("%v\n\n-----------------------------------\n", ent)

	// fmt.Printf("\n\n\nEnd of entrance corridor: \n%v\n\n\n", corr)

	err = ruins.FillExits(corr)
	if err != nil {
		fmt.Printf("Unable to fill ruin: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n\nafter sorting out exits: \n\n%v", ent)

	// fmt.Printf("\n\n\nEnd of entrance corridor after populating exits: \n%v\n\n\n", corr)
}
