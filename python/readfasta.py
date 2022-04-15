import gzip
import random
import math

def read_record(filename):
	
	label = None
	seq = []
	
	fp = None
	if filename.endswith('.gz'): 
		fp = gzip.open(filename, 'rt')
	else: 
		fp = open(filename)

	for line in fp.readlines(): #change to readline() for memory
		line = line.rstrip()
		if line.startswith('>'):
			if len(seq) > 0:
				seq = ''.join(seq)
				yield(label, seq)
				label = line[1:]
				seq = []
			else:
				label = line[1:]
		else:	
			seq.append(line)
	yield(label, ''.join(seq))
	fp.close()
	