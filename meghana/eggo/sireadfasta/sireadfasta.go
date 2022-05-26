package sireadfasta

import(
	"fmt"
	"bufio"
	"os"
	"io"
)
// general fasta types and functions 

type Fasta struct {
	Id string // unique identifier
	//Desc string // description
	Seq string // sequence
}

func NewFasta(id string, seq string) *Fasta {
	r := Fasta{Id: id, Seq: seq}
	return &r
}

func Print(fa *Fasta) {
    fmt.Print(">")
    fmt.Println(fa.Id)
    fmt.Println(fa.Seq)
}

// stateful iterator code

type FastaStatefulIterator struct {
    current *Fasta
    nextid *string
    scanner *FileScanner
}

type StatefulIterator interface {
    Value() Fasta //Fasta
    Next() bool //scanner.Scan()
}

func (fasta *FastaStatefulIterator) Value() *Fasta {
    return fasta.current
}

func (fasta *FastaStatefulIterator) Next() bool {
	scanner := fasta.scanner
	nextid := fasta.nextid
	current, bools := BuildFasta(scanner, nextid)
	fasta.current = current
	return bools
}


func NewFastaStatefulIterator(path *string) *FastaStatefulIterator {
	scanner := GetScannerPtr(path)
	current := NewFasta("","")
	nextid := ""
    return &FastaStatefulIterator{scanner: scanner, current: current, nextid: &nextid}
}

//to pass a scanner as a parameter

type FileScanner struct {
    io.Closer
    *bufio.Scanner
}

func GetScannerPtr(filePath *string) *FileScanner  {
    f, err := os.Open(*filePath)
    if err != nil {
        fmt.Fprint(os.Stderr, "Error opening file\n")
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    return &FileScanner{f, scanner}
}

//build current and determine next
func BuildFasta(scanner *FileScanner, nextid *string) (*Fasta, bool){
    defer scanner.Close()
    seq := ""
    for scanner.Scan() {
    
        line := scanner.Text()
        if line[0] == '>' && len(seq) != 0 {
			var fasta *Fasta

        	fasta = NewFasta(*nextid, seq)
        	*nextid = line[1:]

        	return fasta, true        	
        } else if line[0] == '>'{
        	*nextid = line[1:]
        } else {
        	seq += line
        } 
    }
    fasta := NewFasta(*nextid, seq)
   
   	if fasta.Seq != ""{
   		return fasta, true
   	} else {
    	return fasta, false
    }
}
