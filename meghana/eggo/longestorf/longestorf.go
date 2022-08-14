package main

import (
	"fmt"
	"flag"
	
	"eggo/sireadfasta"
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

func main() {
	fs := flag.String("f", "", "path  to file")
	d := flag.Bool("d", false, "double stranded translation")
	flag.Parse()
	longestorf(fs, *d)
}

var complements = map[string]string {
	"A":"T",	"T":"A",	"U":"A",	"G":"C",
	"C":"G",	"Y":"R",	"R":"Y",	"S":"S",
	"W":"W",	"K":"M",	"M":"K",	"B":"V",
	"D":"H",	"H":"D",	"V":"B",	"N":"N",
}

func reversecomp(seq string) string {
	rc := ""
	for i:=len(seq)-1; i>=0; i-- {
		nt := string(seq[i])
		rc += complements[nt]
	}
	return rc
}

func longestorf(fs *string, d bool) {
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
		longestseq := ""
		ff := it.Value()
		seq := ff.Seq
		reverse := reversecomp(seq)
		for i:=0; i<len(seq)-2; i++ {
			translation := ""
			if CODONS[seq[i:i+3]] == "M" {
				for j:=i; j<len(seq)-2; j+=3{
					aa := CODONS[seq[j:j+3]]
					if aa != "*" {
						translation += aa
					} else {
					if len(translation) > len(longestseq) {longestseq = translation}
					break
					}
				}
			}
			if d {
				if CODONS[reverse[i:i+3]] == "M" {
					translation := "M"
					for j:=i; j<len(reverse)-2; j+=3{
						aa := CODONS[reverse[j:j+3]]
						if aa != "*" {
							translation += aa
						} else {
						if len(translation) > len(longestseq) {longestseq = translation}
						break
						}
					}
				}
			}	
		}
		fmt.Println(">", ff.Id)
		if longestseq == "" {
			fmt.Println("*")
		} else {
			fmt.Println(longestseq)
		}
	}
}

