package main

import (
	"github.com/AlanAloha/si_read_record"
	"flag"
	"fmt"
	"os"
)

type sm_alignment struct{
	qrid string
	dbid string
	qraln string
	dbaln string
	aln string
	qrs int
	qre int
	dbs int
	dbe int
	score int
}

func display_scoremat (mat [][]int) {
	for i := range mat {
		for j := range mat[i] {fmt.Printf("%d ", mat[i][j])}
		fmt.Printf("\n")
	}
}
func display_tracemat (mat [][]string) {
	for i := range mat {
		for j := range mat[i] {fmt.Printf("%s ", mat[i][j])}
		fmt.Printf("\n")
	}
}

func get_max (mat [][]int) (int, int, int) {
	max := 0
	maxi := 0
	maxj := 0
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] > max {
				max = mat[i][j]
				maxi = i
				maxj = j
			}
		}
	}
	return max, maxi, maxj
}

func sm_align (qrr *si_read_record.Record, dbr *si_read_record.Record, m int, n int, g int) *sm_alignment {
	qrid := qrr.Id
	dbid := dbr.Id
	qrseq := qrr.Seq
	dbseq := dbr.Seq
	qraln := ""
	dbaln := ""
	aln := ""
	qrlen := len(qrseq)
	dblen := len(dbseq)
	
	scoremat := make([][]int, qrlen+1)
	for i := range scoremat {
		scoremat[i] = make([]int, dblen+1)
	}
	tracemat := make([][]string, qrlen+1)
	for i := range tracemat {
		tracemat[i] = make([]string, dblen+1)
		for j := range tracemat[i] {
			tracemat[i][j] = "N"
		}
	}
	
	for i := range qrseq {
		for j := range dbseq {
			top := scoremat[i][j+1] + g
			lft := scoremat[i+1][j] + g
			var dgn int
			if qrseq[i] == dbseq[j] {
				dgn = scoremat[i][j] + m
			} else {
				dgn = scoremat[i][j] + n
			}
			
			if dgn > top && dgn > lft && dgn > 0{
				scoremat[i+1][j+1] = dgn
				tracemat[i+1][j+1] = "D"
			} else if top > lft && top > 0 {
				scoremat[i+1][j+1] = top
				tracemat[i+1][j+1] = "T"
			} else if lft > 0 {
				scoremat[i+1][j+1] = lft
				tracemat[i+1][j+1] = "L"
			}
		}
	}
	//display_scoremat(scoremat)
	//display_tracemat(tracemat)
	max, maxi, maxj := get_max(scoremat)
	curri := maxi
	currj := maxj
	if max == 0 {
		return &sm_alignment {
			qrid: qrid,
			dbid: dbid,
			qraln: "",
			dbaln: "",
			qrs: 0,
			qre: 0,
			dbs: 0,
			dbe: 0,
			score: 0}
	} else {
		for scoremat[curri][currj] != 0 {
			if tracemat[curri][currj] == "D" {
				qraln = string(qrseq[curri-1]) + qraln
				dbaln = string( dbseq[currj-1]) + dbaln
				aln = "|" + aln
				curri--
				currj--
			} else if tracemat[curri][currj] == "L" {
				qraln = "-" + qraln
				dbaln = string(dbseq[currj-1]) + dbaln
				aln = " " + aln
				currj--
			} else if tracemat[curri][currj] == "T" {
				qraln = string(qrseq[curri-1]) + qraln
				dbaln = "-" + dbaln
				aln = " " + aln
				curri--
			}
			
		}
		return &sm_alignment {
			qrid: qrid,
			dbid: dbid,
			qraln: qraln,
			dbaln: dbaln,
			aln : aln,
			qrs: curri+1,
			qre: maxi,
			dbs: currj+1,
			dbe: maxj,
			score: max}
	}
}

func main() {
	in := flag.String("in", "", "path to fasta file (required)")
	db := flag.String("db", "", "path to database sequence or STDIN (required)")
	m := flag.Int("m", 1, "match score (default: 1)")
	n := flag.Int("n", -1, "mismatch score (default: -1)")
	g := flag.Int("g", -2, "gap score (default: -2)")
	t := flag.Bool("t", false, "output in tabular format")
	flag.Parse()
	
	if *in == "" || *db == ""{
		flag.Usage()
		os.Exit(1)
	}
	
	qr_records := si_read_record.Read_record(*in)
	var qr_record *si_read_record.Record
	if qr_records.Next() {qr_record = qr_records.Record()}
	db_records := si_read_record.Read_record(*db)
	for db_records.Next() {
		db_record := db_records.Record()
		alignment := sm_align(qr_record, db_record, *m, *n, *g)
		if *t {
			fmt.Printf("%s\t%s\t%d\t%d\t%d\t%d\t%d\n", alignment.qrid, alignment.dbid,
			alignment.score, alignment.qrs, alignment.qre, alignment.dbs, alignment.dbe)
		} else {
			fmt.Printf("Query: %s\n", alignment.qrid)
			fmt.Printf("Sbjct: %s\n", alignment.dbid)
			fmt.Printf("Score: %d\n\n", alignment.score)
			fmt.Printf("%d\t%s\t%d\n", alignment.qrs, alignment.qraln, alignment.qre)
			fmt.Printf(" \t%s\n", alignment.aln)
			fmt.Printf("%d\t%s\t%d\n\n", alignment.dbs, alignment.dbaln, alignment.dbe)
		}
	}
}
