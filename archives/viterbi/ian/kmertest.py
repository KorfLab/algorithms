import random
import itertools
import math

def readfile(file, k):
	n = {}
	total = 0
	with open(file) as fp:
		for line in fp.readlines():
			if line.startswith('>'): continue
			for i in range(len(line) -k):
				kmer = line[i:i+k]
				if kmer not in n: n[kmer] = 0
				n[kmer] += 1
				total += 1
	
	f = {}
	for kmer in n:
		f[kmer] = math.log2(n[kmer]/total)
	return f

def score(seq, model, k):
	p = 0
	for i in range(len(seq) -k + 1):
		kmer = seq[i:i+k]
		p += model[kmer]
	return p



k = 2
rpos = readfile('../235.fa', k)
rneg = readfile('../300.fa', k)
rand = {}
for thing in itertools.product('ACGT', repeat=k):
	kmer = ''.join(thing)
	rand[kmer] = math.log2(0.25 ** k)

# part 1: random seq
rwin = 0
nseqs = 100
rlen = 50
for i in range(nseqs):
	seq = ''
	for j in range(rlen):
		seq += random.choice('ACGT')
	spos = score(seq, rpos, k)
	sneg = score(seq, rneg, k)
	sran = score(seq, rand, k)
	if sran > spos and sran > sneg: rwin += 1
print(rwin/nseqs)

# part 2: pos strand
with open('../235.fa') as fp:
	for line in fp.readlines():
		if line.startswith('>'): continue
		seq = line.rstrip()
		spos = score(seq, rpos, k)
		sneg = score(seq, rneg, k)
		sran = score(seq, rand, k)
		print(seq)
		print(spos)
		print(sneg)
		print(sran)