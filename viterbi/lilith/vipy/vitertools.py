import gzip
import itertools

class vnode:
    def __init__(self):
        self.probs = {}
        self.mp = ""


class fastaseq:
    def __init__(self, name, sequence):
        self.name = name
        self.sequence = sequence
"""
need to train model
need o logodds
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
            model[kmer] = {}
            for base in 'ACTG':
                model[kmer][base] = 0
    return model

def trainmodel(fname, order, model):
    seqs = []
    for name, sequence in readfasta(fname):
        seq = fastaseq(name, sequence)
