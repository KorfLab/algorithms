package si_read_record

import (
	"compress/gzip"
	"strings"
	"bufio"
	"os"
)

type Record struct {
	Seq string
	Id string
}

type StatefulIterator interface {
	Record() Record
	Next() bool
}

type recordStatefulIterator struct {
	scanner *bufio.Scanner
	linecarrier string
	gzfh *gzip.Reader
	idcarrier string
	current *Record
	finished bool
	fh *os.File
}

func Read_record(ff string) *recordStatefulIterator {
	fh, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	
	si := &recordStatefulIterator{
		current: &Record{Id: "", Seq: ""},
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

func (it *recordStatefulIterator) Record() *Record {
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
		if id != "" {
			it.current = &Record{Id: id, Seq: seq}
		} else {
			it.current = &Record{Id: it.idcarrier, Seq: seq}
		}
		return true
	}
}


