#!/usr/bin/env python3
from readfasta import read_record
from itertools import product
import argparse
import random
import math
import json
import sys
import os


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

def locate_gene_start(gff):
	with open(gff) as fh:
		for line in fh:
			if 'mRNA' in line:
				fields = line.split()
				return int(fields[3])-1, int(fields[4])
	return -1, -1

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

def print_all_isoforms(paths, counts=None, total=None):
	for i, path in enumerate(paths):
		if counts:
			print(f'isoform {i+1} {counts[i]/total:.3f}:')
		else:
			print(f'isoform iteration {i+1}:')
		
		intron_start = None
		intron_end = None
		for j in range(len(path)):
			s = path[j]
			if s[0] == 'D0': intron_start = s[1]
			elif s[0] == 'A5':
				intron_end = s[2]
				print('Intron', '\t', intron_start, '\t', intron_end, sep='')
			elif s[0] == 'Exon' or s[0] == 'Genomic':
				print(s[0], '\t', s[1], '\t', s[2], sep='')
		
		print()

def print_prob_mat(probs):
	for i in range(len(probs[0])):
		for j in range(len(probs)):
			normalized = norm_logs(probs[j][i])
			ps = logs2probs(normalized)
			print(i, j, ps)
			
#	for i, path in enumerate(paths):
#		if counts:
#			print(f'isoform {i+1} {counts[i]/total:.2f}:')
#		else:
#			print(f'isoform iteration {i+1}:')
#		if path[0][0] == 'Genomic':
#			s = path[0]
#			print(s[0], '\t', 1, '\t', s[2], sep='')
#		else:
#			print('Genomic', '\t', 1, '\t', 3, sep='')
#			s = path[0]
#			print(s[0], '\t', s[1], '\t', s[2], sep='')
#		
#		# decoded
#		intron_start = None
#		intron_end   = None
#		for i in range(1, len(path)):
#			s = path[i]
#			if s[0] == 'D0': intron_start = s[1]
#			elif s[0] == 'A5':
#				intron_end = s[2]
#				print('Intron', '\t', intron_start, '\t', intron_end, sep='')
#			elif s[0] == 'Exon' or s[0] == 'Genomic':
#				print(s[0], '\t', s[1], '\t', s[2], sep='')
#		print()

def get_distribution(paths):
	isoforms = {}
	for path in paths:
		idn = []
		for s in path: idn.append(''.join(list(map(str, s))))
		idn = ''.join(idn)
		if idn not in isoforms:
			isoforms[idn] = [path, 1]
		else:
			isoforms[idn][1] += 1

	sorted_isoforms = dict(sorted(isoforms.items(), key=lambda x: x[1][1], reverse=True))

	return sorted_isoforms.values()

def print_distribution(dist, top=None):
	ps = []
	cs = []
	total = sum(elem[1] for elem in dist)
	for path, count in dist:
		ps.append(path)
		cs.append(count)
		if top:
			top -= 1
			if top <= 0: break
	print_all_isoforms(ps, cs, total)

def simiso(states, orders, trans, emiss, inits, terms, iterations, no_genomic):
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
	
	# Fill (Can beautify using matrix manipulation and numpy)
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

	
	paths = []
	# Trace back iteration
	for it in range(iterations):
		f_state = None
		fps = []
		for i in range(l):
			fp = sumlogps(probs[i][-1]) + terms[i]
			fps.append(fp)
		f_state = choose_randstate(fps)
		
		tb = [f_state]
		curj = len(seq)-n
		cur_probs = probs[f_state][curj]
		while curj > 0:
			curi = choose_randstate(cur_probs)
			tb.insert(0, curi)
			curj -= 1
			cur_probs = probs[curi][curj]
		
		tb.pop(0)
		
		# Store paths
		cur_state = tb[0]
		start, end = 0, 0
		path = []
		for i in range(1,len(tb)):
			if tb[i] != cur_state:
				end = i-1
				if no_genomic:
					path.append((states[cur_state], start+n+1+MRNA_START, end+n+1+MRNA_START))
				else:
					path.append((states[cur_state], start+n+1, end+n+1))
				start = i
				cur_state = tb[i]
		if no_genomic:
			path.append((states[cur_state],start+1+n+MRNA_START,len(tb)+n+MRNA_START))
		else:
			path.append((states[cur_state],start+1+n,len(tb)+n))
		
		paths.append(path)
	
	return paths
	
##########
## Main ##
##########

parser = argparse.ArgumentParser(
	description='Splicing simulator using stochastic viterbi')
parser.add_argument('fasta', type=str, metavar='<fasta>',
	help='path to FASTA file')
parser.add_argument('jhmm', type=str, metavar='<jhmm>',
	help='path to JHMM file')
parser.add_argument('--iterations', required=False, type=int, metavar='<int>',
	default=500, help='number of iterations [%(default)i]')
parser.add_argument('--top', required=False, type=int, metavar='<int>',
	default=10, help='out put top n isoforms [%(default)i]')
parser.add_argument('-ng', '--no_genomic', action='store_true',
	help='no genomic state (exon, intron, donor, acceptor states only)')
parser.add_argument('-gff', type=str, metavar='<gff>',
	help='path to GFF file, used to locate genomic-exon boundary (required when no_genomic tag is on)')
parser.add_argument('--all', action='store_true',
	help='print all isoforms')
arg = parser.parse_args()


for i, s in read_record(arg.fasta): seq = s

if arg.no_genomic:
	if not arg.gff: sys.exit('No gff input')
	MRNA_START, MRNA_END = locate_gene_start(arg.gff)
	if MRNA_START == -1: sys.exit('Can\'t locate mRNA')
	seq = seq[MRNA_START:MRNA_END]
with open(arg.jhmm) as fh: 
	states, orders, trans, emiss, inits, terms = read_jhmm(fh)

paths = simiso(states, orders, trans, emiss, inits, terms, arg.iterations, arg.no_genomic)	

dist = get_distribution(paths)
if arg.all:
	print(f"cmd: {' '.join(sys.argv)}\n")
	print_distribution(dist)
else:
	print(f"cmd: {' '.join(sys.argv)}\n")
	print_distribution(dist, top=arg.top)
		
		
