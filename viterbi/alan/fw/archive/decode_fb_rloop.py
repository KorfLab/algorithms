#!/usr/bin/env python3

from readfasta import read_record
from itertools import product
import numpy as np
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


###############
## Functions ##
###############
def sumlogp(a, b, mag=40):
	if a == float('-inf') and b == float('-inf'): return float('-inf')
	assert(a <= 0)
	assert(b <= 0)
	if abs(a - b) > mag: return max(a, b)
	if a < b: return math.log(1 + math.exp(a - b)) + b
	return math.log(1 + math.exp(b - a)) + a

def sumlogp_np(a ,b):
	return np.logaddexp(a, b)

def log(prob):
	if prob == 0: return float('-inf')
	else: return math.log(prob)

def draw_mat(mat):
	for row in mat:
		for elem in row: print(f'{elem:.2f}', end='  ')
		print('\n')

def emission(num, probs):
	emiss = {}
	order = int(math.log(num, 4) - 1)
	for i, kmer in enumerate(list(product('ACGT', repeat=order))):
		kmer = ''.join(kmer)
		emiss[kmer] = {}
		for j, letter in enumerate('ACGT'):
			emiss[kmer][letter] = log(probs[4*i+j])
	return order, emiss
	
def decode(seq, states, transition, emission, orders, inits, terms):
	n = max(orders)
	
	# Initializae matrices
	fw = [[0]*(len(seq)-n) for _ in range(len(transition))]
	bw = [[0]*(len(seq)-n) for _ in range(len(transition))]
	# Forward
	# Initialize
	for i, init in enumerate(inits):
		kmer = seq[0:n]
		base = seq[n]
		fw[i][0] = init + emission[i][kmer][base]
	# Fill
	for i in range(1, len(seq)-n):
		for j in range(len(states)):
			total = log(0)
			cur_order = orders[j]
			kmer = seq[i+n-cur_order:i+n]
			base = seq[i+n]
			for s in range(len(states)):
				score = fw[s][i-1] + transition[s][j] + emission[j][kmer][base]
				total = sumlogp(total, score)
			fw[j][i] = total
			
	# Terminal probability
	for i in range(len(states)): fw[i][-1] += terms[i]
	draw_mat(fw)
	
	return -1	
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

print(emiss)

for idn, seq in read_record(arg.input):
	decoded = decode(seq, states, trans, emiss, orders, inits, terms)

