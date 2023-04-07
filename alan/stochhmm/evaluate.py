'''
	This program is used to evaluate the output from simiso using gff file as reference.
	Currently only support output generated assuming no genomic sequence.
'''

import sys

def identical(ref, sim):
	if len(ref) != len(sim): return False
	for fref, fsim in zip(ref, sim):
		if fref[1:] != fsim[1:]: return False
	return True

sim_fp = sys.argv[1]
ref_fp = sys.argv[2]

ref = []
mrna_start = None
with open(ref_fp) as fh:
	for line in fh:
		if 'WormBase' not in line: continue
		if 'mRNA' in line: mrna_start = line.split()[3]
		if 'intron' in line or 'exon' in line:
			f = line.split()
			feature = [f[2], f[3], f[4], f[8]]
			ref.append(feature)

parent = ref[0][-1]
for i, feature in enumerate(ref):
	assert(feature[-1] == parent)
	ref[i].pop(-1)

ref = sorted(ref, key=lambda x: x[-1])

#for feature in ref: print('\t'.join(feature))

cur_probability = None
isoform = []
with open(sim_fp) as fh:
	isoforms = fh.readlines()[2:]
	isoforms = list(map(lambda x: x.strip(), isoforms))
	isoforms = list(filter(lambda x: x!= '', isoforms))
	for line in isoforms:
		if line.startswith('isoform'):
			if len(isoform) > 0:
				isoform[0][1] = mrna_start
				if identical(ref, isoform):
					print(cur_probability)
					for feature in isoform: print('\t'.join(feature))
				isoform = []
			cur_probability = float(line.split()[2][:-1])
		else:
			f = line.split()
			feature = line.split()
			isoform.append(feature)
	
	isoform[0][1] = mrna_start
	if identical(ref, isoform):
		print(cur_probability)
		for feature in isoforms: print('\t'.join(feature))
		
		
