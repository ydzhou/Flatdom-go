package flatdom

import (
	"fmt"
	"math/rand"
	"time"
)

type Updater struct {
}

func (u *Updater) update(target *Grid, surr []*Grid) {
	if target.holder > 0 && u.canDo(devProb) && target.influence < maxGridInf {
		u.dev(target)
		fmt.Printf("DEV: holder %d; inf %.1f\n", target.holder, target.influence)
		return
	}
	for i := range surr {
		if surr[i].holder == target.holder {
			continue
		}
		if surr[i].holder > 0 && surr[i].influence > target.influence && u.canDo(warProb) {
			fmt.Printf("PRE-WAR: holder %d; inf %.1f\n", target.holder, target.influence)
			u.warUpdate(target, surr[i])
			fmt.Printf("POST-WAR: holder %d; inf %.1f\n", target.holder, target.influence)
			break
		}
		if surr[i].holder == 0 && target.influence > 1 && u.canDo(expandProb) {
			u.exploreUpdate(target, surr[i])
			fmt.Printf("EXP: holder %d; inf %.1f\n", target.holder, target.influence)
			break
		}
	}
}

func (u *Updater) canDo(prob int) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) <= prob
}

func (u *Updater) warUpdate(src *Grid, dst *Grid) {
	src.influence = src.influence - dst.influence
	dst.holder = 0
	dst.influence = 0
}

func (u *Updater) exploreUpdate(src *Grid, dst *Grid) {
	dst.holder = src.holder
	dst.influence = src.influence / 2
	src.influence = src.influence - src.influence/2
}

func (u *Updater) dev(g *Grid) {
	rand.Seed(time.Now().UnixNano())
	change := rand.Float64() * (maxGridInf - g.influence)
	fmt.Printf("DEV INF CHANGE: %.1f\n", change)
	g.influence += change
}
