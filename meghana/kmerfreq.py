#Determine the k-mer frequencies in a FASTA file. 
#The value of K should be an argument with a default parameter (e.g. 3). 
#Output format should include tab-separated and JSON.

#Inputs

#Multi-FASTA file (gzipped or STDIN)
#K-mer size

#Outputs

#TSV
#JSON

import argparse
import csv
import json

parser = argparse.ArgumentParser(description='Determine the kmer frequences \
in a fasta file')
parser.add_argument('-f', type = str, required = True, help = 'path to file')
parser.add_argument('-K', type = int, required = False, help = 'K-mer size', default = 3)
args = parser.parse_args()

ff = open(args.f, 'r')
freq = {}
for line in ff:
	if line[0] != '>':
		line = line.strip()
		for i in range(len(line)-args.K+1):
			kmer = line[i:i+args.K]
			if kmer in freq: freq[kmer] += 1
			if kmer not in freq: freq[kmer] = 1

ff.close()
total = sum(freq.values(), 0.0)
freq = {k: v / total for k, v in freq.items()}

with open('kmerfreq.tsv', 'wt') as out_file:
     tsv_writer = csv.writer(out_file, delimiter='\t')
     for key in freq:
     	tsv_writer.writerow([key, freq[key]])

with open('json_data.json', 'w') as outfile:
    json.dump(freq, outfile) # is this how its supposed to look?
