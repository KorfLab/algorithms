#!/usr/bin/env python3
from readfasta import read_record
from itertools import product
import random
import math
import json
import sys


###############
## Functions ##
###############

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

def sumlogp(a, b, mag=40):
	if a == float('-inf') and b == float('-inf'): return float('-inf')
	assert(a <= 0)
	assert(b <= 0)
	if abs(a - b) > mag: return max(a, b)
	if a < b: return math.log(1 + math.exp(a - b)) + b
	return math.log(1 + math.exp(b - a)) + a

def sumlogps(logps):
	prob_sum = logps[0]
	for i in range(1, len(logps)): prob_sum = sumlogp(prob_sum, logps[i])
	return prob_sum

def norm_logs(logps):
	normed = []
	prob_sum = sumlogps(logps)
	for i in range(len(logps)): normed.append(logps[i] - prob_sum)
	return normed

def logs2probs(logps):
	return list(map(math.exp, logps))
	
def choose_randstate(logps):
	normalized = norm_logs(logps)
	probs = logs2probs(normalized)
	choice = random.choices(range(len(logps)), weights=probs)[0]
	return choice

def read_jhmm(fh):
	hmm = json.load(fh)
	
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
	
	return states, orders, trans, emiss, inits, terms

def simiso(states, orders, trans, emiss, inits, terms, iterations):
	# n is the highest order in the states and used for initializing probability matrix
	n = max(orders)
	l = len(states)
	
	# Initialize probability matrix
	probs = [[[0] * l for i in range(len(seq)-n+1)] for j in range(l)]
	
	# Initialize state probabilities
	for i, init in enumerate(inits):
		cur = []
		for j in range(l): cur.append(init - log(l))
		probs[i][0] = cur
	
	# Fill
	for i in range(n, len(seq)):
		base = seq[i]
		for c in range(l):
			cur_probs = []
			cur_order = orders[c]
			kmer = seq[i-cur_order:i]
			for p in range(l):
				prob = sumlogps(probs[p][i-n]) + trans[p][c] + emiss[c][kmer][base]
				cur_probs.append(prob)
			probs[c][i-n+1] = cur_probs

	# Trace Back
	f_state = None
	fps = []
	for i in range(l):
		fp = sumlogps(probs[i][-1]) + terms[i]
		fps.append(fp)
	f_state = choose_randstate(fps)
	
	print(f_state)
	
	tb = [f_state]
	
	
##########
## Main ##
##########
seq_fp  = sys.argv[1]
jhmm_fp = sys.argv[2]
iterations = int(sys.argv[3])

random.seed(1)

for i, s in read_record(seq_fp): seq = s

with open(jhmm_fp) as fh: 
	states, orders, trans, emiss, inits, terms = read_jhmm(fh)

isoforms = simiso(states, orders, trans, emiss, inits, terms, iterations)

