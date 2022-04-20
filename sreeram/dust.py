
import argparse
import math
from readfasta import read_record

arg_parser = argparse.ArgumentParser(
    description='returns a seqience with a mask based on entropy')

arg_parser.add_argument('fasta_file', type=str, metavar='file',
                        help='enter the path to the fasta file you want to use')

arg_parser.add_argument('-window', required=False, type=int, default=11,
                        metavar='window_size', help='enter the window size (defaults to 11)')

arg_parser.add_argument('-entropy', required=False, type=float, default=1.5,
                        metavar='entropy_threshold', help='enter the entropy threshold you want to use (defauts to 1.5)')

arg_parser.add_argument('-lower', required=False, default=False,
                        type=bool, metavar='lower_case_mask', help='enter true if you want to mask to lowercase (default mask is N)')

arg = arg_parser.parse_args()


def entropy_calc(seq):
    A,C,T,G = 0,0,0,0
    for nuctide in seq:
        if nuctide == 'A':
            A+= 1
        if nuctide == 'C':
            C+= 1
        if nuctide == 'G':
            G+= 1
        if nuctide == 'T':
            T+= 1
    total = A+C+G+T

    prob_a = A/total
    prob_c = C/total
    prob_t = T/total
    prob_g = G/total

    entropy = 0
    if prob_a > 0:
        entropy -= prob_a * math.log(prob_a)
    if prob_c > 0:
        entropy -= prob_c * math.log(prob_c)
    if prob_g > 0:
        entropy -= prob_g * math.log(prob_g)
    if prob_t > 0:
        entropy -= prob_t * math.log(prob_t)
    
    entropy /= math.log(2)

    return entropy






cur_seq = ""
total_seq = []
num_seq = []

for id, seq in read_record(arg.fasta_file):
    cur_line = seq

    if seq[0] == '>':
        if len(num_seq) > 0:
            total_seq.append(cur_seq)
            cur_seq = ""
        
        num_seq.append(cur_line[1:])
    else:
        cur_seq += cur_line

total_seq.append(cur_seq)

# entropy calc
lower = arg.lower
threshold = arg.entropy
window = arg.window
for i in range(0, len(num_seq)):
    seq_used = total_seq[i]

    for j in range(0, len(seq_used)-window+1):
        cur_entropy = entropy_calc(seq_used[j:j+window])

        if cur_entropy < threshold:
            if lower:
                cur_seq = cur_seq[:j+int(window/2)] + cur_seq[j+int(window/2)].lower() + cur_seq[j+int((window/2))+1:]
            else:
                cur_seq = cur_seq[:j+int(window/2)] + "N" + cur_seq[j+int((window/2))+1:]

    print('f' + id)
    print(cur_seq)
