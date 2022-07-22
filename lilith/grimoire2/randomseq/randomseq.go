// This package generates a number of random sequences in FASTA format and prints
// it to the console. Users should ensure that their nucleotide frequencies sum
// one as much as possible.


package randomseq

import (
  "error"
  "fmt"
  "math/rand"
  "sort"
)


//Choice pairs a given string to a weight.
type Choice struct {
  Option str
  Weight float32
}

// This function will create a new Choice based on the dictionary and the weights
// input by the user.
func NewChoice(option str, weight float32) Choice {
  return Choice{Option: option, Weight: weight}
}
