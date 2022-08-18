package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func main() {
	size := 10000000
	s := make([]string, size)
	
	nt := "ACGT"
	for i := range s {
		s[i] = string(nt[rand.Intn(4)])
	}
	
	seq := strings.Join(s,"")
	
	// kmer with k = 4
	kmers := make(map[string]float64)
	for _, nt1 := range nt {
		for _, nt2 := range nt {
			for _, nt3 := range nt {
				for _, nt4 := range nt {
					kmer := make([]string, 4)
					kmer[0] = string(nt1)
					kmer[1] = string(nt2)
					kmer[2] = string(nt3)
					kmer[3] = string(nt4)
					key := strings.Join(kmer, "")
					kmers[key] = rand.Float64()
				}
			}
		}
	}
	
	// look up
	count := 0.0
	for i := 0; i < len(seq) -3; i++ {
		s := seq[i:i+4]
		count += kmers[s]
	}

	fmt.Println("done")
}
