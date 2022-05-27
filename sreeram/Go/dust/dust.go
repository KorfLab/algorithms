package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

// function to check for errors and to panic if error detected
func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func entropy_finder(seq string) float64 {
	// create variables for nucleotide counts
	a := 0
	t := 0
	c := 0
	g := 0

	for i := 0; i < len(seq); i++ {
		if seq[i] == 'A' {
			a++
		}
		if seq[i] == 'C' {
			c++
		}
		if seq[i] == 'T' {
			t++
		}
		if seq[i] == 'G' {
			g++
		}
	}
	proba := float64(a) / float64(len(seq))
	probc := float64(c) / float64(len(seq))
	probg := float64(g) / float64(len(seq))
	probt := float64(t) / float64(len(seq))

	entropy := 0.0
	if proba > 0 {
		entropy -= proba * math.Log(proba)
	}
	if probc > 0 {
		entropy -= probc * math.Log(probc)
	}
	if probg > 0 {
		entropy -= probg * math.Log(probg)
	}
	if probt > 0 {
		entropy -= probt * math.Log(probt)
	}

	entropy = entropy / math.Log(2)
	return entropy
}

func main() {

	// get the command line inputs
	fasta_file := flag.String("fasta", "", "enter path to fasta file")
	window_size := flag.Int("window", 11, "enter window size")
	threshold := flag.Float64("t", 1.2, "enter the entropy threshold")
	lower_case := flag.Bool("lc", false, "enter if mask with lower case")

	flag.Parse()

	// if there are less arguments given than needed then prompt user about what the inputs are
	if flag.NArg() != 0 {
		flag.Usage()
	}

	// check for any errors when file is opened
	fasta, err := os.Open(*fasta_file)
	check_error(err)
	// no errors

	// make some variables to store current running seqence, total sequence and the number of seqences there are (for iterating purposes)
	cur_seq := ""
	total_seq := []string{}
	num_seq := []string{}

	// create a reader for reading the fasta file in chunks
	fasta_scan := bufio.NewScanner(fasta)

	// make as own function
	for fasta_scan.Scan() {
		// read in rach chunk from the fasta file
		cur_line := ""
		cur_line = fasta_scan.Text()

		// check if cur line starts with > so we can keep track of the sequence
		// this can be made into its own function later
		if strings.HasPrefix(cur_line, ">") {
			if len(num_seq) > 0 {
				total_seq = append(total_seq, cur_seq)
				cur_seq = ""
			}
			num_seq = append(num_seq, cur_line[1:])
		} else {
			cur_seq += cur_line
		}
	}
	total_seq = append(total_seq, cur_seq)

	// loop to calculate entropy and mask the sequence (also masks as lower case if flag is true)
	for i := 0; i < len(num_seq); i++ {
		seq_used := ""
		seq_used = total_seq[i]

		for j := 0; j < len(seq_used)-*window_size+1; j++ {
			cur_entropy := 0.0
			cur_entropy = entropy_finder(seq_used[j : j+*window_size])

			if cur_entropy < *threshold {
				if *lower_case {
					cur_seq = cur_seq[:j+*window_size/2] + strings.ToLower(string(cur_seq[j+*window_size/2])) + cur_seq[j+*window_size/2+1:]
				} else {
					cur_seq = cur_seq[:j+*window_size/2] + "X" + cur_seq[j+*window_size/2+1:]
				}
			}
		}
		fmt.Print("##", num_seq[i], "\n")
		fmt.Println(cur_seq)
	}
	fasta.Close()
}
