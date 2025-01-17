/*
	debug example
*/

package main

import (
	"fmt"
	"github.com/walterdejong/debug"
)

// calcPi calculates Pi using Nilakantha series.
// We only calculate up to 5 digits correct.
func calcPi() float64 {
	var pi float64 = 3.0

	var d int = 2
	plus := true

	debug.Debug("iteration #%d  pi == %.06f", 0, pi)

	for i := 1; i < 9999; i++ {
		div := float64(d * (d + 1) * (d + 2))
		d += 2

		term := 4.0 / div
		if plus {
			pi += term
		} else {
			pi -= term
		}
		plus = !plus // alternate

		debug.Debug("iteration #%d  pi == %#v", i, pi)

		// stop after 5 digits
		if pi >= 3.141590 && pi < 3.141600 {
			break
		}
	}
	debug.Debug("returning pi == %#v", pi)
	return pi
}

func main() {
	debug.Enabled = true

	/*
		we can customize the colors if we want

		debug.ColorInfo = debug.DarkCyan
		debug.ColorMessage = debug.Cyan
	*/
	/*
		change output to stdout (default is stderr)

		debug.SetOutput(os.Stdout)
	*/

	debug.Debug("hello, debug!")

	fmt.Println("This program calculates (very roughly) the value of Pi")
	pi := calcPi()
	fmt.Printf("calculated Pi : %v\n", pi)

	debug.Debug("exiting")
}

// EOB
