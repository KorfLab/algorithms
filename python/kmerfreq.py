"""
Determine the k-mer frequencies in a FASTA file. 
The value of K should be an argument with a default parameter (e.g. 3). 
Output format should include tab-separated and JSON.

Inputs

Multi-FASTA file (gzipped or STDIN)
K-mer size

Outputs

TSV
JSON
"""

import argparse
import csv
import json
from readfasta import read_record

parser = argparse.ArgumentParser(description='Determine the kmer frequences \
in a fasta file')
parser.add_argument('fasta', type = str, help = 'path to file')
parser.add_argument('-K', type = int, required = False, help = 'K-mer size', default = 3)
parser.add_argument('-j', action = "store_true", help = " ")
arg = parser.parse_args()

total = 0
freq = {}
for id, seq in read_record(arg.fasta):
	for i in range(len(seq) - arg.K + 1):
		kmer = seq[i:i+arg.K]
		if kmer not in freq: freq[kmer] = 0
		freq[kmer] += 1
		total += 1
		
for kmer in freq: freq[kmer] /= total
if arg.j:
	print(json.dumps(freq, indent=4))
else:
	for key in freq:
		print(f'{key}	{freq[key]}')
