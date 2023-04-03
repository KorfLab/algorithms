#!/usr/bin/env python3

from readfasta import read_record
from itertools import product
import argparse
import json
import math
import sys

parser = argparse.ArgumentParser(
	description='Viterbi algorithm for rloops')
parser.add_argument('input', type=str, metavar='<fasta>',
	help='input fasta file to decode')
parser.add_argument('hmm', type=str, metavar='<json>',
	help='input json file for trained hmm models')
arg = parser.parse_args()

#########################################
## Decode Function and other functions ##
#########################################
def decode(seq, states, transition, emission, orders, inits, terms):
	# Determine order
	n = max(orders)

	# Initialize matrices
	prob = [[0]*(len(seq)-n+1) for _ in range(len(transition))]
	tran = [[-1]*(len(seq)-n+1) for _ in range(len(transition))]

	# Initialize state probabilities
	for i, init in enumerate(inits): prob[i][0] = init

	# Fill
	for i in range(len(seq)-n):
		for j in range(len(states)):
			max_score = None
			max_state = None
			cur_order = orders[j]
			kmer = seq[i+n-cur_order:i+n]
			base = seq[i+n]
			for s in range(len(states)):
				score = prob[s][i] + transition[s][j] + emission[j][kmer][base]
				if max_score is None or score > max_score:
					max_score = score
					max_state = s	
			prob[j][i+1] = max_score
			tran[j][i+1] = max_state
	
	# Trace back
	max_score = None
	max_state = None
	for i in range(len(transition)):
		score = prob[i][len(seq)-n] + terms[i]
		if max_score is None or score > max_score:
			max_score = score
			max_state = i
	
	tb = [max_state]
	curj = len(seq) - n
	curi = tran[max_state][curj]
	while curj > 0:
		tb.insert(0, curi)
		curj -= 1
		curi = tran[curi][curj]
	tb.pop(0)
	
	# Print result
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
	
	return decoded

def emission(num, probs):
	emiss = {}
	order = int(math.log(num, 4) - 1)
	for i, kmer in enumerate(list(product('ACGT', repeat=order))):
		kmer = ''.join(kmer)
		emiss[kmer] = {}
		for j, letter in enumerate('ACGT'):
			emiss[kmer][letter] = log(probs[4*i+j])
	return order, emiss

def log(prob):
	if prob == 0: return float('-inf')
	else: return math.log(prob)

def draw_mat(mat):
	for row in mat:
		for elem in row: print(f'{elem:.2f}', end='  ')
		print('\n')

######################
## Harvest HMM Data ##
######################
with open(arg.hmm) as fh: hmm = json.load(fh)

states = []
orders = []
trans = []
emiss = []
inits = []
terms = []


for state in hmm['state']:
	order, emis = emission(state['emissions'], state['emission'])
	orders.append(order)
	emiss.append(emis)
	states.append(state['name'])
	inits.append(log(state['init']))
	terms.append(log(state['term']))
	
for f in hmm['state']:
	cur_trans = []
	for t in states:
		if t in f['transition']: cur_trans.append(log(f['transition'][t]))
		else: cur_trans.append(float('-inf'))
	trans.append(cur_trans)

##########
## Main ##
##########
for idn, seq in read_record(arg.input):
	decoded = decode(seq, states, trans, emiss, orders, inits, terms)
	for d in decoded:
		print(d[0],'\t',d[1],'\t',d[2],sep='')
