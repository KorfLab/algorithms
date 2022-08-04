// This package generates a number of random sequences in FASTA format and prints
// it to the console. Users should ensure that their nucleotide frequencies sum
// one as much as possible.

package randomseq

import (
	//"error"
	"fmt"
	"strings"
	"math/rand"
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
	Ntoptions   []string
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
	ntchoices := make([]string, len(choices))
	var breakpt float32 = 0.0
	var total float32

	for i, nt := range choices {
		breakpt += nt.Weight
		breakpts[i] = breakpt
		fmt.Println(breakpt, total)
		ntoptions[i] = nt.Option
	}

	maxweight := breakpt


	return &Chooser{Options: ntoptions, Breakpts: breakpts, Maxweight: maxweight}
}

// Takes in the string of letters used to make our sequence and pairs them
// in a Choice object in an array, then returns that array.
func MakeChooser(dictionary string, freqs []float32) *Chooser {
	dictslice := strings.Split(dictionary, ",")
	demos := make([]Choice, len(dictslice))

	for i := range dictslice {
		demos[i] = Choice{dictslice[i], freqs[i]}
	}

	ntrand := NewChooser(demos)

	return ntrand
}


func Seqprint(chooser *Chooser, runs int, size int) {
	var i,j,k int = 0, 0, 0
	var nt,printstring string

	for i < runs {
		fmt.Println(">id ", i)
		for j < size {
			r := rand.Float32()
			k = findfloat(chooser.Breakpts, r)
			printstring += chooser.Ntoptions[k]
		}

		fmt.Println(printstring)
	}
}


func findfloat(breaks []float32, selector float32) int {
	index,max := 0, len(breaks)

	for index < max {
		half := int(uint(index+max) >> 1)

		if breaks[half] < selector {
			index = index + half
		} else {
			max = half
		}

		return index
}
