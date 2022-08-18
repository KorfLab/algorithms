package main

import(
	"testing"
)

func TestStr_Reverse(t *testing.T) {
	t.Parallel()
	str_reverse_tests := map[string]struct {
		input string
		output string
	} {
		"empty string": {
			input: "",
			output: "",
		},
		"1 bp": {
			input: "A",
			output: "A",
		},
		"multiple basepairs": {
			input: "ALaN",
			output: "NaLA",
		},
	}
	for name, test := range str_reverse_tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := str_reverse(test.input), test.output; got != expected {
				t.Errorf("str_reverse(%s) returned %s; expected %s\n", test.input, got, expected)
			}
		})
	}
}

func TestLongestPep(t *testing.T) {
	t.Parallel()
	longestPep_tests := map[string]struct {
		input string
		output string
	} {
		"empty string": {
			input: "",
			output: "*",
		},
		"no orf": {
			input: "GACCTTCACTTGCCGCGTAGCAAGCTGGCCGGGCAAACGGGTATATCTGCCGCGACG" + 						"AAAAGCCAAATGGCCATCTGTCGGTTTGTTCCGGCCCAGACGC",
			output:"*",
		},
		"contain orf": {
			input: "CACTTTGCTTTAAACGGCCGCGGATATCTATGTGGGCGGTGTACGTTGCAAGCACT" + 						"TTCAGGCGCGCTTCGGCTCAGGCCCCCTAGCCAACTTCACAGGG",
			output: "MWAVYVASTFRRASAQAP",
		},
	}
	
	for name, test := range longestPep_tests {
		test := test
		t.Run(name, func(t *testing.T) {
			if got, expected := longestPep(test.input), test.output; got != expected {
				t.Errorf("longestPep(%s) returned %s; expected %s\n", test.input, got, expected)
			}
		})
	}
}
