import argparse
from seqlib import read_fasta

parser = argparse.ArgumentParser(description='Get the protein sequence with the longest open reading frame')
parser.add_argument('fasta', type=str, metavar='<file>', help='input fasta file(s)')
parser.add_argument('--frame', type=str, default='single',
	metavar='<single/double>', help='input single(3) or double(6) frame translation [%(default)s]')
arg = parser.parse_args()

codons = {
"UUU" : "F", "CUU" : "L", "AUU" : "I", "GUU" : "V",
"UUC" : "F", "CUC" : "L", "AUC" : "I", "GUC" : "V",
"UUA" : "L", "CUA" : "L", "AUA" : "I", "GUA" : "V",
"UUG" : "L", "CUG" : "L", "AUG" : "M", "GUG" : "V",
"UCU" : "S", "CCU" : "P", "ACU" : "T", "GCU" : "A",
"UCC" : "S", "CCC" : "P", "ACC" : "T", "GCC" : "A",
"UCA" : "S", "CCA" : "P", "ACA" : "T", "GCA" : "A",
"UCG" : "S", "CCG" : "P", "ACG" : "T", "GCG" : "A",
"UAU" : "Y", "CAU" : "H", "AAU" : "N", "GAU" : "D",
"UAC" : "Y", "CAC" : "H", "AAC" : "N", "GAC" : "D", 
"UGU" : "C", "CGU" : "R", "AGU" : "S", "GGU" : "G",
"UGC" : "C", "CGC" : "R", "AGC" : "S", "GGC" : "G", 
"UGG" : "W", "CGG" : "R", "AGG" : "R", "GGG" : "G",
"UAA" : "Stop", "CAA" : "Q", "AAA" : "K", "GAA" : "E",
"UAG" : "Stop", "CAG" : "Q", "AAG" : "K", "GAG" : "E",
"UGA" : "Stop", "CGA" : "R", "AGA" : "R", "GGA" : "G"
}
def complement(seq):
	return seq.replace('A', 'u').replace('U', 'a').replace('C', 'g').replace('G', 'c').upper()[::-1]
	
for idn, seq in read_fasta(arg.fasta):
	seq = seq.replace('T','U')
	peps = []
	
	for i in range(0,len(seq)-3):
		start = (codons[seq[i:i+3]] == 'M')
		if start:
			pep = ''
			for j in range(i,len(seq)-3,3):
				codon = seq[j:j+3]
				if codons[codon] != 'Stop': 
					pep += codons[codon]
				else: 
					if pep not in peps: peps.append(pep)
					break
	
	if arg.frame == 'double':
		comp_seq = complement(seq)
		for i in range(0,len(comp_seq)-3):
			start = (codons[comp_seq[i:i+3]] == 'M')
			if start:
				pep = ''
				for j in range(i,len(seq)-3,3):
					codon = comp_seq[j:j+3]
					if codons[codon] != 'Stop': 
						pep += codons[codon]
					else: 
						if pep not in peps: peps.append(pep)
						break
	
	print(f'>{idn}')
	if len(peps) > 0: print(max(peps, key=len))
	else: print('no orf found')
	
				
			
			

