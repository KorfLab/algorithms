import argparse
import json
from seqlib import read_fasta
from itertools import product

parser = argparse.ArgumentParser(description='Determine the k-mer freqeuncies in a FASTA file')
parser.add_argument('fasta', type=str, metavar='<file>', help='input fasta file(s)')
parser.add_argument('--k', required=False, type=int, default=3,
	metavar='<int>', help='input kmer size [%(default)i]')
arg = parser.parse_args()

kmers = {}
for kmer in product(['A','C','G','T'], repeat=arg.k):
	kmers[''.join(kmer)]=0
	
for iden, seq in read_fasta(arg.fasta):
	for i in range(len(seq)-arg.k+1):
		kmer = seq[i:i+arg.k]
		kmers[kmer]+=1
		
tot=sum(kmers.values())
for key in kmers:
	kmers[key]/=tot



# Write into tsv
with open('kmer.tsv', 'w') as fh:
	for key in kmers:
		fh.write(f'{key}\t{kmers[key]}\n')

# Write into json
with open('kmer.json', 'w') as fh:
	injson = json.dumps(kmers, indent = 4)
	fh.write(injson)
		

