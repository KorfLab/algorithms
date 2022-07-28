package main

import (
	"os"
	"fmt"
	"math"
	"flag"
	"strings"
	"github.com/AlanAloha/si_read_record"
)


func seq_entropy (seq string) float64 {
	a := float64(strings.Count(seq, "A"))
	c := float64(strings.Count(seq, "C"))
	g := float64(strings.Count(seq, "G"))
	t := float64(strings.Count(seq, "T"))
	
	tot := a + c + g + t
	
	pa := a / tot
	pc := c / tot
	pg := g / tot
	pt := t / tot
	
	h := 0.0
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
		
		mask := seq
		for i := 0; i < len(seq)-*w+1; i++ {
			sub := seq[i:i+*w]
			h := seq_entropy(sub)
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
		//fmt.Println(mask)
		for i := 0; i < len(mask); i+=80 {
			if i+80 < len(mask) {
				fmt.Println(mask[i:i+80])
			} else {
				fmt.Println(mask[i:])
			}
		}
	}
}
