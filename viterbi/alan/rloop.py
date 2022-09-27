#!/usr/bin/env python3

from readfasta import read_record
from itertools import product
import argparse
import math
import sys

parser = argparse.ArgumentParser(
	description='Viterbi algorithm for rloops')
parser.add_argument('positive', type=str, metavar='<fasta>',
	help='input fasta file for training data, positive strand')
parser.add_argument('negative', type=str, metavar='<fasta>',
	help='input fasta file for training data, negative strand')
parser.add_argument('-n', type=int, metavar='<int>', default=2,
	help='nth order markov model for emission probabilities [%(default)i]')
arg = parser.parse_args()

###############
## Functions ##
###############

def make_model(n):
	kmers = {}
	if n == 0:
		for kmer in product(list('ACGT'), repeat = 1):
			kmer = ''.join(kmer)
			kmers[kmer] = 0
	elif n > 0:
		for kmer in product(list('ACGT'), repeat = n):
			kmer = ''.join(kmer)
			kmers[kmer] = {}
			for letter in list('ACGT'):
				kmers[kmer][letter] = 0
	else:
		sys.exit('Cannot have negative order')
	return kmers

def prob2score(p):
	if p == 0: return -99
	return math.log2(p/0.25)
	
def train(file, n):
	model = make_model(n)
	for idn, seq in read_record(file):
		if n == 0:
			for bp in seq: model[bp] += 1
		else:
			for i in range(len(seq)-n-1):
				given = seq[i:i+n]
				curbp = seq[i+n+1]
				model[given][curbp] += 1
	if n == 0:
		total = sum(model.values())
		for letter in model:
			p = model[letter] / total
			model[letter] = prob2score(p)
	else:
		for kmer in model:
			total = sum(model[kmer].values())
			for letter in model[kmer]:
				p = model[kmer][letter] / total
				model[kmer][letter] = prob2score(p)
	return model

def seq2int(seq):
	seq = list(seq)
	for i in range(len(seq)):
		if seq[i] == 'A' or seq[i] == 'a' : seq[i] = 0
		if seq[i] == 'C' or seq[i] == 'c' : seq[i] = 1
		if seq[i] == 'G' or seq[i] == 'g' : seq[i] = 2
		if seq[i] == 'T' or seq[i] == 't' : seq[i] = 3
		if seq[i] == 'N' or seq[i] == 'n' : seq[i] = random.randint(0,3)
	return seq
	
##########################
## Emission Probability ##
##########################

# dictionary with probability transformed log score
positive = train(arg.positive, arg.n)
negative = train(arg.negative, arg.n)

############################
## Transition Probability ##
############################

#           genomic    positive    negative
# genomic    g->g        g->p        g->n        
# positive   p->g       (1/300)        0 
# negative   n->g          0        (1/300)
transition = [[0 for _ in range(3)] for _ in range(3)]
transition[0][0], transition[0][1], transition[0][2] = 0, 0, 0
transition[1][0], transition[1][1], transition[1][2] = 1/300, 299/300, 0
transition[2][0], transition[2][1], transition[2][2] = 1/300, 0, 299/300
for r in range(len(transition)):
	for c in range(len(transition[0])):
		transition[r][c] = prob2score(transition[r][c])

##################
## Main Viterbi ##
##################
'''
for idn, seq in read_record(arg.genome):
	seq = seq2int(seq)
	
	# Initialize matrices
	prob = [[0]*(len(seq-arg.n+1)) for _ in range(len(state))]
	tran = [[-1]*(len(seq-arg.n+1)) for _ in range(len(state))]
	
	# Initialize state probabilities
'''
