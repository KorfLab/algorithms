package main

import (
  "flag"
  "fmt"
  "github.com/Korflab/grimoire2/read_record"
  "math"
)

type Vnode struct {
  prob float64
  traceback *Vnode
}

//creates a simple model, only accounting for one position.
func makemodel () map[string]float64 {
  var bases = make(map[string]float64)

  for _, letter := range("ACTG") {
    bases[letter] = 0.0
  }
}

//Turns probabilities into log-odds scores
func logodds (p float64) float64 {
  if p == 0 {
    return math.SmallestNonzeroFloat64
  } else {
    return math.Log2(p/0.25)
  }
}


func trainmodel (file) map[string]float64 {
  model = makemodel()

  fasta = read_record.Read_fasta(file)

  for fasta.Next() {
    record := fasta.Current()

    for _, letter := range(record.Seq) {
      model[letter] += 1.0
    }
  }
  total := 0.0

  for _, count := range model {
    total += count
  }

  for letter, count := range model {
    prob := count / total
    model[letter] = logodds(prob)
  }

  return model
}


func main () {
  flag.StringVar(&infile, "f", "", "query file")
  flag.StringVar(&posfile, "pos", "positive strand training file")
  flag.StringVar(&negfile, "neg", "negative strand training file")

  flag.Parse()

  positivemodel = trainmodel(posfile)
  negativemodel = trainmodel(negfile)

  transition
}
