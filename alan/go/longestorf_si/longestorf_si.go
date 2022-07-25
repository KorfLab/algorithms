package main

import (
	"github.com/AlanAloha/si_read_record"
	"strings"
	"flag"
	"fmt"
	"os"
)

var CODONS = map[string]string {
	"AAA" : "K",	"AAC" : "N",	"AAG" : "K",	"AAT" : "N",
	"AAR" : "K",	"AAY" : "N",	"ACA" : "T",	"ACC" : "T",
	"ACG" : "T",	"ACT" : "T",	"ACR" : "T",	"ACY" : "T",
	"ACK" : "T",	"ACM" : "T",	"ACW" : "T",	"ACS" : "T",
	"ACB" : "T",	"ACD" : "T",	"ACH" : "T",	"ACV" : "T",
	"ACN" : "T",	"AGA" : "R",	"AGC" : "S",	"AGG" : "R",
	"AGT" : "S",	"AGR" : "R",	"AGY" : "S",	"ATA" : "I",
	"ATC" : "I",	"ATG" : "M",	"ATT" : "I",	"ATY" : "I",
	"ATM" : "I",	"ATW" : "I",	"ATH" : "I",	"CAA" : "Q",
	"CAC" : "H",	"CAG" : "Q",	"CAT" : "H",	"CAR" : "Q",
	"CAY" : "H",	"CCA" : "P",	"CCC" : "P",	"CCG" : "P",
	"CCT" : "P",	"CCR" : "P",	"CCY" : "P",	"CCK" : "P",
	"CCM" : "P",	"CCW" : "P",	"CCS" : "P",	"CCB" : "P",
	"CCD" : "P",	"CCH" : "P",	"CCV" : "P",	"CCN" : "P",
	"CGA" : "R",	"CGC" : "R",	"CGG" : "R",	"CGT" : "R",
	"CGR" : "R",	"CGY" : "R",	"CGK" : "R",	"CGM" : "R",
	"CGW" : "R",	"CGS" : "R",	"CGB" : "R",	"CGD" : "R",
	"CGH" : "R",	"CGV" : "R",	"CGN" : "R",	"CTA" : "L",
	"CTC" : "L",	"CTG" : "L",	"CTT" : "L",	"CTR" : "L",
	"CTY" : "L",	"CTK" : "L",	"CTM" : "L",	"CTW" : "L",
	"CTS" : "L",	"CTB" : "L",	"CTD" : "L",	"CTH" : "L",
	"CTV" : "L",	"CTN" : "L",	"GAA" : "E",	"GAC" : "D",
	"GAG" : "E",	"GAT" : "D",	"GAR" : "E",	"GAY" : "D",
	"GCA" : "A",	"GCC" : "A",	"GCG" : "A",	"GCT" : "A",
	"GCR" : "A",	"GCY" : "A",	"GCK" : "A",	"GCM" : "A",
	"GCW" : "A",	"GCS" : "A",	"GCB" : "A",	"GCD" : "A",
	"GCH" : "A",	"GCV" : "A",	"GCN" : "A",	"GGA" : "G",
	"GGC" : "G",	"GGG" : "G",	"GGT" : "G",	"GGR" : "G",
	"GGY" : "G",	"GGK" : "G",	"GGM" : "G",	"GGW" : "G",
	"GGS" : "G",	"GGB" : "G",	"GGD" : "G",	"GGH" : "G",
	"GGV" : "G",	"GGN" : "G",	"GTA" : "V",	"GTC" : "V",
	"GTG" : "V",	"GTT" : "V",	"GTR" : "V",	"GTY" : "V",
	"GTK" : "V",	"GTM" : "V",	"GTW" : "V",	"GTS" : "V",
	"GTB" : "V",	"GTD" : "V",	"GTH" : "V",	"GTV" : "V",
	"GTN" : "V",	"TAA" : "*",	"TAC" : "Y",	"TAG" : "*",
	"TAT" : "Y",	"TAR" : "*",	"TAY" : "Y",	"TCA" : "S",
	"TCC" : "S",	"TCG" : "S",	"TCT" : "S",	"TCR" : "S",
	"TCY" : "S",	"TCK" : "S",	"TCM" : "S",	"TCW" : "S",
	"TCS" : "S",	"TCB" : "S",	"TCD" : "S",	"TCH" : "S",
	"TCV" : "S",	"TCN" : "S",	"TGA" : "*",	"TGC" : "C",
	"TGG" : "W",	"TGT" : "C",	"TGY" : "C",	"TTA" : "L",
	"TTC" : "F",	"TTG" : "L",	"TTT" : "F",	"TTR" : "L",
	"TTY" : "F",	"TRA" : "*",	"YTA" : "L",	"YTG" : "L",
	"YTR" : "L",	"MGA" : "R",	"MGG" : "R",	"MGR" : "R",
}

func str_reverse(str string) string {
	rev := []rune(str)
	for i, j := 0, len(rev) - 1; i < j; i, j = i + 1, j - 1{
		rev[i], rev[j] = rev[j], rev[i]
	}	
	
	return string(rev)	
}

func revcomp(seq string) string {
	from := "ACGTRYMKWSBDHVN"
	to   := "tgcayrkmwsvhdbn"
	
	revcomp := strings.ToUpper(seq)
	for i := 0; i < len(from); i++ {
		revcomp = strings.ReplaceAll(revcomp, string(from[i]), string(to[i]))
	}
	revcomp = strings.ToUpper(revcomp)
	
	return str_reverse(revcomp)
}

func longestPep(seq string) string {
	longest_pep := ""
	for i := 0; i < len(seq) - 2; i++ {
		aa, is_codon := CODONS[seq[i:i+3]]
		if is_codon && aa == "M" {
			pep := ""
			for j := i; j < len(seq) - 2; j += 3 {
				aa, is_codon := CODONS[seq[j:j+3]]
				if !is_codon {
					pep += "X"
				} else if aa == "*" {
					if len(pep) > len(longest_pep) {
						longest_pep = pep	
				 	}
					break
				} else {
					pep += aa
				}
			}
		}
	}
	
	if len(longest_pep) == 0 {longest_pep += "*"}
	return longest_pep
}

func main() {
	fasta := flag.String("in", "", "path to fasta file (required)")
	r := flag.Bool("r", false, "also translate reverse complement strand (defaut: false)")
	flag.Parse()
	
	if *fasta == "" {
		flag.Usage()
		os.Exit(1)
	}
	
	records := si_read_record.Read_record(*fasta)
	for records.Next() {
		record := records.Record()
		fmt.Printf(">%s\n",record.Id)
		seq := record.Seq
		if *r {
			revseq := revcomp(seq)
			if len(longestPep(seq)) > len(longestPep(revseq)) {
				fmt.Println(longestPep(seq))
			} else {
				fmt.Println(longestPep(revseq))
			}
		} else {
			fmt.Println(longestPep(seq))
		}
	}

}
