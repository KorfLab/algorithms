import argparse
import gzip
import itertools
import math

class vnode:
    def __init__(self):
        self.probs = {"geno": 0, "pos": 0, "neg": 0}
        self.trace = ""
        self.maxp = None


class fastaseq:
    def __init__(self, name, sequence):
        self.name = name
        self.sequence = sequence
"""
need to do transition probs
need to do fill
need to do decoding
"""
parser = argparse.ArgumentParser(description="Algorithm to detect R-loops.")
    parser.add_argument('fname', metavar='<file>', help="query file")
    parser.add_argument('--pos', metavar='<file>', help="positive strand training file")
    parser.add_argument('--neg', metavar='<file>', help="negative strand training file")
    parser.add_argument('--n', metavar='<int>', type=int, help="order of HMM to construct")

args = parser.parse_args()

def readfasta (fname):
    seqname = ""
    s = []

    fp = None
    if fname.endswith(".gz"): fp = gzip.open(fname, 'rt')
    else:                     fp = open(fname)

    while True:
        line = fp.readline()
        if line == '': break
        line = line.rstrip()

        if line.startswith('>'):
            if len(s) > 0:
                yield(seqname, ''.join(s))
                seqname = line[1:]
                s = []
            else:
                seqname = line[1:]
        else:
            s.append(line)
    yield(seqname, ''.join(s))
    fp.close()

def makemodel(order):
    model = {}

    if order == 0:
        for base in 'ACTG':
            model[base] = 0
    else:
        for kmer in itertools.product('ACTG', repeat=order):
            joined = ''.join(kmer)
            model[joined] = {}
            for base in 'ACTG':
                model[joined][base] = 0
    return model

def trainmodel(fname, order, model):
    seqs = []
    for name, sequence in readfasta(fname):
        seq = fastaseq(name, sequence)
        seqs.append(seq)

    if order == 0:
        for seq in seqs:
            for letter in seq.sequence:
                model[letter] += 1
    else:
        for seq in seqs:
            j = 0
            for i in range(order, len(seq.sequence)):
                kmer = seq.sequence[j:i]
                letter = seq.sequence[i]

                model[kmer][letter] += 1

                j += 1
    return model

def logodds(prob, nullP):
    if prob == 0:
        return -100
    else:
        prob = math.log2(prob/nullP) #divide by null prob to get logodds

def emissionP(model, order, log):
    sum = 0
    emPs = {'A': 0, 'C': 0, 'G': 0, 'T': 0}

    if order == 0:
        for letter in model:
            sum += model[letter]
            emPs[letter] += model[letter]
    else:
        for kmer in model:
            for letter in kmer:
                sum += model[kmer][letter]
                emPs[letter] += model[kmer][letter]

    for letter in emPs:
        emPs[letter] = round(emPs[letter] / sum, 2)

        if log:
            emPs[letter] = logodds(emPs[letter], 0.25)

    return emPs

"""
Transition Probabilites:

Start   Destination 
Geno    Geno
        R+
        R-

R+      R+
        Geno

R-      R-
        Geno
"""
transitionP = {}

transitionP["geno"] = {"geno" : 0.99999, "pos": 5e-6, "neg": 5e-6}
transitionP["rpos"] = {"rpos" : 299/300, "geno": 1/300}
transitionP["rneg"] = {"rneg" : 299/300, "geno": 1/300}

for start in transitionP:
    for end in start:
        transitionP[start][end] = logodds(transitionP[start][end], 0.25)

#make and train models

posmodel = makemodel(args.n)
negmodel = makemodel(args.n)

posmodel = train(args.pos, args.n, posmodel)
negmodel = train(args.neg, args.n, negmodel)

#emission probs

posemPs = emissionP(posmodel, args.n, True)
negemPs = emissionP(negmodel, args.n, True)
genoemPs = {"A": 0, "C": 0, "G": 0, "T": 0}

#initialize fill

viterbi = []

initial = vnode()

initial.probs["geno"] = 1
initial.probs["pos"] = 0
initial.probs["neg"] = 0

viterbi.append(initial)


#fill

for name, sequence in readfasta(args.fname):
