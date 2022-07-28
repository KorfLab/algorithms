// This package generates a number of random sequences in FASTA format and prints
// it to the console. Users should ensure that their nucleotide frequencies sum
// one as much as possible.


package randomseq

import (
  //"error"
  "fmt"
  //"math/rand"
  //"sort"
)


//Choice pairs a given string to a weight.
type Choice struct {
  Option string
  Weight float32
}

type Chooser struct {
  Choices []Choice
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
  breakpoints := make([]float32, len(choices))
  var total float32 = 0.00


  for i, nt := range choices {
    breakpoints[i] = nt.Weight
    total += choices[i].Weight
  }

  return &Chooser{Choices: choices, Breakpoints: breakpoints, Maxweight: total}
}

func MakeChoices () (*Chooser, error) {
  demos := make([]Choice, 4)
  demos[0] = Choice{"A", 0.10}
  demos[1] = Choice{"C", 0.20}
  demos[2] = Choice{"G", 0.30}
  demos[3] = Choice{"T", 0.40}

  weighted := NewChooser(demos)
}
