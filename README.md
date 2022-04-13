Algorithms
==========

In the Korf Lab, the goal is always to write simple, beautiful code. When one
focuses on speed, efficiency, or cleverness, code becomes difficult to
maintain, extend, and debug.

## Languages ##

Python is a very useful language. However, it is not the fastest language.
Sometimes we need more speed or less memory usage. For these purposes there are
compiled languages like C or Go. Before you make the compiled version of a
program, build the Python version first. That may be sufficient. If you need
something faster later, the Python program will help to ensure that both
programs are getting the correct answers.

## Small Tools Manifesto ##

It is worth reading the [Small Tools
Manifesto](https://github.com/pjotrp/bioinformatics) which is basically a
philosophy that bioinformatics programs should work like other Linux CLI
programs: small, interoperable, open source, documented, etc.

## Demonstration Programs ##

This repo is a collection of several _demonstration_ programs that are designed
to be simple and beautiful. Think of these as _templates_ for your own
projects. Don't extend these source codes. They are purposefully too simple and
too self-contained. Each program is written in multiple languages to get the
feel of how to program in that language. Generally, we seek to follow the
community style guides, however there is always some flexibility.

+ randomseq - generate random FASTA files of DNA sequences
+ kmerfreq - read sequences, output a table of kmer frequencies
+ dust - read sequences, output sequences with low complexity regions masked
+ longestorf - read sequences, output the longest open reading frame of each
+ smithwaterman - read sequences, produce local alignments

## Not Just Programs ##

Providing programs that work is only part of the goal. In order to be 
"beautiful" every program must have other qualities.

+ GitHub
+ Documentation
+ Testing

### GitHub ###

All software should be in GitHub repositories and have open source licenses. 
Following standard practices, all repos should have a `README.md` that 
describes the intent of the software.

### Documentation ###

The minimal documentation for every program is a usage statement that is 
reported when the program is run without arguments or with a help flag such as 
`-h` or `--help`. Usage statements should follow standard Unix practices.

All programs should have command line arguments. Use the most standard CLI 
library for the language.

Most projects should also include a `TUTORIAL.md` that walks a user through 
running a program or using a library.

### Testing ###

Almost every program should come with a small set of test data. This is used 
for both automated testing and tutorials.

Libraries should have unit tests.

Programs should have functional tests. 

## Programs ##

Each program is described in more detail below.

### randomseq ###

Generate random DNA sequences of fixed length. The composition of the sequences
defaults to 25% for each nucleotide, but the program should take other
distributions of mononucleotides.

Inputs

+ Number of sequences to generate
+ Length of each sequence
+ Probability of each letter
+ Random seed

Outputs

+ Multi-FASTA format to STDOUT

### kmerfreq ###

Determine the k-mer frequencies in a FASTA file. The value of K should be an 
argument with a default parameter (e.g. 3). Output format should include 
tab-separated and JSON.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ K-mer size

Outputs

+ TSV
+ JSON

### dust ###

Mask sequences with an entropy filter. The window size and entropy should have
default parameters and command line options. There should be an option to
change the output from N-based (hard) masking to lowercase (soft) masking.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ Window size
+ Entropy threshold
+ N-based or lowercase masking

Outputs

+ Multi-FASTA file to STDOUT

### longestorf ###

Translate each sequence and provide the protein sequence with the longest open 
reading frame. There should be an option to do 3- or 6-frame translations.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ Single or double-stranded translation

Outputs

+ Multi-FASTA file to STDOUT

### smithwaterman ###

Classic local alignment algorithm using match, mismatch, and gap scores. There 
should be a query sequence and database, both in FASTA format. Output formats 
should include tabular (score, coordinates) and human readable (alignments).

Inputs

+ Query sequence in FASTA format
+ Database sequence in Multi-FASTA format (gzipped ok)
+ Match score
+ Mismatach score
+ Gap score

Outputs

+ Tabular format
+ Alignment format

## Checklist ##

+ GitHub
+ Open source license
+ README.md
+ TUTORIAL.md
+ Test data
+ Unit/functional tests
+ Usage statement
+ Unix-standard CLI
+ Inputs as specified
+ Outputs as specified
+ Performs job as expected
