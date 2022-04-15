"""
Mask sequences with an entropy filter. 
The window size and entropy should have default parameters and command line 
options. 
There should be an option to change the output from N-based (hard) masking to 
lowercase (soft) masking.

Inputs

Multi-FASTA file (gzipped or STDIN)
Window size
Entropy threshold
N-based or lowercase masking

Outputs

Multi-FASTA file to STDOUT
"""

import argparse
import math

parser = argparse.ArgumentParser(description='Mask sequences with an entropy \
filter')
parser.add_argument('-f', type = str, required = True, help = 'path to file')
parser.add_argument('-w', type = int, required = False, default = 11,
help = 'window size')
parser.add_argument('-e', type = float, required = False, default = 1.1,
help = 'enthropy threshold')
parser.add_argument('-n', type = str, required = False, default = 'True',
help = 'N based masking = True, lowercase = False')
args = parser.parse_args()


def entropy(freqs):
	a = freqs[0]
	t = freqs[1]
	g = freqs[2]
	c = freqs[3]
	total = sum(freqs)
	if total == 0: return 0
	
	pa = a/total
	pt = t/total
	pg = g/total
	pc = c/total
	
	h = 0
	if (a != 0): h -= pa * math.log(pa)
	if (t != 0): h -= pt * math.log(pt)
	if (g != 0): h -= pg * math.log(pg)
	if (c != 0): h -= pc * math.log(pc)
	
	return h/math.log(2)
	
ff = open(args.f, 'r')
for line in ff:
	line = line.strip()
	if line[0] == '>':
		print(line)
	else:
		newline = line[0:args.w//2]
		for i in range(len(line)-args.w+1):
			string = line[i:i+args.w]
			counts = {'A': 0,
					  'T': 0,
					  'G': 0,
					  'C': 0}
			for nt in string:
				counts[nt] += 1
			count_list = [counts[value] for value in counts]
			if entropy(count_list) < args.e:
				if args.n == 'True':
					newline += 'N'
				elif args.n == 'False':
					newline += line[i+args.w//2 + 1].lower()
			else:
				newline += line[i+args.w//2+1]
		newline += line[-args.w//2 + 1:]	
		print(newline)
		

	
