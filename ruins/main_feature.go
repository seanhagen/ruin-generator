package ruins

/**
 * File: main_feature.go
 * Date: 2021-11-03 15:27:40
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

type mainFeature struct {
	min  int
	max  int
	name string
	doer func(*Exit)
}

var mainFeatureTable = rollingTable{
	diceRoll: "1d100",
	results: []feature{
		{-1, -1, "Exit/Entrance"},
		{1, 20, "Corridor"},
		{21, 58, "Chamber"},
		{59, 70, "Creature"},
		{71, 75, "Explorers"},
		{76, 79, "Interstitial cavity"},
		{80, 82, "Accessway"},
		{83, 85, "Rupture"},
		{86, 88, "Shaft"},
		{89, 91, "Abhuman colony"},
		{92, 93, "Integrated machine"},
		{94, 95, "Matter leak"},
		{96, 97, "Energy discharge"},
		{98, 98, "Weird event"},
		{99, 99, "Vault"},
		{100, 100, "Relic chamber"},
	},
}

func rollMainFeature(level int) (string, int, error) {
	feat, err := coreList.getResult()
	if err != nil {
		return "", -1, err
	}
	return feat.Feature, level + 1, nil
}
