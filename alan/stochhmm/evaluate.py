'''
	This program is used to evaluate the output from simiso using gff file as reference.
	Currently only support output generated assuming no genomic sequence.
'''

import sys
import os
import re

def identical(ref, sim):
	if len(ref) != len(sim): return False
	for fref, fsim in zip(ref, sim):
		if fref[1:] != fsim[1:]: return False
	return True

def search(sim_fp, ref_fp):
	ref = []
	mrna_start = None
	parent = None
	with open(ref_fp) as fh:
		for line in fh:
			if 'WormBase' not in line: continue
			if 'mRNA' in line:
				mrna_start = line.split()[3]
				parent = re.search(r'Transcript:[^;]+', line.split()[8]).group(0)
			if 'intron' in line or 'exon' in line:
				f = line.split()
				feature = [f[2], f[3], f[4], f[8]]
				ref.append(feature)

	tmp = []
	for feature in ref:
		if parent not in feature[-1]: continue
		tmp.append(feature[:-1])
	
	ref = tmp
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
						return isoform, cur_probability
					isoform = []
				cur_probability = float(line.split()[2][:-1])
			else:
				f = line.split()
				feature = line.split()
				isoform.append(feature)
		
		isoform[0][1] = mrna_start
		if identical(ref, isoform): return isoform, cur_probability
	
	return None, None

DATA_DIR = 'data/apc'
SIM_OUT_DIR = 'out'
for f in os.listdir(DATA_DIR):
	if f.endswith('.fa'): continue
	gene_name = os.path.splitext(f)[0]
	sim_fp = os.path.join(SIM_OUT_DIR, f'out_{gene_name}')
	ref_fp = os.path.join(DATA_DIR, f'{gene_name}.gff3')
	isoform, prob = search(sim_fp, ref_fp)
	
	if isoform is not None:
		print(f'{gene_name}\t{prob}')
	else:
		print(f'{gene_name}\t0.000')

#sim_fp = sys.argv[1]
#ref_fp = sys.argv[2]

#isoform, probability = search(sim_fp, ref_fp)
#if isoform is not None:
#	print(probability)
#	for feature in isoform: print('\t'.join(feature))
		
