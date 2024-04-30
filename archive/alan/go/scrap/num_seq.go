package main

import (
	"fmt"
	"math/rand"
)

func get_idx(arr []byte) int {
	idx := 0
	idx += int(arr[3])
	idx += int(arr[2]) * 4
	idx += int(arr[1]) * 16
	idx += int(arr[0]) * 64 
	return idx
}

func main() {
	size := 10000000
	seq := make([]byte, size)
	
	for i := range seq {
		seq[i] = byte(rand.Intn(4))
	}
	
	// kmer with k = 4
	kmers := make([]float64, 256)
	for n := range kmers {
		kmers[n] = rand.Float64()
	}
	
	// look up
	count := 0.0
	for i := 0; i < len(seq) -3; i++ {
		s := seq[i:i+4]
		count += kmers[get_idx(s)]
	}
	

	fmt.Println("done")
}

