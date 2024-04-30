package read_record

import (
	"compress/gzip"
	"strconv"
	"strings"
	"bufio"
	"os"
)

// Fasta reader
type FastaRecord struct {
	Seq string
	Id string
}

type fastaStatefulIterator struct {
	scanner *bufio.Scanner
	gzfh *gzip.Reader
	idcarrier string
	current *FastaRecord
	finished bool
	fh *os.File
}

func Read_fasta(ff string) *fastaStatefulIterator {
	fh, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	
	si := &fastaStatefulIterator{
		finished: false,
		fh: fh}
		
	scanner := bufio.NewScanner(fh)
	if strings.HasSuffix(ff, ".gz") {
		gzfh, err := gzip.NewReader(fh)
		si.gzfh = gzfh
		if err != nil {
			panic(err)
		}
		scanner = bufio.NewScanner(gzfh)
	}
	si.scanner = scanner
	return si
}

func (it *fastaStatefulIterator) Record() *FastaRecord {
	return it.current
}

func (it *fastaStatefulIterator) Next() bool {
	scanner := it.scanner
	not_empty := true
	seq := ""
	id := ""
	if it.finished {
		return false
	}
	for true {
		if scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, ">") {
				if len(seq) == 0 {
					id = line[1:]
				} else {
					if len(id) > 0 {
						it.current = &FastaRecord{Id: id, Seq: seq}
					} else {
						it.current = &FastaRecord{Id: it.idcarrier, Seq: seq}
					}
					it.idcarrier = line[1:]
					break
				}
			} else {
				seq += line
			}
		} else {
			if len(seq) > 0 {
				if len(it.idcarrier) > 0 {
					it.current = &FastaRecord{Id: it.idcarrier, Seq: seq}
				} else {
					it.current = &FastaRecord{Id: id, Seq: seq}
				}
			} else {
				not_empty = false
			}
			it.fh.Close()
			if it.gzfh != nil {
				it.gzfh.Close()
			}
			it.finished  = true
			break
		}
	}
	
	return not_empty
}

// Gff reader

type GffRecord struct {
	Seqid string
	Source string
	Type string
	Beg int64
	End int64
	Score float64
	Strand byte
	Phase byte
	ID string
	Parent []string
}

type gffStatefulIterator struct {
	scanner *bufio.Scanner
	gzfh *gzip.Reader
	current *GffRecord
	fh *os.File
}

func Read_gff(gff string) *gffStatefulIterator{
	fh, err := os.Open(gff)
	if err != nil {
		panic(err)
	}
	
	si := &gffStatefulIterator{fh: fh}
	
	scanner := bufio.NewScanner(fh)
	if strings.HasSuffix(gff, ".gz") {
		gzfh, err := gzip.NewReader(fh)
		si.gzfh = gzfh
		if err != nil {
			panic(err)
		}
		scanner = bufio.NewScanner(gzfh)
	}
	
	si.scanner = scanner
	return si
}

func (it *gffStatefulIterator) Record() *GffRecord {
	return it.current
}

func (it *gffStatefulIterator) Next() bool {
	scanner := it.scanner
	if scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		fields := strings.Split(line, "\t")
		
		it.current = &GffRecord{
			Seqid: fields[0],
			Source: fields[1],
			Type: fields[2]}
		strand := []byte(fields[6])
		phase := []byte(fields[7])
		it.current.Strand = strand[0]
		it.current.Phase = phase[0]
		it.current.Beg, _ = strconv.ParseInt(fields[3], 10, 32)
		it.current.End, _ = strconv.ParseInt(fields[4], 10, 32)
		it.current.Score, _ = strconv.ParseFloat(fields[5], 64)
		
		atts := strings.Split(fields[8], ";")
		for _, att := range atts {
			if strings.HasPrefix(att, "Parent=") {
				parents := strings.ReplaceAll(att, "Parent=", "")
				it.current.Parent = strings.Split(parents, ",")
			}
			if strings.HasPrefix(att, "ID=") {
				id := strings.ReplaceAll(att, "ID=", "")
				it.current.ID = id
			}
		}
		
		return true
	} else {
		it.fh.Close()
		if it.gzfh != nil {
			it.gzfh.Close()
		}
		return false
	}
}
