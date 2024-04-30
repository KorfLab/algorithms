// This package generates a number of random sequences in FASTA format and prints
// it to the console. Users should ensure that their nucleotide frequencies sum
// one as much as possible.

package randomseq

import (
	//"error"
	"fmt"
	"math/rand"
	"strings"
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
	Choices   []Choice
	Breakpts  []float32
	Maxweight float32
}

// This function will create a new Choice based on the dictionary and the weights
// input by the user.
func NewChoice(option string, weight float32) Choice {
	return Choice{Option: option, Weight: weight}
}

// Creates a Chooser object, which sorts the choices in ascending order of their
// weights to prepare them to generate the sequence.
func NewChooser(choices []Choice) *Chooser {
	breakpts := make([]float32, len(choices))
	var breakpt float32 = 0.0
	var total float32

	for i, nt := range choices {
		breakpt += nt.Weight
		breakpts[i] = breakpt
		fmt.Println(breakpt, total)
	}

	maxweight := breakpt

	return &Chooser{Choices: choices, Breakpts: breakpts, Maxweight: maxweight}
}

func MakeChooser(dictionary string, freqs []float32) *Chooser {
	dictslice := strings.Split(dictionary, ",")
	demos := make([]Choice, len(dictslice))

	for i := range dictslice {
		demos[i] = Choice{dictslice[i], freqs[i]}
	}

	ntrand := NewChooser(demos)

	return ntrand
}

func Picknprint(chooser *Chooser, runs int, size int) {
	var i, j, k int = 0, 0, 0
	var printstring string

	for i < runs {
		for j < size {
			rfloat := rand.Float32()
			k := Findfloat(chooser.Breakpts, rfloat)
		}
	}
}

func Findfloat(breaks []float32, selector float32) int {

}
