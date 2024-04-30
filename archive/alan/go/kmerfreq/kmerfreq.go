package main

import (
	"github.com/AlanAloha/read_record"
	"flag"
	"sort"
	"fmt"
	"os"
)

func main() {
	fasta := flag.String("in", "", "path to fasta file (required)")
	k := flag.Int("k", 3, "k-mer size (default: k=3)")
	j := flag.Bool("j", false, "output in json (default: tab-separated)")
	
	flag.Usage = func() {
		flagSet := flag.CommandLine
		fmt.Printf("Kmer frequency in either tab-separated or json format\n")
		order := []string{"in", "k", "j"}
		for _, name := range order {
			flag := flagSet.Lookup(name)
			fmt.Printf("-%s\n", flag.Name)
			fmt.Printf("    %s\n", flag.Usage)
		}
	}
	
	flag.Parse()
	
	if *fasta == "" {
		flag.Usage()
		os.Exit(1)
	}
	
	total := 0.0
	freq := make(map[string]float64)
	
	var kmer string
	records := read_record.Read_record(*fasta)
	for records.Next() {
		record := records.Record()
		for i := 0; i < (len(record.Seq) - *k + 1); i++ {
			kmer = record.Seq[i:i+*k]
			freq[kmer] += 1
			total += 1
		}
	}
	
	for kmer := range freq {freq[kmer] /= total}
	
	//create a slice to sort the keys
	kmers := make([]string, 0)
	for kmer := range freq {kmers = append(kmers, kmer)}
	sort.Strings(kmers)
	
	
	if *j {
		fmt.Printf("{\n")
		for i, kmer := range kmers {
			if i == len(kmers) - 1 {
				fmt.Printf("\t\"%s\": %f\n", kmer, freq[kmer])
			} else {
				fmt.Printf("\t\"%s\": %f,\n", kmer, freq[kmer])
			}
		}
		fmt.Printf("}\n")
	} else {
		for _, kmer := range kmers {
			fmt.Printf("%s\t%f\n", kmer, freq[kmer])
		}
	}

}
