package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"flag"
	"fmt"
	"os"
)

func rand_seq(length int, a float64, c float64, g float64, t float64) string{
	seq := make([]string, length)
	for i := 0; i < length; i++ {
		rf := rand.Float64()
		if rf < a {
			seq[i] = "A"
		} else if rf < a + c {
			seq[i] = "C"
		} else if rf < a + c + g {
			seq[i] = "G"
		} else {
			seq[i] = "T"
		}
	}
	return strings.Join(seq,"")
}

func main() {
	num_seq := flag.Int("num", 10, "number of sequences")
	len_seq := flag.Int("len", 100, "length of sequence")
	prefix := flag.String("prefix", "id", "prefix for identifier")
	seed := flag.String("seed", "random", "random seed" )
	a := flag.Float64("a", 0.25, "weight of A")
	c := flag.Float64("c", 0.25, "weight of C")
	g := flag.Float64("g", 0.25, "weight of G")
	t := flag.Float64("t", 0.25, "weight of T")
	
	flag.Usage = func() {
		flagSet := flag.CommandLine
		fmt.Printf("Generate random sequence and print to STDOUT\n")
		order := []string{"num","len","a", "c", "g", "t", "prefix", "seed"}
		for _, name := range order {
			flag := flagSet.Lookup(name)
			fmt.Printf("-%s\n", flag.Name)
			fmt.Printf("    %s\n", flag.Usage)
		}
	}
	
	flag.Parse()
	
	if *a + *c + *g + *t - 1.0 > 0.00001 {
		fmt.Printf("User input:\nA: %f\nC: %f\nG: %f\nT: %f\n", *a, *c, *g, *t)
		os.Exit(1)
	}

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
		
		
		seq = rand_seq(*len_seq, *a, *c, *g, *t)

		for j := 0; j < len(seq); j += 80 {
			if j + 80 < len(seq) {
				fmt.Println(seq[j:j+80])
			} else {
				fmt.Println(seq[j:])
			} 
		}
	}
	
	
}

