package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// make custom function that does weighted picks based on the CDF of the weights?
// reference: https://softwareengineering.stackexchange.com/questions/150616/get-weighted-random-item

// pass prob of each base later and do weighted calculation inside
func randseq(n int) string {
	var bases = []rune("ACGT")
	ret_seq := make([]rune, n)
	for i := range ret_seq {
		ret_seq[i] = bases[rand.Intn(len(bases))]
	}
	return string(ret_seq)
}

func main() {
	// set all the flags
	num_seq := flag.Int("num_seqs", 20, "enter the number of seequences you want to generate (defaults to 20)")
	len_each := flag.Int("len_each", 50, "enter the length of a sequence (defaults to 50)")
	// proba := flag.Float64("prob_A", 0.25, "enter probability of nucleotide A (defaults to 0.25")
	// probc := flag.Float64("prob_C", 0.25, "enter probability of nucleotide C (defaults to 0.25")
	// probg := flag.Float64("prob_G", 0.25, "enter probability of nucleotide G (defaults to 0.25")
	// probt := flag.Float64("prob_T", 0.25, "enter probability of nucleotide T (defaults to 0.25")
	seed := flag.String("seed", "random", "enter the random seed")

	flag.Parse()

	if *seed != "random" {
		entered_seed, err := strconv.ParseInt(*seed, 0, 32)
		if err != nil {
			panic(err)
		}
		// pass seed into rand seed
		rand.Seed(entered_seed)
	}
	// if no entered seed then make a random seed
	rand.Seed(time.Now().UnixNano())

	// print required number of seq
	for i := 0; i < *num_seq; i++ {
		var cur_seq string = randseq(*len_each)
		fmt.Print(">", i+1., " ")
		fmt.Println(cur_seq)
	}
}
