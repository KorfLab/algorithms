package main

import (
	"os"
	"fmt"
	"math"
	"flag"
	"strings"
	"github.com/AlanAloha/si_read_record"
)

func get_entropy (freqs map[string]float64) float64 {
	tot := 0.0
	h := 0.0
	for _, freq := range freqs {tot += freq}
	for _, freq := range freqs {
		prop := freq / tot
		if prop > 0 {h -= prop * math.Log(prop)}
	}
	
	h /= math.Log(2)
	return h
}

func main() {
	fasta := flag.String("in", "", "path to fasta file (required)")
	w := flag.Int("w", 11, "window size (default: 11)")
	t := flag.Float64("t", 1.5, "entropy threshold (default: 1.1)")
	s := flag.Bool("s", false, "mask to lowercase, (default mast to N)")
	flag.Parse()
	
	if *fasta == "" {
		flag.Usage()
		os.Exit(1)
	}

	records := si_read_record.Read_record(*fasta)
	for records.Next() {
		record := records.Record()
		seq := strings.ToUpper(record.Seq)
		
		mask := seq
		freqs := make(map[string]float64)
		for i := 0; i < len(seq)-*w+1; i++ {
			if i == 0 {
				sub := seq[i:i+*w]
				freqs["A"] = float64(strings.Count(sub, "A"))
				freqs["C"] = float64(strings.Count(sub, "C"))
				freqs["G"] = float64(strings.Count(sub, "G"))
				freqs["T"] = float64(strings.Count(sub, "T"))
			} else {
				previous := string(seq[i-1])
				current := string(seq[i+*w-1])
				freqs[previous] -= 1
				freqs[current] += 1
			}
			h := get_entropy(freqs)
			if h < *t {
				pos := i + int(*w/2)
				if *s {
					mask = mask[:pos] + strings.ToLower(string(mask[pos])) + mask[pos+1:]
				} else {
					mask = mask[:pos] + "N" + mask[pos+1:]
				}
			}
		}
		
		fmt.Printf(">%s\n", record.Id)
		for i := 0; i < len(mask); i+=80 {
			if i+80 < len(mask) {fmt.Println(mask[i:i+80])} else {fmt.Println(mask[i:])}
		}
	}
}
