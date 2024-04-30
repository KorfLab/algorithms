import argparse
import json
import math
import readfasta
from itertools import product

parser = argparse.ArgumentParser(description='Create file with data for hmm in .json')
parser.add_argument('pos', type=str, help='fasta file for positive strand r-loops')
parser.add_argument('neg', type=str, help='fasta file for negative strand r-loops')
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

dictionary = {
	"name": "demo",
	"author": "Meghana",
	"states": ["s1", "s2", "s3"],
	"inits": {
		"s1": pscore(.33),
		"s2": pscore(.33),
		"s3": pscore(.33)
	},
	"transitions": {
		"s1": {
			"s1": pscore(99/100),
			"s2": pscore(.5/100),
			"s3": pscore(.5/100)
		},
		"s2": {
			"s1": pscore(1/90),
			"s2": pscore(89/90),
			"s3": pscore(0/90)
		},
		"s3": {
			"s1": pscore(1/90),
			"s2": pscore(0),
			"s3": pscore(89/90)
		}
	},
	"emissions": {
		"s1": kmers_dict(arg.n),
		"s2": freqs_dict(arg.pos, arg.n),
		"s3": freqs_dict(arg.neg, arg.n)
	}
}

filename = str(arg.n) + ".json"
json_object = json.dumps(dictionary, indent = 4)
with open(filename, "w") as outfile:
	outfile.write(json_object)
