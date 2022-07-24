package si_read_record

import (
	"os"
	"bufio"
	"strings"
	"compress/gzip"
)

type Record struct {
	Id string
	Seq string
}

type StatefulIterator interface {
	Value() Record
	Next() bool
}

type recordStatefulIterator struct {
	current *Record
	scanner *bufio.Scanner
	fh *os.File
	gzfh *gzip.Reader
	idcarrier string
	linecarrier string
	finished bool
}

func NewRecordStatefulIterator(ff string) *recordStatefulIterator {
	si := &recordStatefulIterator{current: &Record{Id: "", Seq: ""}, finished: false}
	
	fh, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fh)
	if strings.HasSuffix(ff, ".gz") {
		gzfh, err := gzip.NewReader(fh)
		si.gzfh = gzfh
		if err != nil {
			panic(err)
		}
		scanner = bufio.NewScanner(gzfh)
	}
	si.fh = fh
	si.scanner = scanner
	return si
	/*
	return &recordStatefulIterator {
		current: &Record{Id: "", Seq: ""},
		scanner: scanner,
		fh: fh,
		gzfh: gzfh,
		finished: false}
	*/
}

func (it *recordStatefulIterator) Value() *Record {
	return it.current
}

func (it *recordStatefulIterator) Next() bool {
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
					it.current = &Record{Id: it.idcarrier, Seq: seq}
				} else {
					it.current = &Record{Id: id, Seq: seq}
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
		it.current = &Record{Id: it.idcarrier, Seq: seq}
		return true
	}
}


