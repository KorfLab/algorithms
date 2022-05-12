package read_fasta

import (
	"os"
	"bufio"
	"strings"
)

type Read struct {
	Id string
	Seq string
}


func Read_fasta(ff string) chan Read{
	fh, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	
	seq := ""
	id := ""
	read := make(chan Read)
	
	scanner := bufio.NewScanner(fh)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, ">") {
				if len(seq) > 0 {
					read <- Read{Id: id, Seq: seq}
					id = line[1:]
					seq = ""
				} else {
					id = line[1:]
				}
			 } else {
			 	seq += line
			 }
		}
	}()
	
	return read
}


