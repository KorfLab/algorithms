import argparse
import gzip
import itertools
import math

"""
need to do transition probs
need to do fill
need to do decoding
"""

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
    for name, sequence in readfasta(fname):
        if order == 0:
            for letter in sequence:
                model[letter] += 1
        else:
            j = 0
            for i in range(order, len(sequence)):
                kmer = sequence[j:i]
                letter = sequence[i]

                model[kmer][letter] += 1
                j += 1
    return model

def prob2log(prob):
    if prob == 0:
        return -100
    else:
        prob = math.log2(prob)
        return prob

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
            emPs[letter] = prob2log(emPs[letter])

    return emPs

"""
Transition Probabilites:

Start   index      Destination
Geno    0          Geno
                    R+
                    R-

R+      1           Geno
                    R+

R-      2           Geno
                    R-
"""
def transPs():
    transP = [[0 for i in range(3)] for j in range(3)]]



    transP[0][0], transP[0][1], transP[0][2] = 0.99999, 5e-6, "5e-6
    transP[1][0], transP[1][1], transP[1][2] = 299/300, 1/300, 0
    transP[2][0], transP[2][1], transp[2][2] = 299/300, 0, 1/300

    for col in range(len(transP)):
        for row in range(len(transP[col])):
            transP[col][row] = prob2log(transP[col][row])

    return transP
