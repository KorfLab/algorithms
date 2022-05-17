package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func make_seq(length int) string{
	bases := "ACGT"
	seq := ""
	for i := 0; i < length; i++ {
		seq += string(bases[rand.Intn(4)])
	}
	return seq
}

func main() {
	num_seq := flag.Int("num", 10, "number of sequences")
	len_seq := flag.Int("len", 100, "length of sequence")
	prefix := flag.String("prefix", "id", "prefix for identifier")
	seed := flag.String("seed", "random", "random seed" )
	flag.Parse()
	
	if *seed != "random" {
		seed, err:= strconv.ParseInt(*seed, 0, 64)
		if err != nil {
		panic(err)
	}
		rand.Seed(seed)
	} else {
		rand.Seed(time.Now().UnixNano())
	}
	
	
	var seq string
	for i := 0; i < *num_seq; i++ {
		fmt.Printf(">%s%d\n", *prefix, i+1)
		
		seq = make_seq(*len_seq)
		for j := 0; j < len(seq); j += 80 {
			if j + 80 < len(seq) {
				fmt.Println(seq[j:j+80])
			} else {
				fmt.Println(seq[j:])
			} 
		}
	}
	
	
}

