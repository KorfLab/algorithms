/*
Grimoire is a collection of bioinformatics tools.

randomseq: generates a number of FASTA format DNA sequences of arbitrary length,
both values supplied by the user.

Warning! This program determines sequence using given frequencies as weights,
and as a result will function even if those weights do not add up to 1 or close
enough to it. It is the responsibility of the user to ensure that their frequencies
are correct!
*/



package main


import (
  "grimoire2/randomseq"
  "flag"
  "fmt"
  "math/rand"
  "time"
)


func main() {
  var count, size int
  var A, C, G, T float64

  flag.IntVar(&count, "count", 1000, "Desired # of sequences, default 100")
  flag.IntVar(&size, "len", 4000, "Desired size of sequences, default 1000")
  flag.Float64Var(&A, "A", 0.25, "Frequency of A")
  flag.Float64Var(&C, "C", 0.25, "Frequency of C")
  flag.Float64Var(&G, "G", 0.25, "Frequency of G")
  flag.Float64Var(&T, "T", 0.25, "Frequency of T")

  flag.Parse()

  dictionary := "A,C,G,T"
  freqs := []float64{A, C, G, T}
  rand.Seed(time.Now().UnixNano())
  fmt.Println(time.Now().UnixNano())

  ntchoices := randomseq.MakeChooser(dictionary, freqs)

  randomseq.Seqprint(ntchoices, count, size)
}
