import sys
import json
from readfasta import read_record

mm_file = sys.argv[1]
km_file = sys.argv[2]
fa_file = sys.argv[3]

with open(mm_file) as fh: mm_data = json.load(fh)
with open(km_file) as fh: km_data = json.load(fh)

mm, km = mm_data['emissions'], km_data['emissions']

mm_count = {'pos': 0, 'neg': 0, 'gnm': 0}
km_count = {'pos': 0, 'neg': 0, 'gnm': 0}

for idn, seq in read_record(fa_file):
	# markov
	mm_highest_score = None
	mm_highest_state = None
	for state in mm:
		cur_score = 0
		model = mm[state]
		for kmer in model:
			n = len(kmer)
			break
		for i in range(len(seq)-n):
			given = seq[i:i+n]
			curbp = seq[i+n]
			cur_score += model[given][curbp]
		if mm_highest_score is None or cur_score > mm_highest_score:
			mm_highest_score = cur_score
			mm_highest_state = state
	mm_count[mm_highest_state] += 1
	
	# kmer
	km_highest_score = None
	km_highest_state = None
	for state in km:
		cur_score = 0
		model = km[state]
		for kmer in model:
			k = len(kmer)
			break
		for i in range(len(seq)-k+1):
			kmer = seq[i:i+k]
			cur_score += model[kmer]
		if km_highest_score is None or cur_score > km_highest_score:
			km_highest_score = cur_score
			km_highest_state = state
	km_count[km_highest_state] += 1
	
print(mm_count)	
print(km_count)


