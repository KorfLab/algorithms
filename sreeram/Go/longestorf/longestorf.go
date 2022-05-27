package main

import (
	"flag"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func find_starts(seq string) []int {
	start_indices := []int{}
	cur_start := 0

	for cur_start < len(seq)-2 {
		cur_start = strings.Index(seq[cur_start+1:], "ATG")
		if cur_start == -1 {
			break
		}
		start_indices = append(start_indices, cur_start)
	}

	return start_indices
}

func find_stop(seq string) int {
	var i = 0
	var stop_codons = []string{"TAG", "TAA", "TGA"}
	for i < len(seq)-2 {
		if !(slices.Contains(stop_codons, seq[i:i+3])) {
			i = i + 3
		}
	}

	if i < len(seq)-2 {
		return i + 3
	}
	return -1
}

func get_orfs(seq string) [][]int {
	var start_indices = find_starts(seq)
	var stop_indices = []int{}

	for start := range start_indices {
		var stop = find_stop(seq[start:])
		stop = start + stop
		stop_indices = append(stop_indices, stop)
	}

	var orf_pairs = [][]int{}

	for i := range start_indices {
		var cur_pair = []int{start_indices[i], stop_indices[i]}
		orf_pairs = append(orf_pairs, cur_pair)
	}

	return orf_pairs
}

func seq_comp(seq string) string {
	seq = strings.ReplaceAll(seq, "A", "T")
	seq = strings.ReplaceAll(seq, "T", "A")
	seq = strings.ReplaceAll(seq, "C", "G")
	seq = strings.ReplaceAll(seq, "G", "C")
	return seq
}

func longest_orf(seq_orfs [][]int) (int, int, int) {
	var len_longest = -1
	var longest_start = -1
	var longest_stop = -1

	for orf := range seq_orfs {
		var cur_orf = seq_orfs[orf]
		var start = cur_orf[0]
		var stop = cur_orf[1]

		if (stop - start) > len_longest {
			len_longest = stop - start
			longest_start = start
			longest_stop = stop
		}
	}
	return len_longest, longest_start, longest_stop
}

func main() {
	ff := flag.String("fasta_file", "", "enter path to fasta file (required input)")
	// frame := flag.Bool("frame", false, "enter if you want to include the reverse complement (defaults to false)")

	flag.Parse()
	if *ff == "" {
		println("Please enter the path to fasta file")
		os.Exit(1)
	}
	// read fasta file (still have to do this) and call all functions
}
