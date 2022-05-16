package read_fasta

import (
	//"fmt"
	"os"
	"bufio"
	"strings"
)

type Read struct {
	Id string
	Seq string
}


func Read_fasta(ff string, callBack func(Read)) {
	fh, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	
	seq := ""
	id := ""
	var read Read
	
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if len(seq) > 0 {
				read = Read{Id: id, Seq: seq}
				callBack(read)
				id = line[1:]
				seq = ""
			} else {
				id = line[1:]
			}
		 } else {
		 	seq += line
		 }
	}
	read = Read{Id: id, Seq: seq}
	callBack(read)
}
