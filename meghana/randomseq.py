"""
Generate random DNA sequences of fixed length.
The composition of the sequences defaults to 25% for each nucleotide, 
but the program should take other distributions of mononucleotides.

Inputs

Number of sequences to generate
Length of each sequence
Probability of each letter
Random seed

Outputs
Multi-FASTA format to STDOUT
"""

import argparse
import random
import time

parser = argparse.ArgumentParser(description='Generate random DNA sequences of \
fixed length.')
parser.add_argument('-s', type = int, required = False, default = 10, 
metavar = 'seqs', help = 'Number of sequences to generate.')
parser.add_argument('-l', type = int, required = False, default = 80,
metavar = 'length', help = 'Length of each sequence')
parser.add_argument('-p', type = float, nargs='+', required = False, 
default = [.25, .25, .25, .25], help = 'Probability of each nt [A, T, C, G]')
parser.add_argument('-seed', type = float, required = False, default = time.time(),
help = 'Random seed')
args = parser.parse_args()

random.seed(args.seed)

for i in range(args.s):
	print(f'>ID {i+1}')
	seq = ''.join(random.choices(population = ['A', 'T', 'C', 'G'], weights=args.p,
	k=args.l))
	print(seq)
	
	
	