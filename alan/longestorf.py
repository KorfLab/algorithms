import argparse
from readfasta import read_record

parser = argparse.ArgumentParser(
	description='Get the protein sequence with the longest open reading frame')
parser.add_argument('fasta', type=str, metavar='<file>',
	help='input fasta file(s)')
parser.add_argument('--frame', type=str, default='single', metavar='<single/double>',
	help='input single(3) or double(6) frame translation [%(default)s]')
arg = parser.parse_args()

codons = {
"TTT" : "F", "CTT" : "L", "ATT" : "I", "GTT" : "V",
"TTC" : "F", "CTC" : "L", "ATC" : "I", "GTC" : "V",
"TTA" : "L", "CTA" : "L", "ATA" : "I", "GTA" : "V",
"TTG" : "L", "CTG" : "L", "ATG" : "M", "GTG" : "V",
"TCT" : "S", "CCT" : "P", "ACT" : "T", "GCT" : "A",
"TCC" : "S", "CCC" : "P", "ACC" : "T", "GCC" : "A",
"TCA" : "S", "CCA" : "P", "ACA" : "T", "GCA" : "A",
"TCG" : "S", "CCG" : "P", "ACG" : "T", "GCG" : "A",
"TAT" : "Y", "CAT" : "H", "AAT" : "N", "GAT" : "D",
"TAC" : "Y", "CAC" : "H", "AAC" : "N", "GAC" : "D", 
"TAA" : "*", "CAA" : "Q", "AAA" : "K", "GAA" : "E",
"TAG" : "*", "CAG" : "Q", "AAG" : "K", "GAG" : "E",
"TGT" : "C", "CGT" : "R", "AGT" : "S", "GGT" : "G",
"TGC" : "C", "CGC" : "R", "AGC" : "S", "GGC" : "G", 
"TGA" : "*", "CGA" : "R", "AGA" : "R", "GGA" : "G",
"TGG" : "W", "CGG" : "R", "AGG" : "R", "GGG" : "G"
}

def revcomp(seq):
	rc = ''
	for nt in seq:
		if nt == 'A': rc+='T'
		if nt == 'C': rc+='G'
		if nt == 'G': rc+='C'
		if nt == 'T': rc+='A'
	return rc

def findallpeps(seq):
	peps = []
	for i in range(0,len(seq)-2):
		start = (codons[seq[i:i+3]] == 'M')
		if start:
			pep = ''
			for j in range(i,len(seq)-2,3):
				codon = seq[j:j+3]
				pep += codons[codon]
				if codons[codon] == '*':
					if pep not in peps: peps.append(pep)
					break
	return peps

# Main
for idn, seq in read_record(arg.fasta):
	peps = findallpeps(seq) 
	if arg.frame == 'double':
		peps+=findallpeps(revcomp(seq))
	
	print(f'>{idn}')
	if len(peps) > 0: print(max(peps, key=len))
	else:             print('*')
	
				
			
			

