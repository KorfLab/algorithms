// This package generates a number of random sequences in FASTA format and prints
// it to the console. Users should ensure that their nucleotide frequencies sum
// one as much as possible.

package randomseq

import (
	//"error"
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Choice pairs a given string to a weight.
type Choice struct {
	Option string
	Weight float64
}

// The Chooser object contains all the data needed to generate the sequence,
// including the options, their weights, and the running total
type Chooser struct {
	Ntoptions []string
	Breakpts  []float64
	Maxweight float64
}

// This function will create a new Choice based on the dictionary and the weights
// input by the user.
func NewChoice(option string, weight float64) Choice {
	return Choice{Option: option, Weight: weight}
}

// Creates a Chooser object, which pairs the choices off with summative weights,
// which serve as breakpoints to shift from one option to the next. Sorting is
// unnecessary as the order of the options doesn't matter.
func NewChooser(choices []Choice) *Chooser {
	breakpts := make([]float64, len(choices))
	ntoptions := make([]string, len(choices))
	var breakpt float64 = 0.0

	for i, nt := range choices {
		breakpt += nt.Weight
		breakpts[i] = breakpt
		ntoptions[i] = nt.Option
	}

	maxweight := breakpt

	return &Chooser{Ntoptions: ntoptions, Breakpts: breakpts, Maxweight: maxweight}
}

// Takes in the string of letters used to make our sequence and pairs them
// in a Choice object in an array, then returns that array.
func MakeChooser(dictionary string, freqs []float64) *Chooser {
	dictslice := strings.Split(dictionary, ",")
	demos := make([]Choice, len(dictslice))

	for i := range dictslice {
		demos[i] = Choice{dictslice[i], freqs[i]}
	}

	ntrand := NewChooser(demos)

	return ntrand
}

// Seqprint is a funciton that prints a user-specified number of sequences of an
// arbitrary length. It uses the chooser object made previously, printing a FASTA
// format sequence.
func Seqprint(ch *Chooser, count int, size int) {
	var i, j, k, printlen int
	const printwidth int = 80
	divider := float64(size/printwidth)
	extra := int(math.Floor(divider))
	extra += size
	printstring := make([]string, extra)

	for i < count {
		fmt.Println(">id ", i+1)

		for j < extra {
			r := rand.Float64()
			k = conditional(ch.Breakpts, r)
			printstring[j] = ch.Ntoptions[k]
			printlen++

			if printlen == printwidth {
				j++
				printstring[j] = "\n"
				printlen = 0
			}
			j++
		}

		i++
		fmt.Print(strings.Join(printstring, ""), "\n")
		j = 0
	}
}

// This is an implemented binary search tree to find the correct base to call
// for a given position, using a random number and a set of breakpoints that
// were determined in NewChooser. The goal of this is to be more efficient with
// repeated calls to the same structure.
func findfloat(breaks []float64, selector float64) int {
	index, max := 0, len(breaks)

	for index < max {
		half := int(uint(index+max) >> 1)

		if breaks[half] < selector {
			index = half + 1
		} else {
			max = half
		}
	}

	return index
}

// Scans through the breakpoints and returns the index of the base that was
// selected for by the random number generator. 
func conditional(breaks []float64, selector float64) int {

	if selector < breaks[0] {
		return 0
	} else if selector < breaks[1] {
		return 1
	} else if selector < breaks[2] {
		return 2
	} else {
		return 3
	}
}
