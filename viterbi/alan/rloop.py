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
parser.add_argument('test', type=str, metavar='<fasta>',
	help='input fasta file for testing')	
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

	
##########################
## Emission Probability ##
##########################

# dictionary with probability transformed log score
n = arg.n
genomic  = make_model(n)
positive = train(arg.positive, n)
negative = train(arg.negative, n)
emission = [genomic, positive, negative]
############################
## Transition Probability ##
############################

#           genomic    positive    negative
# genomic    g->g        g->p        g->n        
# positive   p->g       (1/300)        0 
# negative   n->g          0        (1/300)
transition = [[0 for _ in range(3)] for _ in range(3)]
transition[0][0], transition[0][1], transition[0][2] = 99/100, 0.5/100, 0.5/100
transition[1][0], transition[1][1], transition[1][2] = 1/90, 89/90, 0
transition[2][0], transition[2][1], transition[2][2] = 1/90, 0, 89/90
for r in range(len(transition)):
	for c in range(len(transition[0])):
		transition[r][c] = prob2score(transition[r][c])
		
##################
## Main Viterbi ##
##################

for idn, seq in read_record(arg.test):
	# Initialize matrices
	prob = [[0]*(len(seq)-n+1) for _ in range(len(transition))]
	tran = [[-1]*(len(seq)-n+1) for _ in range(len(transition))]

	# Initialize state probabilities
	prob[0][0], prob[1][0], prob[2][0] = prob2score(100/280), prob2score(90/280), prob2score(90/280)

	
	# Fill
	for i in range(len(seq)-n):
		if n == 0:
			letter = seq[i]
			for j in range(len(transition)):
				max_score = None
				max_state = None
				for s in range(len(transition)):
					score = prob[s][i-1] + transition[s][j] + emission[j][letter]
					if max_score is None or score > max_score:
						max_score = score
						max_state = s
				prob[j][i+1] = max_score
				tran[j][i+1] = max_state
		else:
			kmer = seq[i:i+n]
			letter = seq[i+n]
			for j in range(len(transition)):
				max_score = None
				max_state = None
				for s in range(len(transition)):
					score = prob[s][i-1] + transition[s][j] + emission[j][kmer][letter]
					if max_score is None or score > max_score:
						max_score = score
						max_state = s	
				prob[j][i+1] = max_score
				tran[j][i+1] = max_state
	
	# Trace back
	max_score = None
	max_state = None
	for i in range(len(transition)):
		score = prob[i][len(seq)-n]
		if max_score is None or score > max_score:
			max_score = score
			max_state = i
	
	tb = []
	curi = max_state
	curj = len(seq) - n
	while tran[curi][curj] != -1:
		tb.insert(0, tran[curi][curj])
		curj -= 1
		curi = tran[curi][curj]
	
	# Print result
	states = ['gnm', 'pos', 'neg']
	cur_state = tb[0]
	start, end = 0, 0
	decoded = []
	for i in range(1,len(tb)):
		if tb[i] != cur_state:
			end = i-1
			decoded.append((states[cur_state], start+n+1, end+n+1))
			start = i
			cur_state = tb[i]
	decoded.append((states[cur_state],start+1+n,len(tb)+n))
	
	for d in decoded:
		print(d[0],'\t',d[1],'\t',d[2],sep='')
			
			
			
