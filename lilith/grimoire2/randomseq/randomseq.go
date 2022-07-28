// This package generates a number of random sequences in FASTA format and prints
// it to the console. Users should ensure that their nucleotide frequencies sum
// one as much as possible.

package randomseq

import (
	//"error"
	"fmt"
	"strings"
	//"math/rand"
	//"sort"
)

// Choice pairs a given string to a weight.
type Choice struct {
	Option string
	Weight float32
}

// The Chooser object contains all the data needed to generate the sequence,
// including the options, their weights, and the running total
type Chooser struct {
	Choices []Choices
	Breakpoints []float32
	Maxweight float32
}

// This function will create a new Choice based on the dictionary and the weights
// input by the user.
func NewChoice(option string, weight float32) Choice {
	return Choice{Option: option, Weight: weight}
}

// Creates a Chooser object, which sorts the choices in ascending order of their
// weights to prepare them to generate the sequence.
func NewChooser(choices []Choice) (*Chooser) {
	//  sort.Slice(choices, func(i, j) int) bool {
	//    return choices[i].Weight < choices[j].Weight
	//  }
	//breakpoints := make([]float32, len(choices))
	var total float32

	for _, nt := range choices {
		weight := nt.Weight
		total += nt.Weight
		fmt.Println(weight, total)
	}
}

func Run() {
	dictionary := "A,C,T,G"
	dictslice := strings.Split(dictionary, ",")
	freqs := [4]float32 {0.24, 0.36, 0.30, 0.10}
	demos := make([]Choice, len(dictslice))

	for i := range dictslice {
		demos[i] = Choice{dictslice[i], freqs[i]}
	}

	NewChooser(demos)
}
