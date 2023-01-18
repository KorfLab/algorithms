import readfasta
import argparse
import math
from itertools import product

parser = argparse.ArgumentParser(description='Determine the location of r-loops\
using hmms')
parser.add_argument('pos', type=str, help='fasta file for positive strand r-loops')
parser.add_argument('neg', type=str, help='fasta file for negative strand r-loops')
parser.add_argument('test', type=str, help='fasta file to decode')
parser.add_argument('-n', type=int, default=1, help='n of hmm')
arg = parser.parse_args()

def pscore(num):
	if num == 0: return 99
	return -math.log2(num)
	
def kmers_dict(n):
	if n == 1:
		freqs = {'A': pscore(0.25),
				 'C': pscore(0.25),
				 'G': pscore(0.25),
				 'T': pscore(0.25)}
		return freqs
	else:
		kmers = product('ACGT', repeat=(n-1))
		freqs = {}
		for kmer in kmers:
			kmer = ''.join(kmer)
			freqs[''.join(kmer)] = kmers_dict(1)
		return freqs
	
def freqs_dict(fasta, n):
	freqs = kmers_dict(arg.n)
	for id, seq in readfasta.read_record(fasta):
		for i in range(len(seq)-n+1):
			if n == 1:
				freqs[seq[i]] += 1
			else:
				freqs[seq[i:i+n-1]][seq[i+n-1]] += 1
			
	if n==1:
		total = sum(freqs.values())
		for k in freqs:
			freqs[k] = pscore(freqs[k]/total)
		return freqs
	
	for i in freqs.keys():
		total = sum(freqs[i].values())
		for j in freqs[i].keys():
			freqs[i][j] = pscore(freqs[i][j]/total)
	return freqs

def viterbi(emission, initial, transition, n, file):
	for id, seq in readfasta.read_record(file):
		#initialization
		scores = [[0]*(len(seq)) for _ in range(len(initial))]
		states = [[-1]*(len(seq)) for _ in range(len(initial))]
		for i in range(len(initial)):
			if n == 1:
				scores[i][0] = initial[i]+emission[i][seq[0]]
				states[i][0] = -1
			else:
				scores[i][n-1] = initial[i]+emission[i][seq[0:n-1]][seq[n-1]]
				states[i][n-1] = 0
		
		#fill
		for i in range(0, len(initial)):
			for j in range(n, len(seq)):
				for k in range(len(initial)):
					maxx = None
					argmax = None
					if n==1 :
						for l in range(len(initial)):
							score = scores[i][j-1]+transition[i][k]+emission[i][seq[j]]
							if maxx is None or score < maxx:
								maxx = score
								argmax = l
						scores[i][j] = maxx
						states[i][j] = argmax
					else:
						for l in range(len(initial)):
							score = scores[i][j-1]+transition[i][k]+emission[i][seq[j-n+1:j]][seq[j]]
							if maxx is None or score < maxx:
								maxx = score
								argmax = l
						scores[i][j] = maxx
						states[i][j] = argmax
					

		#print(scores[0][0:100], scores[1][0:100], scores[2][0:100])
		#traceback
		maxx = scores[0][len(seq)-1]
		pointer = 0
		
		#print(scores[0][len(seq)-1], scores[1][len(seq)-1], scores[2][len(seq)-1])
		for i in range(1, len(initial)):
			if scores[i][len(seq)-1] > maxx:
				maxx = initial[i]
				pointer = i
				
		#print(pointer)
		bestpath = [-1]*len(seq)
		bestpath[len(seq)-1] = pointer	
		for i in range(len(seq)-2, n-1, -1):
			bestpath[i] = states[bestpath[i+1]][i+1]
		
		#print(bestpath)


#emission 
genome = kmers_dict(arg.n)
positive = freqs_dict(arg.pos, arg.n)
negative = freqs_dict(arg.neg, arg.n)
emission = [genome, positive, negative]

#initial 
initial = [pscore(.33), pscore(.33), pscore(.33)] #genome, positive, negative

#transition
transition = [[pscore(99/100), pscore(.5/100), pscore(.5/100)], #gnm->gnm, gnm->pos, gnm->neg
			  [pscore(1/90),   pscore(89/90),  pscore(0)],     			 #r+->gnm,  r+->r+,   r+->r-
			  [pscore(1/90),   pscore(0),                 pscore(89/90)]]  #r-->gnm,  r-->r+    r-->r-
			  
n = arg.n
file = arg.test

viterbi(emission, initial, transition, n, file)			  
			  
			  