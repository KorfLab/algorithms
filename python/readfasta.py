import gzip
import random
import math

def read_fasta(filename):
	
	label = None
	seq = []
	
	fp = None
	if filename.endswith('.gz'): 
		fp = gzip.open(filename, 'rt')
	else: 
		fp = open(filename)

	for line in fp.readlines():
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
	
def read_fastq(filename):
	
	if filename.endswith('.gz'): 
		fp = gzip.open(filename, 'rt')
	else: 
		fp = open(filename)

	while True:
		name = fp.readline()
		seq = fp.readline()
		plus = fp.readline()
		qual = fp.readline()
		if name == '': break
		yield(name, seq, qual)
		
def random_dna(length, A, C, G, T):
	
	assert(math.isclose(1, A+C+G+T, abs_tol = 0.001))
	seq = []
	for i in range(length):
		r = random.random()
		if r < A:		seq.append('A')
		elif r < A + C:	seq.append('C')
		elif r < A + C +G:	seq.append('G')
		else:			seq.append('T')
	return ''.join(seq)
		
		

		
		
		
		
		
		
		
		
		
		
		
		
		
		

