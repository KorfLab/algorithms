package read_record

import (
	"compress/gzip"
	"strconv"
	"strings"
	"bufio"
	"os"
)

type FastaRecord struct {
	Seq string
	Id string
}

type GffRecord struct {
	Seqid string
	Source string
	Type string
	Beg int64
	End int64
	Score float64
	Strand byte
	Phase byte
	Id string
	Parent []string
}


type fastaStatefulIterator struct {
	scanner *bufio.Scanner
	linecarrier string
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
	if it.finished {
		it.fh.Close()
		if it.gzfh != nil {
			it.gzfh.Close()
		}
		return false
	}
	seq := ""
	id := ""
	scanner := it.scanner
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if len(seq) > 0 {
				if id == "" {
					seq = it.linecarrier + seq
					it.current = &FastaRecord{Id: it.idcarrier, Seq: seq}
				} else {
					it.current = &FastaRecord{Id: id, Seq: seq}
				}
				it.idcarrier = line[1:]
				break
			} else {
				id = line[1:]
			}
		} else {
			seq += line
		}
	}
	//
	if scanner.Scan() == true {
		it.linecarrier = scanner.Text()
		return true
	} else {
		it.finished = true
		seq = it.linecarrier + seq
		if id != "" {
			it.current = &FastaRecord{Id: id, Seq: seq}
		} else {
			it.current = &FastaRecord{Id: it.idcarrier, Seq: seq}
		}
		return true
	}
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
				it.current.Parent = strings.Split(att, ",")
			}
			if strings.HasPrefix(att, "ID=") {
				it.current.Id = att	
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
