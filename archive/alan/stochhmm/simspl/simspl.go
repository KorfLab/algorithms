package main

import (
	"errors"
	"fmt"
	"flag"
	"github.com/AlanAloha/read_record"
)

func nonempty(path string, arg_name string) error {
	message := fmt.Sprintf("missing input %s", arg_name)
	if path == "" {
		return errors.New(message)
	}
	
	return nil
}

func check_input(path string, arg_name string)  {
	err := nonempty(path, arg_name)
	if err != nil {
		flag.Usage()
		panic(err)
	}
}

//func read_jhmm(jhmm_fp string) 

func locate_gene_start(gff_fp string) (int64, int64, error) {
	gff_iterator := read_record.Read_gff(gff_fp)
	for gff_iterator.Next() {
		record := gff_iterator.Record()
		
		if record.Source != "WormBase" {
			continue
		}
		
		if record.Type == "mRNA" {
			start := record.Beg - 1
			end   := record.End
			
			return start, end, nil
		}
	}
	
	return -1, -1, errors.New("Can\\'t locate gene beginning and ending position")
}

func main() {
	fasta := flag.String("fa", "", "path to FASTA file")
	jhmm := flag.String("jhmm", "", "path to JHMM file")
	iterations := flag.Int("it", 500, "number of iterations [500]")
	top := flag.Int("top", 10, "out put top n isoforms [10]")
	ng := flag.Bool("ng", false, "no genomic state")
	gff := flag.String("gff", "", "path to GFF file, used to locate genomic-exon boundary (require when ng tag is on)")
	all := flag.Bool("all", false, "print all isoforms simulated")
	flag.Parse()
	
//	check_input(*fasta, "fa")
//	check_input(*jhmm, "jhmm")
	
	fmt.Printf("%s, %s, %d, %d, %t, %s, %t\n", *fasta, *jhmm, *iterations, *top, *ng, *gff, *all)
	
	var seq string
	
	fasta_reader := read_record.Read_fasta(*fasta)
	for fasta_reader.Next() {
		record := fasta_reader.Record()
		seq = record.Seq
	}

	_=seq
	

	if *ng {
		check_input(*gff, "gff")
		mrna_start, mrna_end, err := locate_gene_start(*gff)
		if err != nil {
			panic(err)
		}
		seq = seq[mrna_start:mrna_end]
	}
	
//	states, orders, trans, emiss, inits, terms = read_jhmm(*gff)
	

}
