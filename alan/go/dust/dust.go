package main

import (
	"os"
	"fmt"
	"math"
	"flag"
	"strings"
	"github.com/AlanAloha/si_read_record"
)

func get_entropy (a float64, c float64, g float64, t float64) float64 {
	h := 0.0
	tot := a + c + g + t
	
	pa := a / tot
	pc := c / tot
	pg := g / tot
	pt := t / tot

	if pa > 0 {h -= pa * math.Log(pa)}
	if pc > 0 {h -= pc * math.Log(pc)}
	if pg > 0 {h -= pg * math.Log(pg)}
	if pt > 0 {h -= pt * math.Log(pt)}
	
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
		
		mask_list := strings.Split(seq, "")
		var freqs [4] float64
		for i := 0; i < len(seq)-*w+1; i++ {
			if i == 0 {
				sub := seq[i:i+*w]
				freqs[0] = float64(strings.Count(sub, "A"))
				freqs[1] = float64(strings.Count(sub, "C"))
				freqs[2] = float64(strings.Count(sub, "G"))
				freqs[3] = float64(strings.Count(sub, "T"))
			} else {
				previous := string(seq[i-1])
				current := string(seq[i+*w-1])
				if previous == "A" {
					freqs[0] -= 1.0
				} else if previous == "C" {
					freqs[1] -= 1.0
				} else if previous == "G" {
					freqs[2] -= 1.0
				} else {
					freqs[3] -= 1.0
				}
				
				if current == "A" {
					freqs[0] += 1.0
				} else if current == "C" {
					freqs[1] += 1.0
				} else if current == "G" {
					freqs[2] += 1.0
				} else {
					freqs[3] += 1.0
				}
			}
			h := get_entropy(freqs[0], freqs[1], freqs[2], freqs[3])
			if h < *t {
				pos := i + int(*w/2)
				if *s {
					mask_list[pos] = strings.ToLower(string(mask_list[pos]))
				} else {
					mask_list[pos] = "N"
				}
			}
		}
		mask := strings.Join(mask_list,"")
		fmt.Printf(">%s\n", record.Id)
		for i := 0; i < len(mask); i+=80 {
			if i+80 < len(mask) {fmt.Println(mask[i:i+80])} else {fmt.Println(mask[i:])}
		}
	}
}
