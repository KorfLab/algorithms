import random

def kmers(filename, k):
	kmers = []
	with open(filename) as fp:
		for line in fp.readlines():
			if line.startswith('>'): continue
			for i in range(len(line) -k):
				kmers.append(line[i:i+k])
	return kmers
				

pos = kmers('235.fa', 3)
neg = kmers('300.fa', 3)

pnt = kmers('235.fa', 1)
nnt = kmers('300.fa', 1)

scheme = (
	[None, 'ACGT', 100],
	['ip10', pos, 10],
	[None, 'ACGT', 100],
	['in10', neg, 10],
	[None, 'ACGT', 100],
	['ip20', pos, 20],
	[None, 'ACGT', 100],
	['in20', neg, 20],
	[None, 'ACGT', 100],
	['ip30', pos, 30],
	[None, 'ACGT', 100],
	['in30', neg, 30],
	[None, 'ACGT', 100],
	['ip40', pos, 40],
	[None, 'ACGT', 100],
	['in40', neg, 40],
	[None, 'ACGT', 100],
	['ip50', pos, 50],
	[None, 'ACGT', 100],
	['in50', neg, 50],
	
	[None, pnt, 100],
	['pp10', pos, 10],
	[None, pnt, 100],
	['pn10', neg, 10],
	[None, pnt, 100],
	['pp20', pos, 20],
	[None, pnt, 100],
	['pn20', neg, 20],
	[None, pnt, 100],
	['pp30', pos, 30],
	[None, pnt, 100],
	['pn30', neg, 30],
	[None, pnt, 100],
	['pp40', pos, 40],
	[None, pnt, 100],
	['pn40', neg, 40],
	[None, pnt, 100],
	['pp50', pos, 50],
	[None, pnt, 100],
	['pn50', neg, 50],

	[None, nnt, 100],
	['np10', pos, 10],
	[None, nnt, 100],
	['nn10', neg, 10],
	[None, nnt, 100],
	['np20', pos, 20],
	[None, nnt, 100],
	['nn20', neg, 20],
	[None, nnt, 100],
	['np30', pos, 30],
	[None, nnt, 100],
	['nn30', neg, 30],
	[None, nnt, 100],
	['np40', pos, 40],
	[None, nnt, 100],
	['nn40', neg, 40],
	[None, nnt, 100],
	['np50', pos, 50],
	[None, nnt, 100],
	['nn50', neg, 50],
)

fa = open('testseq.fa', 'w')
bed = open('testseq.bed', 'w')

total = 0
fa.write('>testseq\n')
for i in range(10):
	for name, dist, length in scheme:
		seq = ''
		for j in range(length):
			seq += random.choice(dist)
		if name is not None:
			bed.write(f'{name}\t{total +1}\t{total + len(seq)}\n')
		fa.write(f'{seq}\n')
		total += len(seq)
