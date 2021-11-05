package ruins

import (
	"github.com/justinian/dice"
)

/**
 * File: main_feature.go
 * Date: 2021-11-03 15:27:40
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

var mainFeatureTable = rollingTable{
	diceRoll: "1d100",
	results: []result{
		chamberFeature{baseFeature{1, 20, "Corridor"}},
		corridorFeature{baseFeature{21, 58, "Chamber"}},
		creatureFeature{baseFeature{59, 70, "Creature"}},
		explorerFeature{baseFeature{71, 75, "Explorers"}},
		interstitialFeature{baseFeature{76, 79, "Interstitial cavity"}},
		accesswayFeature{baseFeature{80, 82, "Accessway"}},
		ruptureFeature{baseFeature{83, 85, "Rupture"}},
		shaftFeature{baseFeature{86, 88, "Shaft"}},
		abhumanFeature{baseFeature{89, 91, "Abhuman colony"}},
		machineFeature{baseFeature{92, 93, "Integrated machine"}},
		leakFeature{baseFeature{94, 95, "Matter leak"}},
		dischargeFeature{baseFeature{96, 97, "Energy discharge"}},
		weirdFeature{baseFeature{98, 98, "Weird event"}},
		vaultFeature{baseFeature{99, 99, "Vault"}},
		relicFeature{baseFeature{100, 100, "Relic chamber"}},
	},
}

// GetEntrance ...
func GetEntrance() (*Room, *Room, error) {
	ent, err := createChamber(0, nil)
	if err != nil {
		return nil, nil, err
	}
	ent.Type = "Entrance"

	// err = addExits(ent)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// if len(ent.Exits) <= 0 {
	ex := &Exit{nextExitID(), ExitNormal, ent, nil}
	ent.Exits = []*Exit{ex}
	// }

	corr, err := createCorridor(1, ex)
	if err != nil {
		return nil, nil, err
	}

	return ent, corr, nil
}

// FillExits ...
func FillExits(r *Room) error {
	if r.Level == MaxDepth {
		r.Exits = []*Exit{}
		return nil
	}

	for _, e := range r.Exits {
		if e.Child != nil {
			continue
		}
		res, err := mainFeatureTable.getResult()
		if err != nil {
			return err
		}

		new, err := res.Apply(e)
		if err != nil {
			return err
		}

		err = FillExits(new)
		if err != nil {
			return err
		}
	}
	return nil
}

func rollMainFeature(level int) (result, error) {
	feat, err := mainFeatureTable.getResult()
	if err != nil {
		return nil, err
	}
	return feat, nil
}

func rollRoom(level int, parent *Exit) (*Room, error) {
	ch, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}
	r := ch.Int()

	if r <= 33 {
		return createCorridor(level, parent)
	}

	if r > 33 && r <= 75 {
		return createChamber(level, parent)
	}

	if r > 75 && r <= 90 {
		return createInterstitial(level, parent)
	}

	if r > 90 && r <= 96 {
		return createShaft(level, parent)
	}

	return createRupture(level, parent)
}
