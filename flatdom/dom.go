package flatdom

import (
	"fmt"
	"math/rand"
)

// Dom is to represent a flatdom containing the land and updater.
type Dom struct {
	land    [landLength][landWidth]*Grid
	updater *Updater
}

// Init is to initialize the land.
func (d *Dom) Init() {
	rand.Seed(42)

	for i := range d.land {
		for j := range d.land[i] {
			d.land[i][j] = &Grid{}
			seed := rand.Intn(100)
			if seed < initFactionNum {
				d.land[i][j].holder = rand.Intn(100)%2 + 1
				d.land[i][j].influence = float64(rand.Intn(10)) + 2
				fmt.Printf("holder: %d, influence: %f\n", d.land[i][j].holder, d.land[i][j].influence)
			} else {
				d.land[i][j].holder = 0
				d.land[i][j].influence = 0
			}
		}
	}

	d.printout(-1)

	d.updater = &Updater{}
}

// Run is to execute one phase, compute events and apply the afterwards results.
func (d *Dom) Run(num int) {
	totalNum := num
	for num > 0 {
		for i := range d.land {
			for j := range d.land[i] {
				if d.land[i][j].holder == 0 {
					continue
				}
				surr := []*Grid{}
				if i > 0 {
					surr = append(surr, d.land[i-1][j])
				}
				if i < landLength-1 {
					surr = append(surr, d.land[i+1][j])
				}
				if j > 0 {
					surr = append(surr, d.land[i][j-1])
				}
				if j < landWidth-1 {
					surr = append(surr, d.land[i][j+1])
				}
				d.updater.update(d.land[i][j], surr)
			}
		}

		d.printout(totalNum - num + 1)

		num--
	}

}

func (d *Dom) printout(k int) {
	fmt.Println("Holder:")
	for i := range d.land {
		for j := range d.land[i] {
			fmt.Printf("%d ", d.land[i][j].holder)
		}
		fmt.Println()
	}
	fmt.Println("Influence:")
	for i := range d.land {
		for j := range d.land[i] {
			fmt.Printf("%.1f ", d.land[i][j].influence)
		}
		fmt.Println()
	}
	fmt.Printf("Iteration %d finished\n", k)
}
